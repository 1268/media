package amc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "net/url"
   "os"
   "strings"
)

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

type playback_request struct {
   Ad_Tags struct {
      Lat int `json:"lat"`
      Mode string `json:"mode"`
      PPID int `json:"ppid"`
      Player_Height int `json:"playerHeight"`
      Player_Width int `json:"playerWidth"`
      URL string `json:"url"`
   } `json:"adtags"`
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

type Content struct {
   Data	struct {
      Children []struct {
         Properties json.RawMessage
         Type string
      }
   }
}

func (c Content) Video_Player() (*Video_Player, error) {
   for _, child := range c.Data.Children {
      if child.Type == "video-player-ap" {
         vid := new(Video_Player)
         err := json.Unmarshal(child.Properties, vid)
         if err != nil {
            return nil, err
         }
         return vid, nil
      }
   }
   return nil, errors.New("video-player-ap not present")
}

type Video_Player struct {
   Content_Type string `json:"contentType"`
   Current_Video struct {
      Meta struct {
         Show_Title string `json:"showTitle"`
         
         Title string
         Airdate string // 1996-01-01T00:00:00.000Z
      }
   } `json:"currentVideo"`
}

const sep_big = " - "

func (v Video_Player) Name() (string, error) {
   year, _, found := strings.Cut(v.Current_Video.Meta.Airdate, "-")
   if !found {
      return "", errors.New("year not found")
   }
   var b strings.Builder
   b.WriteString(v.Current_Video.Meta.Title)
   if v.Content_Type == "movie" {
      b.WriteString(sep_big)
      b.WriteString(year)
   }
   return b.String(), nil
}
