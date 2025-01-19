package amc

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
   "time"
)

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

func cache_hash() string {
   return base64.StdEncoding.EncodeToString([]byte("ff="))
}

func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   a.Path = strings.TrimPrefix(s, "amcplus.com")
   var found bool
   _, a.Nid, found = strings.Cut(a.Path, "--")
   if !found {
      return errors.New("--")
   }
   return nil
}

type Address struct {
   Nid string
   Path string
}

func (a *Address) String() string {
   return a.Path
}

func (a *Authorization) Content(path string) (*ContentCompiler, error) {
   req, err := http.NewRequest("", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return nil, err
   }
   // If you request once with headers, you can request again without any
   // headers for 10 minutes, but then headers are required again
   req.Header = http.Header{
      "authorization": {"Bearer " + a.Data.AccessToken},
      "x-amcn-cache-hash": {cache_hash()},
      "x-amcn-network": {"amcplus"},
      "x-amcn-tenant": {"amcn"},
      "x-amcn-user-cache-hash": {cache_hash()},
   }
   // Shows must use `path`, and movies must use `path/watch`. If trial has
   // expired, you will get `.data.type` of `redirect`. You can remove the
   // `/watch` to resolve this, but the resultant response will still be
   // missing `video-player-ap`.
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/content-compiler-cr/api/v1/content/amcn/amcplus/path")
      if strings.HasPrefix(path, "/movies/") {
         b.WriteString("/watch")
      }
      b.WriteString(path)
      return b.String()
   }()
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
      Data ContentCompiler
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data, nil
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

type Authorization struct {
   Data struct {
      AccessToken string `json:"access_token"`
      RefreshToken string `json:"refresh_token"`
   }
}

func (a *Authorization) Unmarshal(data []byte) error {
   return json.Unmarshal(data, a)
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

type ContentCompiler struct {
   Children []struct {
      Properties struct {
         CurrentVideo CurrentVideo
      }
      Type string
   }
}

func (c *ContentCompiler) Video() (*CurrentVideo, bool) {
   for _, child := range c.Children {
      if child.Type == "video-player-ap" {
         return &child.Properties.CurrentVideo, true
      }
   }
   return nil, false
}

type CurrentVideo struct {
   Meta struct {
      Airdate time.Time // 1996-01-01T00:00:00.000Z
      EpisodeNumber int
      Season int `json:",string"`
      ShowTitle string
   }
   Text struct {
      Title string
   }
}

func (c *CurrentVideo) Title() string {
   return c.Text.Title
}

func (c *CurrentVideo) Year() int {
   return c.Meta.Airdate.Year()
}

func (c *CurrentVideo) Episode() int {
   return c.Meta.EpisodeNumber
}

func (c *CurrentVideo) Show() string {
   return c.Meta.ShowTitle
}

func (c *CurrentVideo) Season() int {
   return c.Meta.Season
}
