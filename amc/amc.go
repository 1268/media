package amc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "net/url"
   "os"
   "strings"
)

func (a Auth) Playback(ref string) (*Playback, error) {
   // request body
   req_body := func(r *http.Request) error {
      var s struct {
         Ad_Tags struct {
            Lat int `json:"lat"`
            Mode string `json:"mode"`
            PPID int `json:"ppid"`
            Player_Height int `json:"playerHeight"`
            Player_Width int `json:"playerWidth"`
            URL string `json:"url"`
         } `json:"adtags"`
      }
      s.Ad_Tags.Mode = "on-demand"
      s.Ad_Tags.URL = "-"
      b, err := json.MarshalIndent(s, "", " ")
      if err != nil {
         return err
      }
      r.Body_Bytes(b)
      return nil
   }
   // request URL path
   req_URL_path := func(r *http.Request) error {
      _, nID, found := strings.Cut(ref, "--")
      if !found {
         return errors.New("nid not found")
      }
      r.URL.Path += nID
      return nil
   }
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "gw.cds.amcn.com",
      Path: "/playback-id/api/v1/playback/",
   })
   err := req_URL_path(req)
   if err != nil {
      return nil, err
   }
   if err := req_body(req); err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "Content-Type": {"application/json"},
      "X-Amcn-Device-Ad-ID": {"-"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Service-ID": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Ccpa-Do-Not-Sell": {"doNotPassData"},
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var p Playback
   if err := p.body(res); err != nil {
      return nil, err
   }
   p.Header = res.Header
   return &p, nil
}

type Playback struct {
   http.Header
   sources []Source
}

func (p *Playback) body(res *http.Response) error {
   var s struct {
      Data struct {
         Playback_JSON_Data struct {
            Sources []Source
         } `json:"playbackJsonData"`
      }
   }
   err := json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return err
   }
   p.sources = s.Data.Playback_JSON_Data.Sources
   return nil
}

func (p Playback) HTTP_DASH() *Source {
   for _, source := range p.sources {
      if strings.HasPrefix(source.Src, "http://") {
         if source.Type == "application/dash+xml" {
            return &source
         }
      }
   }
   return nil
}

func (Playback) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (p Playback) Request_Header() http.Header {
   token := p.Get("X-AMCN-BC-JWT")
   head := make(http.Header)
   head.Set("bcov-auth", token)
   return head
}

func (p Playback) Request_URL() string {
   return p.HTTP_DASH().Key_Systems.Widevine.License_URL
}

func (a *Auth) Login(email, password string) error {
   body := func(r *http.Request) error {
      b, err := json.Marshal(map[string]string{
         "email": email,
         "password": password,
      })
      if err != nil {
         return err
      }
      r.Body_Bytes(b)
      return nil
   }
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "gw.cds.amcn.com",
      Path: "/auth-orchestration-id/api/v1/login",
   })
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "Content-Type": {"application/json"},
      "X-Amcn-Device-Ad-ID": {"-"},
      "X-Amcn-Device-ID": {"-"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Service-Group-ID": {"10"},
      "X-Amcn-Service-ID": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Ccpa-Do-Not-Sell": {"doNotPassData"},
   }
   err := body(req)
   if err != nil {
      return err
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

func (a *Auth) Refresh() error {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "gw.cds.amcn.com",
      Path: "/auth-orchestration-id/api/v1/refresh",
   })
   req.Header.Set("Authorization", "Bearer " + a.Data.Refresh_Token)
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

func (a Auth) Write_File(name string) error {
   data, err := json.MarshalIndent(a, "", " ")
   if err != nil {
      return err
   }
   return os.WriteFile(name, data, 0666)
}

// JSON
type Auth struct {
   Data struct {
      Access_Token string
      Refresh_Token string
   }
}

func Read_Auth(name string) (*Auth, error) {
   data, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   a := new(Auth)
   if err := json.Unmarshal(data, a); err != nil {
      return nil, err
   }
   return a, nil
}

func Unauth() (*Auth, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "gw.cds.amcn.com",
      Path: "/auth-orchestration-id/api/v1/unauth",
   })
   req.Header = http.Header{
      "X-Amcn-Device-ID": {"-"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Tenant": {"amcn"},
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   auth := new(Auth)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

type Source struct {
   Key_Systems *struct {
      Widevine struct {
         License_URL string
      } `json:"com.widevine.alpha"`
   }
   Src string
   Type string
}
