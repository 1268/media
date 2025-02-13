package rakuten

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (a *Address) String() string {
   var data strings.Builder
   data.WriteString(a.MarketCode)
   data.WriteByte('/')
   if a.SeasonId != "" {
      data.WriteString("player/episodes/stream/")
      data.WriteString(a.SeasonId)
   } else {
      data.WriteString("movies")
   }
   data.WriteByte('/')
   data.WriteString(a.ContentId)
   return data.String()
}

func (g *GizmoContent) Hd(classification_id int, language string) *OnDemand {
   return g.video(classification_id, language, "HD")
}

// hard geo block
func (o *OnDemand) Streamings() (*StreamInfo, error) {
   data, err := json.MarshalIndent(o, "", " ")
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
         StreamInfos []StreamInfo `json:"stream_infos"`
      }
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   // you can trigger this with wrong location
   if err := value.Errors; len(err) >= 1 {
      return nil, errors.New(err[0].Message)
   }
   return &value.Data.StreamInfos[0], nil
}

func (g *GizmoContent) String() string {
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
   b = append(b, "\ntype = "...)
   b = append(b, g.Type...)
   return string(b)
}

func (a *Address) Movie(classification_id int) (*GizmoContent, error) {
   req, _ := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   req.URL.Path = "/v3/movies/" + a.ContentId
   req.URL.RawQuery = url.Values{
      "classification_id": {strconv.Itoa(classification_id)},
      "device_identifier": {"atvui40"},
      "market_code":       {a.MarketCode},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var value struct {
      Data GizmoContent
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data, nil
}

func (a *Address) Set(data string) error {
   data = strings.TrimPrefix(data, "https://")
   data = strings.TrimPrefix(data, "www.")
   data = strings.TrimPrefix(data, "rakuten.tv")
   data = strings.TrimPrefix(data, "/")
   a.MarketCode, data, _ = strings.Cut(data, "/")
   var found bool
   data, a.ContentId, found = strings.Cut(data, "movies/")
   if !found {
      data = strings.TrimPrefix(data, "player/episodes/stream/")
      a.SeasonId, a.ContentId, _ = strings.Cut(data, "/")
   }
   return nil
}

func (g GizmoSeason) Content(web *Address) (*GizmoContent, bool) {
   for _, episode := range g.Episodes {
      if episode.Id == web.ContentId {
         return &episode, true
      }
   }
   return nil, false
}

func (a *Address) Season(classification_id int) (*GizmoSeason, error) {
   req, _ := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   req.URL.Path = "/v3/seasons/" + a.SeasonId
   req.URL.RawQuery = url.Values{
      "classification_id": {strconv.Itoa(classification_id)},
      "device_identifier": {"atvui40"},
      "market_code":       {a.MarketCode},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var value struct {
      Data GizmoSeason
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data, nil
}

// github.com/pandvan/rakuten-m3u-generator/blob/master/rakuten.py
func (a *Address) ClassificationId() (int, bool) {
   switch a.MarketCode {
   case "at":
      return 300, true
   case "ch":
      return 319, true
   case "cz":
      return 272, true
   case "de":
      return 307, true
   case "fr":
      return 23, true
   case "ie":
      return 41, true
   case "nl":
      return 69, true
   case "pl":
      return 277, true
   case "se":
      return 282, true
   case "uk":
      return 18, true
   }
   return 0, false
}

func (g *GizmoContent) video(
   classification_id int, language, quality string,
) *OnDemand {
   return &OnDemand{
      DeviceStreamVideoQuality: quality,
      AudioLanguage:            language,
      AudioQuality:             "2.0",
      ClassificationId:         classification_id,
      DeviceIdentifier:         "atvui40",
      DeviceSerial:             "not implemented",
      Player:                   "atvui40:DASH-CENC:WVM",
      SubtitleLanguage:         "MIS",
      VideoType:                "stream",
      ContentId:                g.Id,
      ContentType:              g.Type,
   }
}

func (g *GizmoContent) Fhd(classification_id int, language string) *OnDemand {
   return g.video(classification_id, language, "FHD")
}

func (s *StreamInfo) Mpd() (*http.Response, error) {
   return http.Get(s.Url)
}

func (s *StreamInfo) License(data []byte) (*http.Response, error) {
   return http.Post(
      s.LicenseUrl, "application/x-protobuf", bytes.NewReader(data),
   )
}

///

type OnDemand struct {
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

type StreamInfo struct {
   LicenseUrl   string `json:"license_url"`
   Url          string
   VideoQuality string `json:"video_quality"`
}

type Address struct {
   MarketCode string
   SeasonId   string
   ContentId  string
}

type GizmoSeason struct {
   Episodes []GizmoContent
}

type GizmoContent struct {
   ViewOptions  struct {
      Private struct {
         Streams []struct {
            AudioLanguages []struct {
               Id string
            } `json:"audio_languages"`
         }
      }
   } `json:"view_options"`
   Id           string
   Type         string
}

