package roku

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func (c Cross_Site) Playback(id string) (*Playback, error) {
   body := func(r *http.Request) error {
      m := map[string]string{
         "mediaFormat": "mpeg-dash",
         "providerId": "rokuavod",
         "rokuId": id,
      }
      b, err := json.MarshalIndent(m, "", " ")
      if err != nil {
         return err
      }
      r.Body_Bytes(b)
      return nil
   }
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "therokuchannel.roku.com",
      Path: "/api/v3/playback",
   })
   // we could use Request.AddCookie, but we would need to call it after this,
   // otherwise it would be clobbered
   req.Header = http.Header{
      "CSRF-Token": {c.token},
      "Content-Type": {"application/json"},
      "Cookie": {c.cookie.Raw},
   }
   err := body(req)
   if err != nil {
      return nil, err
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(Playback)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

type Cross_Site struct {
   cookie *http.Cookie // has own String method
   token string
}
func New_Content(id string) (*Content, error) {
   req, err := http.NewRequest(
      "GET", "https://therokuchannel.roku.com/api/v2/homescreen/content", nil,
   )
   if err != nil {
      return nil, err
   }
   {
      include := []string{
         "episodeNumber",
         "releaseDate",
         "seasonNumber",
         "series.title",
         "title",
         "viewOptions",
      }
      expand := url.URL{
         Scheme: "https",
         Host: "content.sr.roku.com",
         Path: "/content/v1/roku-trc/" + id,
         RawQuery: url.Values{
            "expand": {"series"},
            "include": {strings.Join(include, ",")},
         }.Encode(),
      }
      homescreen := url.PathEscape(expand.String())
      req.URL = req.URL.JoinPath(homescreen)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var con Content
   if err := json.NewDecoder(res.Body).Decode(&con.s); err != nil {
      return nil, err
   }
   return &con, nil
}

func (c Content) Title() string {
   return c.s.Title
}

func (c Content) Series() string {
   return c.s.Series.Title
}

func (c Content) Season() (int64, error) {
   return c.s.Season_Number, nil
}

func (c Content) Episode() (int64, error) {
   return c.s.Episode_Number, nil
}

func (c Content) Date() (time.Time, error) {
   return time.Parse(time.RFC3339, c.s.Release_Date)
}

type Content struct {
   s struct {
      Series struct {
         Title string
      }
      Season_Number int64 `json:"seasonNumber,string"`
      Episode_Number int64 `json:"episodeNumber,string"`
      Title string
      Release_Date string `json:"releaseDate"` // 2007-01-01T000000Z
      View_Options []struct {
         Media struct {
            Videos []Video
         }
      } `json:"viewOptions"`
   }
}

type Video struct {
   DRM_Authentication *struct{} `json:"drmAuthentication"`
   URL string
   Video_Type string `json:"videoType"`
}

func (c Content) DASH() *Video {
   for _, option := range c.s.View_Options {
      for _, vid := range option.Media.Videos {
         if vid.Video_Type == "DASH" {
            return &vid
         }
      }
   }
   return nil
}

type Playback struct {
   DRM struct {
      Widevine struct {
         License_Server string `json:"licenseServer"`
      }
   }
}

func (p Playback) Request_URL() string {
   return p.DRM.Widevine.License_Server
}

func (Playback) Request_Header() http.Header {
   return nil
}

func (Playback) Request_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (Playback) Response_Body(buf []byte) ([]byte, error) {
   return buf, nil
}
