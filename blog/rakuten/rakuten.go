package rakuten

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (s *stream_info) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      s.LicenseUrl, "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

// geo block
func (o *on_demand) streamings() (*stream_info, error) {
   data, err := json.Marshal(o)
   if err != nil {
      return nil, err
   }
   resp, err := http.Post(
      "https://gizmo.rakuten.tv/v3/avod/streamings",
      "application/json", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   var value struct {
      Data struct {
         StreamInfos []stream_info `json:"stream_infos"`
      }
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   if err := value.Errors; len(err) >= 1 {
      return nil, errors.New(err[0].Message)
   }
   return &value.Data.StreamInfos[0], nil
}

type gizmo_season struct {
   Episodes []gizmo_content
}

func (g *gizmo_content) String() string {
   var (
      audio = map[string]struct{}{}
      b     []byte
   )
   for _, stream := range g.ViewOptions.Private.Streams {
      for _, language := range stream.AudioLanguages {
         _, ok := audio[language.Id]
         if !ok {
            if b != nil {
               b = append(b, '\n')
            }
            b = append(b, "audio language = "...)
            b = append(b, language.Id...)
            audio[language.Id] = struct{}{}
         }
      }
   }
   b = append(b, "\nid = "...)
   b = append(b, g.Id...)
   if g.Number >= 1 {
      b = append(b, "\nnumber = "...)
      b = strconv.AppendInt(b, int64(g.Number), 10)
   }
   if g.SeasonNumber >= 1 {
      b = append(b, "\nseason number = "...)
      b = strconv.AppendInt(b, int64(g.SeasonNumber), 10)
   }
   b = append(b, "\ntitle = "...)
   b = append(b, g.Title...)
   if g.TvShowTitle != "" {
      b = append(b, "\ntv show = "...)
      b = append(b, g.TvShowTitle...)
   }
   b = append(b, "\ntype = "...)
   b = append(b, g.Type...)
   b = append(b, "\nyear = "...)
   b = strconv.AppendInt(b, int64(g.Year), 10)
   return string(b)
}

type stream_info struct {
   LicenseUrl   string `json:"license_url"`
   Url          string
   VideoQuality string `json:"video_quality"`
}

func (g *gizmo_content) video(
   classification_id int, language, quality string,
) *on_demand {
   return &on_demand{
      AudioLanguage:            language,
      AudioQuality:             "2.0",
      ClassificationId:         classification_id,
      ContentId:                g.Id,
      ContentType:              g.Type,
      DeviceIdentifier:         "atvui40",
      DeviceSerial:             "not implemented",
      DeviceStreamVideoQuality: quality,
      Player:                   "atvui40:DASH-CENC:WVM",
      SubtitleLanguage:         "MIS",
      VideoType:                "stream",
   }
}

func (g *gizmo_content) Fhd(classification_id int, language string) *on_demand {
   return g.video(classification_id, language, "FHD")
}

func (g *gizmo_content) hd(classification_id int, language string) *on_demand {
   return g.video(classification_id, language, "HD")
}

type gizmo_content struct {
   Id           string
   Number       int
   SeasonNumber int `json:"season_number"`
   Title        string
   TvShowTitle  string `json:"tv_show_title"`
   Type         string
   Year         int
   ViewOptions  struct {
      Private struct {
         Streams []struct {
            AudioLanguages []struct {
               Id string
            } `json:"audio_languages"`
         }
      }
   } `json:"view_options"`
}

type on_demand struct {
   AudioLanguage            string `json:"audio_language"`
   AudioQuality             string `json:"audio_quality"`
   ClassificationId         int    `json:"classification_id"`
   ContentId                string `json:"content_id"`
   ContentType              string `json:"content_type"`
   DeviceIdentifier         string `json:"device_identifier"`
   DeviceSerial             string `json:"device_serial"`
   DeviceStreamVideoQuality string `json:"device_stream_video_quality"`
   Player                   string `json:"player"`
   SubtitleLanguage         string `json:"subtitle_language"`
   VideoType                string `json:"video_type"`
}

func (a *address) season(classification_id int) (*gizmo_season, error) {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/seasons/" + a.season_id
   req.URL.RawQuery = url.Values{
      "classification_id": {strconv.Itoa(classification_id)},
      "device_identifier": {"atvui40"},
      "market_code":       {a.market_code},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   var value struct {
      Data gizmo_season
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data, nil
}

func (a *address) movie(classification_id int) (*gizmo_content, error) {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/movies/" + a.content_id
   req.URL.RawQuery = url.Values{
      "classification_id": {strconv.Itoa(classification_id)},
      "device_identifier": {"atvui40"},
      "market_code":       {a.market_code},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var value struct {
      Data   gizmo_content
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   if err := value.Errors; len(err) >= 1 {
      return nil, errors.New(err[0].Message)
   }
   return &value.Data, nil
}

func (n namer) Show() string {
   return n.g.TvShowTitle
}

func (n namer) Season() int {
   return n.g.SeasonNumber
}

func (n namer) Episode() int {
   return n.g.Number
}

func (n namer) Title() string {
   return n.g.Title
}

type namer struct {
   g *gizmo_content
}

func (n namer) Year() int {
   return n.g.Year
}

func (a *address) content(season *gizmo_season) (*gizmo_content, bool) {
   for _, episode := range season.Episodes {
      if episode.Id == a.content_id {
         return &episode, true
      }
   }
   return nil, false
}

func (a *address) classification_id() (int, bool) {
   switch a.market_code {
   case "cz":
      return 272, true
   case "dk":
      return 283, true
   case "fi":
      return 284, true
   case "fr":
      return 23, true
   case "ie":
      return 41, true
   case "it":
      return 36, true
   case "nl":
      return 323, true
   case "no":
      return 286, true
   case "pt":
      return 64, true
   case "se":
      return 282, true
   case "ua":
      return 276, true
   case "uk":
      return 18, true
   }
   return 0, false
}

type address struct {
   market_code string
   season_id   string
   content_id  string
}

func (a *address) String() string {
   var data strings.Builder
   data.WriteString(a.market_code)
   data.WriteByte('/')
   if a.season_id != "" {
      data.WriteString("player/episodes/stream/")
      data.WriteString(a.season_id)
      data.WriteByte('/')
   } else {
      data.WriteString("movies/")
   }
   data.WriteString(a.content_id)
   return data.String()
}

func (a *address) Set(data string) error {
   data = strings.TrimPrefix(data, "https://")
   data = strings.TrimPrefix(data, "www.")
   data = strings.TrimPrefix(data, "rakuten.tv")
   data = strings.TrimPrefix(data, "/")
   var found bool
   a.market_code, data, found = strings.Cut(data, "/")
   if !found {
      return errors.New("market code not found")
   }
   data, a.content_id, found = strings.Cut(data, "movies/")
   if !found {
      data = strings.TrimPrefix(data, "player/episodes/stream/")
      a.season_id, a.content_id, found = strings.Cut(data, "/")
      if !found {
         return errors.New("episode not found")
      }
   }
   return nil
}
