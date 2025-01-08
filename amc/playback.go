package amc

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

func (a *Authorization) Playback(nid string) (*Playback, error) {
   var value struct {
      AdTags struct {
         Lat int `json:"lat"`
         Mode string `json:"mode"`
         Ppid int `json:"ppid"`
         PlayerHeight int `json:"playerHeight"`
         PlayerWidth int `json:"playerWidth"`
         Url string `json:"url"`
      } `json:"adtags"`
   }
   value.AdTags.Mode = "on-demand"
   value.AdTags.Url = "-"
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/playback-id/api/v1/playback/" + nid
   req.Header = http.Header{
      "authorization": {"Bearer " + a.Data.AccessToken},
      "content-type": {"application/json"},
      "x-amcn-device-ad-id": {"-"},
      "x-amcn-language": {"en"},
      "x-amcn-network": {"amcplus"},
      "x-amcn-platform": {"web"},
      "x-amcn-service-id": {"amcplus"},
      "x-amcn-tenant": {"amcn"},
      "x-ccpa-do-not-sell": {"doNotPassData"},
   }
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
   var play struct {
      Data struct {
         PlaybackJsonData struct {
            Sources []DataSource
         }
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&play)
   if err != nil {
      return nil, err
   }
   return &Playback{resp.Header, play.Data.PlaybackJsonData.Sources}, nil
}

type DataSource struct {
   KeySystems *struct {
      Widevine struct {
         LicenseUrl string `json:"license_url"`
      } `json:"com.widevine.alpha"`
   } `json:"key_systems"`
   Src string
   Type string
}

type Playback struct {
   Header http.Header
   Sources []DataSource
}

func (p *Playback) Dash() (*Wrapper, bool) {
   for _, source := range p.Sources {
      if source.Type == "application/dash+xml" {
         return &Wrapper{p.Header, source}, true
      }
   }
   return nil, false
}

type Wrapper struct {
   Header http.Header
   Source DataSource
}

func (w *Wrapper) Wrap(data []byte) ([]byte, error) {
   req, err := http.NewRequest(
      "POST", w.Source.KeySystems.Widevine.LicenseUrl, bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("bcov-auth", w.Header.Get("x-amcn-bc-jwt"))
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}
