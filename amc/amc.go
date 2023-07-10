package amc

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func (a *Auth_ID) Login(email, password string) error {
   body, err := json.Marshal(map[string]string{
      "email": email,
      "password": password,
   })
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/auth-orchestration-id/api/v1/login",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
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
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return json.NewDecoder(res.Body).Decode(a)
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

// This accepts full URL or path only.
func (a Auth_ID) Content(ref string) (*Content, error) {
   req, err := http.NewRequest("GET", "https://gw.cds.amcn.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/content-compiler-cr/api/v1/content/amcn/amcplus/path"
   // Shows must use `path`, and movies must use `path/watch`. If trial has
   // expired, you will get `.data.type` of `redirect`. You can remove the
   // `/watch` to resolve this, but the resultant response will still be
   // missing `video-player-ap`.
   {
      p, err := url.Parse(ref)
      if err != nil {
         return nil, err
      }
      if strings.HasPrefix(p.Path, "/movies/") {
         req.URL.Path += "/watch"
      }
      req.URL.Path += p.Path
   }
   // If you request once with headers, you can request again without any
   // headers for 10 minutes, but then headers are required again
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   con := new(Content)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}

func (a Auth_ID) Playback(ref string) (*Playback, error) {
   body, err := func() ([]byte, error) {
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
      return json.MarshalIndent(s, "", " ")
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/playback-id/api/v1/playback/",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   _, nID, found := strings.Cut(ref, "--")
   if !found {
      return nil, errors.New("nid not found")
   }
   req.URL.Path += nID
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
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var play Playback
   {
      var s struct {
         Data struct {
            Playback_JSON_Data struct {
               Sources []Source
            } `json:"playbackJsonData"`
         }
      }
      err := json.NewDecoder(res.Body).Decode(&s)
      if err != nil {
         return nil, err
      }
      play.sources = s.Data.Playback_JSON_Data.Sources
   }
   play.Header = res.Header
   return &play, nil
}

func (a *Auth_ID) Refresh() error {
   req, err := http.NewRequest(
      "POST",
      "https://gw.cds.amcn.com/auth-orchestration-id/api/v1/refresh",
      nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set("Authorization", "Bearer " + a.Data.Refresh_Token)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return json.NewDecoder(res.Body).Decode(a)
}

func Unauth() (*Auth_ID, error) {
   req, err := http.NewRequest(
      "POST", "https://gw.cds.amcn.com/auth-orchestration-id/api/v1/unauth",
      nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "X-Amcn-Device-ID": {"-"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Tenant": {"amcn"},
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   auth := new(Auth_ID)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}
