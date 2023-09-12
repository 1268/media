package roku

import (
   "154.pages.dev/encoding/json"
   "errors"
   "io"
   "net/http"
   "time"
)

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

func (Playback) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}

func New_Cross_Site() (*Cross_Site, error) {
   // this has smaller body than www.roku.com
   res, err := http.Get("https://therokuchannel.roku.com")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var site Cross_Site
   site.cookies = res.Cookies()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   _, text = json.Cut(text, []byte("\tcsrf:"), nil)
   if err := json.Unmarshal(text, &site.token); err != nil {
      return nil, err
   }
   return &site, nil
}

type Cross_Site struct {
   cookies []*http.Cookie
   token string
}

func (c Cross_Site) csrf() *http.Cookie {
   for _, cookie := range c.cookies {
      if cookie.Name == "_csrf" {
         return cookie
      }
   }
   return nil
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
