package rakuten

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

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

type stream_info struct {
   LicenseUrl   string `json:"license_url"`
   Url          string
   VideoQuality string `json:"video_quality"`
}

func (a *address) classification_id() (string, error) {
   var v int
   switch a.market_code {
   case "cz":
      v = 272
   case "dk":
      v = 283
   case "fi":
      v = 284
   case "fr":
      v = 23
   case "ie":
      v = 41
   case "it":
      v = 36
   case "nl":
      v = 323
   case "no":
      v = 286
   case "pt":
      v = 64
   case "se":
      v = 282
   case "ua":
      v = 276
   case "uk":
      v = 18
   default:
      return "", errors.New(a.market_code)
   }
   return strconv.Itoa(v), nil
}

func (a *address) season() (*gizmo_season, error) {
   classification, err := a.classification_id()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/seasons/" + a.season_id
   req.URL.RawQuery = url.Values{
      "classification_id": {classification},
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

func (a *address) movie() (*gizmo_content, error) {
   classification, err := a.classification_id()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/movies/" + a.content_id
   req.URL.RawQuery = url.Values{
      "classification_id": {classification},
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
      Data gizmo_content
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data, nil
}

type gizmo_season struct {
   Episodes []gizmo_content
}

func (a *address) content(season *gizmo_season) (*gizmo_content, bool) {
   for _, episode := range season.Episodes {
      if episode.Id == a.content_id {
         return &episode, true
      }
   }
   return nil, false
}

type gizmo_content struct {
   Id           string
   Number       int
   SeasonNumber int `json:"season_number"`
   Title        string
   TvShowTitle  string `json:"tv_show_title"`
   Type         string
   Year int
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
