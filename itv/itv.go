package itv

import (
   "bytes"
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (h *Href) UnmarshalText(data []byte) error {
   h.S = strings.Replace(string(data), "itvpnpctv", "itvpnpdotcom", 1)
   return nil
}

type Href struct {
   S string
}

func (p *Playlist) Resolution1080() (*MediaFile, bool) {
   for _, file := range p.Playlist.Video.MediaFiles {
      if file.Resolution == "1080" {
         return &file, true
      }
   }
   return nil, false
}

type Playlist struct {
   Playlist struct {
      Video struct {
         MediaFiles []MediaFile
      }
   }
}

func (m *MediaFile) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      m.KeyServiceUrl, "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type MediaFile struct {
   Href Href
   KeyServiceUrl string
   Resolution string
}

// hard geo block
func (d *DiscoveryTitle) Playlist() (*Playlist, error) {
   var value struct {
      Client struct {
         Id string `json:"id"`
      } `json:"client"`
      VariantAvailability struct {
         Drm         struct {
            MaxSupported string `json:"maxSupported"`
            System       string `json:"system"`
         } `json:"drm"`
         FeatureSet  []string `json:"featureset"`
         PlatformTag string   `json:"platformTag"`
      } `json:"variantAvailability"`
   }
   value.Client.Id = "browser"
   value.VariantAvailability.Drm.MaxSupported = "L3"
   value.VariantAvailability.Drm.System = "widevine"
   // need all these to get 720:
   value.VariantAvailability.FeatureSet = []string{
      "hd",
      "mpeg-dash",
      "single-track",
      "widevine",
   }
   // 1080p
   value.VariantAvailability.PlatformTag = "ctv"
   data, err := json.MarshalIndent(value, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", d.LatestAvailableVersion.PlaylistUrl, bytes.NewReader(data),
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

func (i LegacyId) Discovery() (*DiscoveryTitle, error) {
   req, err := http.NewRequest(
      "", "https://content-inventory.prd.oasvc.itv.com/discovery", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "query": {fmt.Sprintf(graphql_compact(query_discovery), i)},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var value struct {
      Data struct {
         Titles []DiscoveryTitle
      }
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   if v := value.Errors; len(v) >= 1 {
      return nil, errors.New(v[0].Message)
   }
   return &value.Data.Titles[0], nil
}

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(data string) string {
   field := strings.Fields(data)
   return strings.Join(field, " ")
}

const query_discovery = `
{
   titles(filter: {
      legacyId: %q
   }) {
      ... on Episode {
         seriesNumber
         episodeNumber
      }
      ... on Film {
         productionYear
      }
      brand {
         title
      }
      latestAvailableVersion {
         playlistUrl
      }
      title
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

type LegacyId [3]string

type DiscoveryTitle struct {
   LatestAvailableVersion struct {
      PlaylistUrl string
   }
   Brand *struct {
      Title string
   }
   EpisodeNumber int
   ProductionYear int
   SeriesNumber int
   Title string
}
