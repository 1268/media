package amc

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

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

func (a *Auth) Unmarshal(data []byte) error {
   return json.Unmarshal(data, a)
}

func (a *Auth) Refresh() ([]byte, error) {
   req, _ := http.NewRequest("POST", "https://gw.cds.amcn.com", nil)
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

type Auth struct {
   Data struct {
      AccessToken string `json:"access_token"`
      RefreshToken string `json:"refresh_token"`
   }
}

func (a *Auth) Login(email, password string) ([]byte, error) {
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

func (a *Auth) Unauth() error {
   req, _ := http.NewRequest("POST", "https://gw.cds.amcn.com", nil)
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

func (a *Auth) Playback(web Address) (*Playback, error) {
   data, err := json.Marshal(map[string]any{
      "adtags": map[string]any{
         "lat": 0,
         "mode": "on-demand",
         "playerHeight": 0,
         "playerWidth": 0,
         "ppid": 0,
         "url": "-",
      },
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
   var play Playback
   err = json.NewDecoder(resp.Body).Decode(&play.Body)
   if err != nil {
      return nil, err
   }
   play.Header = resp.Header
   return &play, nil
}

type Source struct {
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
   Body struct {
      Data struct {
         PlaybackJsonData struct {
            Sources []Source
         }
      }
   }
}

type Client struct {
   Header http.Header
   Source Source
}

func (c *Client) Mpd() (*http.Response, error) {
   return http.Get(c.Source.Src)
}

func (p *Playback) Dash() (*Client, bool) {
   for _, source1 := range p.Body.Data.PlaybackJsonData.Sources {
      if source1.Type == "application/dash+xml" {
         return &Client{p.Header, source1}, true
      }
   }
   return nil, false
}

func (c *Client) License(data []byte) ([]byte, error) {
   req, err := http.NewRequest(
      "POST", c.Source.KeySystems.Widevine.LicenseUrl, bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("bcov-auth", c.Header.Get("x-amcn-bc-jwt"))
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}
