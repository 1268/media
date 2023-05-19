package roku

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/json"
   "net/url"
   "strings"
)

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
