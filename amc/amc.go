package amc

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

func (a *Authorization) Playback(web Address) (*Playback, error) {
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
   req.URL.Path = "/playback-id/api/v1/playback/" + web[1]
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
      var data strings.Builder
      resp.Write(&data)
      return nil, errors.New(data.String())
   }
   var value1 struct {
      Data struct {
         PlaybackJsonData struct {
            Sources []DataSource
         }
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value1)
   if err != nil {
      return nil, err
   }
   return &Playback{resp.Header, value1.Data.PlaybackJsonData.Sources}, nil
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

type Address [2]string

func (a *Address) Set(data string) error {
   data = strings.TrimPrefix(data, "https://")
   data = strings.TrimPrefix(data, "www.")
   data = strings.TrimPrefix(data, "amcplus.com")
   var found bool
   (*a)[0], (*a)[1], found = strings.Cut(data, "--")
   if !found {
      return errors.New("--")
   }
   return nil
}

func (a Address) String() string {
   return strings.Join(a[:], "--")
}

func (a *Authorization) Login(email, password string) ([]byte, error) {
   data, err := json.Marshal(map[string]string{
      "email": email,
      "password": password,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/auth-orchestration-id/api/v1/login"
   req.Header = http.Header{
      "authorization": {"Bearer " + a.Data.AccessToken},
      "content-type": {"application/json"},
      "x-amcn-device-ad-id": {"-"},
      "x-amcn-device-id": {"-"},
      "x-amcn-language": {"en"},
      "x-amcn-network": {"amcplus"},
      "x-amcn-platform": {"web"},
      "x-amcn-service-group-id": {"10"},
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
      return nil, errors.New(resp.Status)
   }
   return io.ReadAll(resp.Body)
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

type Authorization struct {
   Data struct {
      AccessToken string `json:"access_token"`
      RefreshToken string `json:"refresh_token"`
   }
}

func (a *Authorization) Unmarshal(data []byte) error {
   return json.Unmarshal(data, a)
}

func (a *Authorization) Unauth() error {
   req, err := http.NewRequest("POST", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/auth-orchestration-id/api/v1/unauth"
   req.Header = http.Header{
      "x-amcn-device-id": {"-"},
      "x-amcn-language": {"en"},
      "x-amcn-network": {"amcplus"},
      "x-amcn-platform": {"web"},
      "x-amcn-tenant": {"amcn"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   return json.NewDecoder(resp.Body).Decode(a)
}

func (a *Authorization) Refresh() ([]byte, error) {
   req, err := http.NewRequest("POST", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/auth-orchestration-id/api/v1/refresh"
   req.Header.Set("authorization", "Bearer " + a.Data.RefreshToken)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   return io.ReadAll(resp.Body)
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
