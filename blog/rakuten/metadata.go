package rakuten

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type metadata struct {
   Title       string
   ViewOptions struct {
      Public struct {
         Trailers []struct {
            AudioLanguages []struct {
               Id string
            } `json:"audio_languages"`
         }
      }
   } `json:"view_options"`
   Year int
}

func (a *address) metadata() (*metadata, error) {
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
      Data metadata
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data, nil
}
