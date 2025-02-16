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

func (a *Address) Set(data string) error {
   for {
      var (
         key string
         ok  bool
      )
      key, data, ok = strings.Cut(data, "/")
      if !ok {
         return nil
      }
      switch key {
      case "movies":
         (*a)[0] = data
      case "series":
         (*a)[0], data, ok = strings.Cut(data, "/")
         if !ok {
            return errors.New("episode")
         }
      case "episode":
         (*a)[1] = data
      }
   }
}

type Address [2]string

func (v Vod) Clips() (*Clips, error) {
   req, err := http.NewRequest("", "https://api.pluto.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var data strings.Builder
      data.WriteString("/v2/episodes/")
      if v.Id != "" {
         data.WriteString(v.Id)
      } else {
         data.WriteString(v.Episode)
      }
      data.WriteString("/clips.json")
      return data.String()
   }()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var clips1 []Clips
   err = json.NewDecoder(resp.Body).Decode(&clips1)
   if err != nil {
      return nil, err
   }
   return &clips1[0], nil
}

func (a Address) Vod(forward string) (*Vod, error) {
   req, _ := http.NewRequest("", "https://boot.pluto.tv/v4/start", nil)
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
      Vod []Vod
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
   for _, season1 := range demand.Seasons {
      for _, episode := range season1.Episodes {
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

type Client struct{}

func (Client) License(data []byte) ([]byte, error) {
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

type Vod struct {
   Episode string `json:"_id"`
   Id      string
   Name    string
   Seasons []struct {
      Episodes []*Vod
   }
   Slug    string
}

type Clips struct {
   Sources []struct {
      File File
      Type string
   }
}

func (c *Clips) Dash() (*File, bool) {
   for _, source := range c.Sources {
      if source.Type == "DASH" {
         return &source.File, true
      }
   }
   return nil, false
}

type File [1]url.URL

// these return a valid response body, but response status is "403 OK":
// http://siloh-fs.plutotv.net
// http://siloh-ns1.plutotv.net
// https://siloh-fs.plutotv.net
// https://siloh-ns1.plutotv.net
func (f *File) UnmarshalText(data []byte) error {
   err := (*f)[0].UnmarshalBinary(data)
   if err != nil {
      return err
   }
   (*f)[0].Scheme = "http"
   (*f)[0].Host = "silo-hybrik.pluto.tv.s3.amazonaws.com"
   return nil
}

// The Request's URL and Header fields must be initialized
func (f File) Mpd() (*http.Response, error) {
   var req http.Request
   req.Header = http.Header{}
   req.URL = &f[0]
   return http.DefaultClient.Do(&req)
}
