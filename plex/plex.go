package plex

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (a *Anonymous) Match(web *Address) (*DiscoverMatch, error) {
   req, err := http.NewRequest(
      "", "https://discover.provider.plex.tv/library/metadata/matches", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("accept", "application/json")
   req.URL.RawQuery = url.Values{
      "url": {web.Path},
      "x-plex-token": {a.AuthToken},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var value struct {
      MediaContainer struct {
         Metadata []DiscoverMatch
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.MediaContainer.Metadata[0], nil
}

type Anonymous struct {
   AuthToken string
}

func (a *Anonymous) New() error {
   req, err := http.NewRequest(
      "POST", "https://plex.tv/api/v2/users/anonymous", nil,
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "accept": {"application/json"},
      "x-plex-product": {"Plex Mediaverse"},
      "x-plex-client-identifier": {"!"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(a)
}

func (a *Anonymous) Video(
   match *DiscoverMatch, forward string,
) (*OnDemand, error) {
   req, err := http.NewRequest("", "https://vod.provider.plex.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/library/metadata/" + match.RatingKey
   req.Header.Set("accept", "application/json")
   req.Header.Set("x-plex-token", a.AuthToken)
   if forward != "" {
      req.Header.Set("x-forwarded-for", forward)
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var value struct {
      MediaContainer struct {
         Metadata []OnDemand
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   metadata := value.MediaContainer.Metadata[0]
   for _, media := range metadata.Media {
      for _, part := range media.Part {
         part.Key.Url.RawQuery = url.Values{
            "x-plex-token": {a.AuthToken},
         }.Encode()
         if part.License != nil {
            part.License.Url.RawQuery = url.Values{
               "x-plex-drm": {"widevine"},
               "x-plex-token": {a.AuthToken},
            }.Encode()
         }
      }
   }
   return &metadata, nil
}
func (m *MediaPart) Wrap(data []byte) ([]byte, error) {
   var req http.Request
   req.Body = io.NopCloser(bytes.NewReader(data))
   req.Method = "POST"
   req.URL = m.License.Url
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type MediaPart struct {
   Key Url
   License *Url
}

type Address struct {
   Path string
}

func (a *Address) String() string {
   return a.Path
}

func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "watch.plex.tv")
   a.Path = strings.TrimPrefix(s, "/watch")
   return nil
}

func (o *OnDemand) Dash() (*MediaPart, bool) {
   for _, media := range o.Media {
      if media.Protocol == "dash" {
         for _, part := range media.Part {
            return &part, true
         }
      }
   }
   return nil, false
}

type OnDemand struct {
   Media []struct {
      Part []MediaPart
      Protocol string
   }
}

type Url struct {
   Url *url.URL
}

func (u *Url) UnmarshalText(data []byte) error {
   u.Url = &url.URL{}
   err := u.Url.UnmarshalBinary(data)
   if err != nil {
      return err
   }
   u.Url.Scheme = "https"
   u.Url.Host = "vod.provider.plex.tv"
   return nil
}

func (n Namer) Show() string {
   return n.Match.GrandparentTitle
}

func (n Namer) Title() string {
   return n.Match.Title
}

type Namer struct {
   Match *DiscoverMatch
}

type DiscoverMatch struct {
   GrandparentTitle string
   Index int
   ParentIndex int
   RatingKey string
   Title string
   Year int
}

func (n Namer) Episode() int {
   return n.Match.Index
}

func (n Namer) Season() int {
   return n.Match.ParentIndex
}

func (n Namer) Year() int {
   return n.Match.Year
}
