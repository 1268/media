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

// geo block
func (o on_demand) streamings() ([]stream_info, error) {
   o.AudioQuality = "2.0"
   o.ContentType = "movies"
   o.DeviceIdentifier = "atvui40"
   o.DeviceStreamVideoQuality = "FHD"
   o.Player = "atvui40:DASH-CENC:WVM"
   o.SubtitleLanguage = "MIS"
   o.VideoType = "stream"
   o.DeviceSerial = "not implemented"
   // cz = fail
   // fr = pass
   o.AudioLanguage = "ENG"
   // cz = pass
   // fr = fail
   //o.AudioLanguage = "SPA"
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
   return value.Data.StreamInfos, nil
}

type stream_info struct {
   LicenseUrl   string `json:"license_url"`
   Url          string
   VideoQuality string `json:"video_quality"`
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
type data struct {
   Year int
   Title string
   ViewOptions struct {
      Private struct {
         Streams []struct {
            AudioLanguages []struct {
               Id string
            } `json:"audio_languages"`
         }
      }
   } `json:"view_options"`
}

type season struct {
   data
   Episodes []data
}

func (a *address) get_season() (*season, error) {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v3/")
      if a.season != "" {
         b.WriteString("seasons/")
         b.WriteString(a.season)
      } else {
         b.WriteString("movies/")
         b.WriteString(a.content_id)
      }
      return b.String()
   }()
   req.URL.RawQuery = url.Values{
      "classification_id": {
         strconv.Itoa(classification_id[a.market_code]),
      },
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
      Data season
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data, nil
}
var classification_id = map[string]int{
   "cz": 272,
   "dk": 283,
   "fi": 284,
   "fr": 23,
   "ie": 41,
   "it": 36,
   "nl": 323,
   "no": 286,
   "pt": 64,
   "se": 282,
   "ua": 276,
   "uk": 18,
}

type address struct {
   market_code string
   season      string
   content_id string
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
      a.season, a.content_id, found = strings.Cut(data, "/")
      if !found {
         return errors.New("episode not found")
      }
   }
   return nil
}
