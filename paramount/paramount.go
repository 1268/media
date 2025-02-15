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
)

const secret_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"

func pad(data []byte) []byte {
   length := aes.BlockSize - len(data) % aes.BlockSize
   for high := byte(length); length >= 1; length-- {
      data = append(data, high)
   }
   return data
}

const encoding = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

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
   AppSecret string
   SecretKey string
}

type SessionToken struct {
   LsSession string `json:"ls_session"`
   Url string
}

type Item struct {
   AssetType string
   CmsAccountId string
   ContentId string
}

// must use app token and IP address for correct location
func (Item) Marshal(token *AppToken, cid string) ([]byte, error) {
   req, _ := http.NewRequest("", "https://www.paramountplus.com", nil)
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/apps-api/v2.0/androidphone/video/cid/")
      b.WriteString(cid)
      b.WriteString(".json")
      return b.String()
   }()
   token1, err := token.encode()
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{"at": {token1}}.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var data strings.Builder
      resp.Write(&data)
      return nil, errors.New(data.String())
   }
   return io.ReadAll(resp.Body)
}

func (i *Item) Unmarshal(data []byte) error {
   var value struct {
      Error string
      ItemList []Item
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
   *i = value.ItemList[0]
   return nil
}

// hard geo block
func (i *Item) Mpd() string {
   b := []byte("https://link.theplatform.com/s/")
   b = append(b, i.CmsAccountId...)
   b = append(b, "/media/guid/"...)
   b = strconv.AppendInt(b, cms_account(i.CmsAccountId), 10)
   b = append(b, '/')
   b = append(b, i.ContentId...)
   b = append(b, "?assetTypes="...)
   b = append(b, i.AssetType...)
   b = append(b, "&formats=MPEG-DASH"...)
   return string(b)
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

func (a *AppToken) encode() (string, error) {
   key, err := hex.DecodeString(a.SecretKey)
   if err != nil {
      return "", err
   }
   block, err := aes.NewCipher(key)
   if err != nil {
      return "", err
   }
   var iv [aes.BlockSize]byte
   data := []byte{'|'}
   data = append(data, a.AppSecret...)
   data = pad(data)
   cipher.NewCBCEncrypter(block, iv[:]).CryptBlocks(data, data)
   data1 := []byte{0, aes.BlockSize}
   data1 = append(data1, iv[:]...)
   data1 = append(data1, data...)
   return base64.StdEncoding.EncodeToString(data1), nil
}

func (s *SessionToken) License(data []byte) ([]byte, error) {
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

// must use app token and IP address for US
func (a *AppToken) Session(content_id string) (*SessionToken, error) {
   req, _ := http.NewRequest("", "https://www.paramountplus.com", nil)
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
      var data strings.Builder
      resp.Write(&data)
      return nil, errors.New(data.String())
   }
   session := &SessionToken{}
   err = json.NewDecoder(resp.Body).Decode(session)
   if err != nil {
      return nil, err
   }
   return session, nil
}
