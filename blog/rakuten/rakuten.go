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

func (a *address) classification_id() int {
   switch a.market_code {
   case "cz":
      return 272
   case "dk":
      return 283
   case "fi":
      return 284
   case "fr":
      return 23
   case "ie":
      return 41
   case "it":
      return 36
   case "nl":
      return 323
   case "no":
      return 286
   case "pt":
      return 64
   case "se":
      return 282
   case "ua":
      return 276
   case "uk":
      return 18
   }
   panic(a.market_code)
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
   o.AudioLanguage = "ENG"
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

func (a *address) movie() (*hello, error) {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/movies/" + a.content_id
   req.URL.RawQuery = url.Values{
      "classification_id": {
         strconv.Itoa(a.classification_id()),
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
      Data hello
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data, nil
}

type hello struct {
   Number       int
   SeasonNumber int `json:"season_number"`
   Title        string
   TvShowTitle  string `json:"tv_show_title"`
   ViewOptions  struct {
      Private struct {
         Streams []struct {
            AudioLanguages []struct {
               Id string
            } `json:"audio_languages"`
         }
      }
   } `json:"view_options"`
   Year int
}

type address struct {
   market_code string
   season_id   string
   content_id  string
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

func (a *address) season() ([]hello, error) {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/seasons/" + a.season_id
   req.URL.RawQuery = url.Values{
      "classification_id": {
         strconv.Itoa(a.classification_id()),
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
      Data struct {
         Episodes []hello
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return value.Data.Episodes, nil
}
