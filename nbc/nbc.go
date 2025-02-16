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

func (m *Metadata) Vod() (*Vod, error) {
   req, _ := http.NewRequest("", "https://lemonade.nbc.com", nil)
   req.URL.Path = func() string {
      b := []byte("/v1/vod/")
      b = strconv.AppendInt(b, m.MpxAccountId, 10)
      b = append(b, '/')
      b = strconv.AppendInt(b, m.MpxGuid, 10)
      return string(b)
   }()
   req.URL.RawQuery = url.Values{
      "platform":        {"web"},
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
   video := &Vod{}
   err = json.NewDecoder(resp.Body).Decode(video)
   if err != nil {
      return nil, err
   }
   return video, nil
}

type Metadata struct {
   MpxAccountId     int64 `json:",string"`
   MpxGuid          int64 `json:",string"`
   ProgrammingType  string
}

func (m *Metadata) New(guid int) error {
   data, err := json.Marshal(map[string]any{
      "query": graphql_compact(bonanza_page),
      "variables": map[string]any{
         "app": "nbc",
         "name": strconv.Itoa(guid),
         "oneApp": true,
         "platform": "android",
         "type": "VIDEO",
         "userId": "",
      },
   })
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
   if err := value.Errors; len(err) >= 1 {
      return errors.New(err[0].Message)
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
            mpxAccountId
            mpxGuid
            programmingType
         }
      }
   }
}
`

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(data string) string {
   return strings.Join(strings.Fields(data), " ")
}

const drm_proxy_secret = "Whn8QFuLFM7Heiz6fYCYga7cYPM8ARe6"

func (v Vod) Mpd() (*http.Response, error) {
   return http.Get(v.PlaybackUrl)
}

type Vod struct {
   PlaybackUrl string
}

type Client struct {
   Time string
   Hash string
}

func (c *Client) New() {
   c.Time = fmt.Sprint(time.Now().UnixMilli())
   c.Hash = func() string {
      hash := hmac.New(sha256.New, []byte(drm_proxy_secret))
      fmt.Fprint(hash, c.Time, "widevine")
      return fmt.Sprintf("%x", hash.Sum(nil))
   }()
}

func (c *Client) License(data []byte) ([]byte, error) {
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
      "hash":   {c.Hash},
      "time":   {c.Time},
   }.Encode()
   req.Header.Set("content-type", "application/octet-stream")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}
