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

// must use app token and IP address for US
func (a AppToken) Session(content_id string) (*SessionToken, error) {
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
   a.Values.Set("contentId", content_id)
   req.URL.RawQuery = a.Values.Encode()
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

func (a *AppToken) New(app_secret string) error {
   key, err := hex.DecodeString(secret_key)
   if err != nil {
      return err
   }
   block, err := aes.NewCipher(key)
   if err != nil {
      return err
   }
   var src []byte
   src = append(src, '|')
   src = append(src, app_secret...)
   src = pad(src)
   var iv [aes.BlockSize]byte
   cipher.NewCBCEncrypter(block, iv[:]).CryptBlocks(src, src)
   var dst []byte
   dst = append(dst, 0, aes.BlockSize)
   dst = append(dst, iv[:]...)
   dst = append(dst, src...)
   a.Values = url.Values{
      "at": {base64.StdEncoding.EncodeToString(dst)},
   }
   return nil
}

// 15.0.28
func (a *AppToken) ComCbsApp() error {
   return a.New("a624d7b175f5626b")
}

// 15.0.28
func (a *AppToken) ComCbsCa() error {
   return a.New("c0b1d5d6ed27a3f6")
}

const secret_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"

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

type AppToken struct {
   Values url.Values
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

// hard geo block
func (v *VideoItem) Mpd() string {
   data := []byte("https://link.theplatform.com/s/")
   data = append(data, v.CmsAccountId...)
   data = append(data, "/media/guid/"...)
   data = strconv.AppendInt(data, cms_account(v.CmsAccountId), 10)
   data = append(data, '/')
   data = append(data, v.ContentId...)
   data = append(data, "?assetTypes="...)
   data = append(data, v.asset_type()...)
   data = append(data, "&formats=MPEG-DASH"...)
   return string(data)
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

// must use app token and IP address for correct location
func (*VideoItem) Marshal(token AppToken, cid string) ([]byte, error) {
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
   req.URL.RawQuery = token.Values.Encode()
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

type VideoItem struct {
   AirDateIso time.Time `json:"_airDateISO"`
   CmsAccountId string
   ContentId string
   EpisodeNum Number
   Label string
   RegionalRatings []struct {
      Region string
   }
   SeasonNum Number
   SeriesTitle string
}

func (v *VideoItem) asset_type() string {
   for _, rating := range v.RegionalRatings {
      if rating.Region == "GB" {
         return "DASH_CENC_PRECON"
      }
   }
   return "DASH_CENC"
}
