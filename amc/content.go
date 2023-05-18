package amc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "net/url"
   "strings"
   "time"
)

type Content struct {
   Data	struct {
      Children []struct {
         Properties json.RawMessage
         Type string
      }
   }
}

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

func (v Video) Season() int64 {
   return v.Meta.Season
}

func (v Video) Episode() int64 {
   return v.Meta.Episode_Number
}

func (v Video) Title() string {
   return v.Text.Title
}

func (v Video) Date() (time.Time, error) {
   return time.Parse(time.RFC3339, v.Meta.Airdate)
}
