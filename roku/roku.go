package roku

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/json"
   "io"
   "net/url"
   "strings"
   "time"
)

func (c Cross_Site) Playback(id string) (*Playback, error) {
   body := map[string]string{
      "mediaFormat": "mpeg-dash",
      "providerId": "rokuavod",
      "rokuId": id,
   }
   raw, err := json.MarshalIndent(body, "", " ")
   if err != nil {
      return nil, err
   }
   req := http.Post()
   req.AddCookie(c.cookie)
   req.Body_Bytes(raw)
   req.Header = http.Header{
      "CSRF-Token": {c.token},
      "Content-Type": {"application/json"},
   }
   req.URL.Host = "therokuchannel.roku.com"
   req.URL.Path = "/api/v3/playback"
   req.URL.Scheme = "https"
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

func (c Content) DASH() *Video {
   for _, opt := range c.View_Options {
      for _, vid := range opt.Media.Videos {
         if vid.Video_Type == "DASH" {
            return &vid
         }
      }
   }
   return nil
}
func (c Content) Name() string {
   var b strings.Builder
   if c.Meta.Media_Type == "episode" {
      b.WriteString(c.Series.Title)
      b.WriteByte('-')
      b.WriteString(c.Season_Number)
      b.WriteByte('-')
      b.WriteString(c.Episode_Number)
      b.WriteByte('-')
   }
   b.WriteString(c.Title)
   b.WriteByte('-')
   year, _, _ := strings.Cut(c.Release_Date, "-")
   b.WriteString(year)
   return b.String()
}

func (c Content) String() string {
   var b strings.Builder
   b.WriteString("ID: ")
   b.WriteString(c.Meta.ID)
   b.WriteString("\nType: ")
   b.WriteString(c.Meta.Media_Type)
   b.WriteString("\nTitle: ")
   b.WriteString(c.Title)
   if c.Meta.Media_Type == "episode" {
      b.WriteString("\nSeries: ")
      b.WriteString(c.Series.Title)
      b.WriteString("\nSeason: ")
      b.WriteString(c.Season_Number)
      b.WriteString("\nEpisode: ")
      b.WriteString(c.Episode_Number)
   }
   b.WriteString("\nDate: ")
   b.WriteString(c.Release_Date)
   b.WriteString("\nDuration: ")
   b.WriteString(c.Duration().String())
   return b.String()
}

func (c Content) Duration() time.Duration {
   return time.Duration(c.Run_Time_Seconds) * time.Second
}

type Content struct {
   Episode_Number string `json:"episodeNumber"`
   Meta struct {
      ID string
      Media_Type string `json:"mediaType"`
   }
   Release_Date string `json:"releaseDate"` // 2007-01-01T000000Z
   Run_Time_Seconds int64 `json:"runTimeSeconds"`
   Season_Number string `json:"seasonNumber"`
   Series struct {
      Title string
   }
   Title string
   View_Options []struct {
      License string
      Media struct {
         Videos []Video
      }
   } `json:"viewOptions"`
}

func New_Content(id string) (*Content, error) {
   include := []string{
      "episodeNumber",
      "releaseDate",
      "runTimeSeconds",
      "seasonNumber",
      // this needs to be exactly as is, otherwise size blows up
      "series.seasons.episodes.viewOptions\u2008",
      "series.title",
      "title",
      "viewOptions",
   }
   var expand url.URL
   expand.Scheme = "https"
   expand.Host = "content.sr.roku.com"
   expand.Path = "/content/v1/roku-trc/" + id
   expand.RawQuery = url.Values{
      "expand": {"series"},
      "include": {strings.Join(include, ",")},
   }.Encode()
   var home strings.Builder
   home.WriteString("https://therokuchannel.roku.com")
   home.WriteString("/api/v2/homescreen/content/")
   home.WriteString(url.PathEscape(expand.String()))
   res, err := http.Default_Client.Get(home.String())
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   screen := new(Content)
   if err := json.NewDecoder(res.Body).Decode(screen); err != nil {
      return nil, err
   }
   return screen, nil
}

type Video struct {
   DRM_Authentication *struct{} `json:"drmAuthentication"`
   URL string
   Video_Type string `json:"videoType"`
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

type Cross_Site struct {
   cookie *http.Cookie // has own String method
   token string
}

func New_Cross_Site() (*Cross_Site, error) {
   // this has smaller body than www.roku.com
   req := http.Get()
   req.URL.Scheme = "https"
   req.URL.Host = "therokuchannel.roku.com"
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var site Cross_Site
   for _, cook := range res.Cookies() {
      if cook.Name == "_csrf" {
         site.cookie = cook
      }
   }
   var scan json.Scanner
   scan.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Sep = []byte("\tcsrf:")
   scan.Scan()
   scan.Sep = nil
   if err := scan.Decode(&site.token); err != nil {
      return nil, err
   }
   return &site, nil
}

