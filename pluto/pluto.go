package pluto

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (a Address) Video(forward string) (*OnDemand, error) {
   req, err := http.NewRequest("", "https://boot.pluto.tv/v4/start", nil)
   if err != nil {
      return nil, err
   }
   if forward != "" {
      req.Header.Set("x-forwarded-for", forward)
   }
   req.URL.RawQuery = url.Values{
      "appName":           {"web"},
      "appVersion":        {"9"},
      "clientID":          {"9"},
      "clientModelNumber": {"9"},
      "drmCapabilities":   {"widevine:L3"},
      "seriesIDs":         {a[0]},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var value struct {
      Vod []OnDemand
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   demand := value.Vod[0]
   if demand.Slug != a[0] {
      if demand.Id != a[0] {
         return nil, errors.New(demand.Slug)
      }
   }
   for _, season := range demand.Seasons {
      season.show = &demand
      for _, episode := range season.Episodes {
         episode.season = season
         if episode.Episode == a[1] {
            return episode, nil
         }
         if episode.Slug == a[1] {
            return episode, nil
         }
      }
   }
   return &demand, nil
}

type VideoSeason struct {
   Episodes []*OnDemand
   show     *OnDemand
}

type OnDemand struct {
   Episode string `json:"_id"`
   Id      string
   Name    string
   Seasons []*VideoSeason
   Slug    string
   season  *VideoSeason
}

func (o OnDemand) Clip() (*EpisodeClip, error) {
   req, err := http.NewRequest("", "https://api.pluto.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v2/episodes/")
      if o.Id != "" {
         b.WriteString(o.Id)
      } else {
         b.WriteString(o.Episode)
      }
      b.WriteString("/clips.json")
      return b.String()
   }()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var clips []EpisodeClip
   err = json.NewDecoder(resp.Body).Decode(&clips)
   if err != nil {
      return nil, err
   }
   return &clips[0], nil
}

type EpisodeClip struct {
   Sources []struct {
      File Url
      Type string
   }
}

func (e *EpisodeClip) Dash() (*url.URL, bool) {
   for _, source := range e.Sources {
      if source.Type == "DASH" {
         return &source.File.Url, true
      }
   }
   return nil, false
}

func (a Address) String() string {
   var data strings.Builder
   if a[0] != "" {
      if a[1] != "" {
         data.WriteString("series/")
         data.WriteString(a[0])
         data.WriteString("/episode/")
         data.WriteString(a[1])
      } else {
         data.WriteString("movies/")
         data.WriteString(a[0])
      }
   }
   return data.String()
}

func (Wrapper) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      "https://service-concierge.clusters.pluto.tv/v1/wv/alt",
      "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(string(data))
   }
   return data, nil
}

type Wrapper struct{}

type Address [2]string

func (a *Address) Set(text string) error {
   for {
      var (
         key string
         ok  bool
      )
      key, text, ok = strings.Cut(text, "/")
      if !ok {
         return nil
      }
      switch key {
      case "movies":
         (*a)[0] = text
      case "series":
         (*a)[0], text, ok = strings.Cut(text, "/")
         if !ok {
            return errors.New("episode")
         }
      case "episode":
         (*a)[1] = text
      }
   }
}

var Base = []FileBase{
   {"http", "silo-hybrik.pluto.tv.s3.amazonaws.com", "200 OK"},
   {"http", "siloh-fs.plutotv.net", "403 OK"},
   {"http", "siloh-ns1.plutotv.net", "403 OK"},
   {"https", "siloh-fs.plutotv.net", "403 OK"},
   {"https", "siloh-ns1.plutotv.net", "403 OK"},
}

type FileBase struct {
   Scheme string
   Host   string
   Status string
}

type Url struct {
   Url url.URL
}

func (u *Url) UnmarshalText(text []byte) error {
   return u.Url.UnmarshalBinary(text)
}
