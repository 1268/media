package amc

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
   "time"
)

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

type Source struct {
   Key_Systems *struct {
      Widevine struct {
         License_URL string
      } `json:"com.widevine.alpha"`
   }
   Src string
   Type string
}

func (c Content) Video() (*Video, error) {
   for _, child := range c.Data.Children {
      if child.Type == "video-player-ap" {
         var s struct {
            Current_Video Video `json:"currentVideo"`
         }
         err := json.Unmarshal(child.Properties, &s)
         if err != nil {
            return nil, err
         }
         return &s.Current_Video, nil
      }
   }
   return nil, errors.New("video-player-ap not present")
}

func (v Video) Season() (int64, error) {
   return v.Meta.Season, nil
}

func (v Video) Episode() (int64, error) {
   return v.Meta.Episode_Number, nil
}

func (v Video) Series() string {
   return v.Meta.Show_Title
}

type Video struct {
   Meta struct {
      Show_Title string `json:"showTitle"`
      Season int64 `json:",string"`
      Episode_Number int64 `json:"episodeNumber"`
      Airdate string // 1996-01-01T00:00:00.000Z
   }
   Text struct {
      Title string
   }
}

func (v Video) Title() string {
   return v.Text.Title
}

func (v Video) Date() (time.Time, error) {
   return time.Parse(time.RFC3339, v.Meta.Airdate)
}

type Content struct {
   Data	struct {
      Children []struct {
         Properties json.RawMessage
         Type string
      }
   }
}

func (p Playback) Request_URL() string {
   return p.HTTP_DASH().Key_Systems.Widevine.License_URL
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

/////////////////////////////////////////////

// This accepts full URL or path only.
func (a Auth) Content(ref string) (*Content, error) {
   // Shows must use `path`, and movies must use `path/watch`. If trial has
   // expired, you will get `.data.type` of `redirect`. You can remove the
   // `/watch` to resolve this, but the resultant response will still be
   // missing `video-player-ap`.
   url_path := func(r *http.Request) error {
      p, err := url.Parse(ref)
      if err != nil {
         return err
      }
      if strings.HasPrefix(p.Path, "/movies/") {
         r.URL.Path += "/watch"
      }
      r.URL.Path += p.Path
      return nil
   }
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "gw.cds.amcn.com",
      Path: "/content-compiler-cr/api/v1/content/amcn/amcplus/path",
   })
   err := url_path(req)
   if err != nil {
      return nil, err
   }
   // If you request once with headers, you can request again without any
   // headers for 10 minutes, but then headers are required again
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   con := new(Content)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

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
   a := new(Auth)
   if err := json.NewDecoder(res.Body).Decode(a); err != nil {
      return nil, err
   }
   return a, nil
}
