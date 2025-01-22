package nbc

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

func (d *DrmProxy) New() {
   d.Time = fmt.Sprint(time.Now().UnixMilli())
   d.Hash = func() string {
      hash := hmac.New(sha256.New, []byte(drm_proxy_secret))
      fmt.Fprint(hash, d.Time, "widevine")
      return fmt.Sprintf("%x", hash.Sum(nil))
   }()
}

type DrmProxy struct {
   Hash string
   Time string
}

const drm_proxy_secret = "Whn8QFuLFM7Heiz6fYCYga7cYPM8ARe6"

func (d *DrmProxy) Wrap(data []byte) ([]byte, error) {
   req, err := http.NewRequest(
      "POST", "https://drmproxy.digitalsvc.apps.nbcuni.com",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/drm-proxy/license/widevine"
   req.URL.RawQuery = url.Values{
      "device": {"web"},
      "hash": {d.Hash},
      "time": {d.Time},
   }.Encode()
   req.Header.Set("content-type", "application/octet-stream")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type OnDemand struct {
   PlaybackUrl string
}

type Metadata struct {
   AirDate time.Time
   EpisodeNumber int `json:",string"`
   MovieShortTitle string
   MpxAccountId int64 `json:",string"`
   MpxGuid int64 `json:",string"`
   ProgrammingType string
   SeasonNumber int `json:",string"`
   SecondaryTitle string
   SeriesShortTitle string
}

func (m *Metadata) OnDemand() (*OnDemand, error) {
   req, err := http.NewRequest("", "https://lemonade.nbc.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/v1/vod/")
      b = strconv.AppendInt(b, m.MpxAccountId, 10)
      b = append(b, '/')
      b = strconv.AppendInt(b, m.MpxGuid, 10)
      return string(b)
   }()
   req.URL.RawQuery = url.Values{
      "platform": {"web"},
      "programmingType": {m.ProgrammingType},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   video := &OnDemand{}
   err = json.NewDecoder(resp.Body).Decode(video)
   if err != nil {
      return nil, err
   }
   return video, nil
}

func (m *Metadata) New(guid int) error {
   var page page_request
   page.Query = graphql_compact(bonanza_page)
   page.Variables.App = "nbc"
   page.Variables.Name = strconv.Itoa(guid)
   page.Variables.OneApp = true
   page.Variables.Platform = "android"
   page.Variables.Type = "VIDEO"
   data, err := json.Marshal(page)
   if err != nil {
      return err
   }
   resp, err := http.Post(
      "https://friendship.nbc.co/v2/graphql", "application/json",
      bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   var value struct {
      Data struct {
         BonanzaPage struct {
            Metadata Metadata
         }
      }
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return err
   }
   if v := value.Errors; len(v) >= 1 {
      return errors.New(v[0].Message)
   }
   *m = value.Data.BonanzaPage.Metadata
   return nil
}

// NO ANONYMOUS QUERY
const bonanza_page = `
query bonanzaPage(
   $app: NBCUBrands!
   $name: String!
   $oneApp: Boolean
   $platform: SupportedPlatforms!
   $type: EntityPageType!
   $userId: String!
) {
   bonanzaPage(
      app: $app
      name: $name
      oneApp: $oneApp
      platform: $platform
      type: $type
      userId: $userId
   ) {
      metadata {
         ... on VideoPageData {
            airDate
            episodeNumber
            movieShortTitle
            mpxAccountId
            mpxGuid
            programmingType
            seasonNumber
            secondaryTitle
            seriesShortTitle
         }
      }
   }
}
`

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(s string) string {
   field := strings.Fields(s)
   return strings.Join(field, " ")
}

type page_request struct {
   Query string `json:"query"`
   Variables struct {
      App string `json:"app"` // String cannot represent a non string value
      Name string `json:"name"`
      OneApp bool `json:"oneApp"`
      Platform string `json:"platform"`
      Type string `json:"type"` // can be empty
      UserId string `json:"userId"`
   } `json:"variables"`
}
