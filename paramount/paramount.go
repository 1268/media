package paramount

import (
   "bytes"
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

const encoding = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func pad(b []byte) []byte {
   length := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(length); length >= 1; length-- {
      b = append(b, high)
   }
   return b
}

func cms_account(id string) int64 {
   var (
      i = 0
      j = 1
   )
   for _, value := range id {
      i += strings.IndexRune(encoding, value) * j
      j *= len(encoding)
   }
   return int64(i)
}

type Number int

func (n *Number) UnmarshalText(data []byte) error {
   if len(data) >= 1 {
      v, err := strconv.Atoi(string(data))
      if err != nil {
         return err
      }
      *n = Number(v)
   }
   return nil
}

func (n Number) MarshalText() ([]byte, error) {
   return strconv.AppendInt(nil, int64(n), 10), nil
}

func (s *SessionToken) Wrap(data []byte) ([]byte, error) {
   req, err := http.NewRequest("POST", s.Url, bytes.NewReader(data))
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "authorization": {"Bearer " + s.LsSession},
      "content-type": {"application/x-protobuf"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type SessionToken struct {
   LsSession string `json:"ls_session"`
   Url string
}

func (v *VideoItem) Show() string {
   if v.SeasonNum >= 1 {
      return v.SeriesTitle
   }
   return ""
}

func (v *VideoItem) Unmarshal(data []byte) error {
   var value struct {
      Error string
      ItemList []VideoItem
   }
   err := json.Unmarshal(data, &value)
   if err != nil {
      return err
   }
   if value.Error != "" {
      return errors.New(value.Error)
   }
   if len(value.ItemList) == 0 {
      return errors.New(`"itemList":[]`)
   }
   *v = value.ItemList[0]
   return nil
}

func (v *VideoItem) Title() string {
   return v.Label
}

func (v *VideoItem) Season() int {
   return int(v.SeasonNum)
}

func (v *VideoItem) Episode() int {
   return int(v.EpisodeNum)
}

func (v *VideoItem) Year() int {
   return v.AirDateIso.Year()
}

// hard geo block
func (v *VideoItem) Mpd() string {
   b := []byte("https://link.theplatform.com/s/")
   b = append(b, v.CmsAccountId...)
   b = append(b, "/media/guid/"...)
   b = strconv.AppendInt(b, cms_account(v.CmsAccountId), 10)
   b = append(b, '/')
   b = append(b, v.ContentId...)
   b = append(b, "?assetTypes="...)
   b = append(b, v.AssetType...)
   b = append(b, "&formats=MPEG-DASH"...)
   return string(b)
}

type VideoItem struct {
   AirDateIso time.Time `json:"_airDateISO"`
   AssetType string
   CmsAccountId string
   ContentId string
   EpisodeNum Number
   Label string
   SeasonNum Number
   SeriesTitle string
}

func (a *AppToken) encode() (string, error) {
   key, err := hex.DecodeString(a.SecretKey)
   if err != nil {
      return "", err
   }
   block, err := aes.NewCipher(key)
   if err != nil {
      return "", err
   }
   src := []byte{'|'}
   src = append(src, a.AppSecret...)
   src = pad(src)
   var iv [aes.BlockSize]byte
   cipher.NewCBCEncrypter(block, iv[:]).CryptBlocks(src, src)
   var dst []byte
   dst = append(dst, 0, aes.BlockSize)
   dst = append(dst, iv[:]...)
   dst = append(dst, src...)
   return base64.StdEncoding.EncodeToString(dst), nil
}

const secret_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"

type AppToken struct {
   AppSecret string
   SecretKey string
}

// 15.0.52
var ComCbsApp = AppToken{
   AppSecret: "4fb47ec1f5c17caa",
   SecretKey: secret_key,
}

// 15.0.52
var ComCbsCa = AppToken{
   AppSecret: "e55edaeb8451f737",
   SecretKey: secret_key,
}

// must use app token and IP address for US
func (a *AppToken) Session(content_id string) (*SessionToken, error) {
   req, err := http.NewRequest("", "https://www.paramountplus.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/apps-api/v3.1/androidphone/irdeto-control")
      b.WriteString("/anonymous-session-token.json")
      return b.String()
   }()
   token, err := a.encode()
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "at": {token},
      "contentId": {content_id},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   session := &SessionToken{}
   err = json.NewDecoder(resp.Body).Decode(session)
   if err != nil {
      return nil, err
   }
   return session, nil
}

// must use app token and IP address for correct location
func (*VideoItem) Marshal(app *AppToken, cid string) ([]byte, error) {
   req, err := http.NewRequest("", "https://www.paramountplus.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/apps-api/v2.0/androidphone/video/cid/")
      b.WriteString(cid)
      b.WriteString(".json")
      return b.String()
   }()
   token, err := app.encode()
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{"at": {token}}.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   return io.ReadAll(resp.Body)
}
