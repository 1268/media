package roku

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/json"
   "net/url"
   "strings"
   "time"
)

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

func New_Content(id string) (*Content, error) {
   homescreen := func() string {
      include := []string{
         "series.title",
         "seasonNumber",
         "episodeNumber",
         "title",
         "releaseDate",
         // this needs to be exactly as is, otherwise size blows up
         "series.seasons.episodes.viewOptions\u2008",
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
      return url.PathEscape(expand.String())
   }
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "therokuchannel.roku.com",
      Path: "/api/v2/homescreen/content/" + homescreen(),
   })
   res, err := http.Default_Client.Do(req)
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

