package itv

import (
   "bytes"
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
   "net/http/cookiejar"
   "net/url"
   "strings"
)

type Title struct {
   LatestAvailableVersion struct {
      PlaylistUrl string
   }
}

func (i LegacyId) Title() (*Title, error) {
   req, _ := http.NewRequest(
      "", "https://content-inventory.prd.oasvc.itv.com/discovery", nil,
   )
   req.URL.RawQuery = "query=" + url.QueryEscape(
      fmt.Sprintf(graphql_compact(query_discovery), i),
   )
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var value struct {
      Data struct {
         Titles []Title
      }
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   if err := value.Errors; len(err) >= 1 {
      return nil, errors.New(err[0].Message)
   }
   return &value.Data.Titles[0], nil
}

type LegacyId [3]string

type MediaFile struct {
   Href Href
   KeyServiceUrl string
   Resolution string
}

type Playlist struct {
   Playlist struct {
      Video struct {
         MediaFiles []MediaFile
      }
   }
}

func (m *MediaFile) License(data []byte) (*http.Response, error) {
   return http.Post(
      m.KeyServiceUrl, "application/x-protobuf", bytes.NewReader(data),
   )
}

func (h Href) Mpd() (*http.Response, error) {
   var err error
   http.DefaultClient.Jar, err = cookiejar.New(nil)
   if err != nil {
      return nil, err
   }
   return http.Get(h[0])
}

func (h *Href) UnmarshalText(data []byte) error {
   (*h)[0] = strings.Replace(string(data), "itvpnpctv", "itvpnpdotcom", 1)
   return nil
}

type Href [1]string

const query_discovery = `
{
   titles(filter: {
      legacyId: %q
   }) {
      latestAvailableVersion {
         playlistUrl
      }
   }
}
`

func (i LegacyId) String() string {
   var data strings.Builder
   for key, value := range i {
      if value != "" {
         if key >= 1 {
            data.WriteByte('/')
         }
         data.WriteString(value)
      }
   }
   return data.String()
}

func (i *LegacyId) Set(text string) error {
   var found bool
   (*i)[0], text, found = strings.Cut(text, "a")
   if !found {
      return errors.New(`"a" not found`)
   }
   (*i)[1], (*i)[2], found = strings.Cut(text, "a")
   if !found {
      (*i)[2] = "0001"
   }
   return nil
}

func (p *Playlist) Resolution1080() (*MediaFile, bool) {
   for _, file := range p.Playlist.Video.MediaFiles {
      if file.Resolution == "1080" {
         return &file, true
      }
   }
   return nil, false
}

// hard geo block
func (t *Title) Playlist() (*Playlist, error) {
   data, err := json.Marshal(map[string]any{
      "client": map[string]string{
         "id": "browser",
      },
      "variantAvailability": map[string]any{
         "drm": map[string]string{
            "maxSupported": "L3",
            "system": "widevine",
         },
         "featureset": []string{ // need all these to get 720p
            "hd",
            "mpeg-dash",
            "single-track",
            "widevine",
         },
         "platformTag": "ctv", // 1080p
      },
   })
   req, err := http.NewRequest(
      "POST", t.LatestAvailableVersion.PlaylistUrl, bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("accept", "application/vnd.itv.vod.playlist.v4+json")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   play := &Playlist{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(data string) string {
   return strings.Join(strings.Fields(data), " ")
}
