package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
)

type Session struct {
   URL string
   LS_Session string
}

func (Session) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (s Session) Request_Header() http.Header {
   h := make(http.Header)
   h.Set("Authorization", "Bearer " + s.LS_Session)
   return h
}

func (s Session) Request_URL() string {
   return s.URL
}

func (Session) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (at App_Token) Session(content_ID string) (*Session, error) {
   req, err := http.NewRequest("GET", "https://www.paramountplus.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/apps-api/v3.0/androidphone/irdeto-control/anonymous-session-token.json"
   req.URL.RawQuery = url.Values{
      // this needs to be encoded
      "at": {at.value},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   sess := new(Session)
   if err := json.NewDecoder(res.Body).Decode(sess); err != nil {
      return nil, err
   }
   sess.URL += content_ID
   return sess, nil
}

type app_details struct {
   version string
   code int
}

// com.cbs.app
var app_secrets = map[app_details]string{
   { "4.8.06",   1648603}: "a958002817953588",
   { "7.3.58", 210735833}: "4be3d46aecbcd26d",
   { "8.0.00", 210800061}: "5d1d865f929d3daa",
   { "8.0.10", 210801034}: "685c401ff9a4a2d9",
   { "8.0.12", 210801228}: "20021bb2eda91db4",
   { "8.0.14", 210801417}: "f3577b860abfa76d",
   { "8.0.16", 210801619}: "79e71194ad8b97d4",
   { "8.0.20", 210802025}: "003ff1f049feb54a",
   { "8.0.22", 210802235}: "36a841291cfecc4e",
   { "8.0.24", 210802415}: "2b7feb264967d94f",
   { "8.0.26", 210802628}: "860c7062bb69759d",
   { "8.0.28", 210802834}: "1fc4f2e07173b30c",
   { "8.0.30", 210803027}: "90946a66385ceeb5",
   { "8.0.32", 210803227}: "262d30953b16032b",
   { "8.0.34", 210803426}: "0f84a8e9f62594ad",
   { "8.0.36", 210803633}: "6dfcc58b09fca975",
   { "8.0.38", 210803826}: "423187842fdd7eac",
   { "8.0.40", 210804027}: "d08c12908070b2ac",
   { "8.0.42", 210804235}: "c824c27d68eb9fc3",
   { "8.0.44", 210804421}: "9c5b3eda87e38402",
   { "8.0.46", 210804638}: "88065c1d30bc15ed",
   { "8.0.48", 210804826}: "00a7ea18c54d674c",
   { "8.0.50", 210805068}: "2e6cd61ba21dd0c4",
   { "8.0.52", 210805212}: "5fcf8c6937dba442",
   { "8.0.54", 210805434}: "6c70b33080758409",
   { "8.1.10", 210811006}: "8150802ffdeffaf0",
   { "8.1.12", 210811230}: "3395c01da67a358b",
   { "8.1.14", 210811423}: "acacc94f4214ddc1",
   { "8.1.16", 210811625}: "add603b54be2fc3c",
   { "8.1.18", 210811824}: "1705732089ff4d60",
   { "8.1.20", 210812024}: "817774cbafb2b797",
   { "8.1.22", 210812237}: "ddca2f16bfa3d937",
   { "8.1.23", 210812302}: "c0966212aa651e8b",
   { "8.1.26", 210812630}: "a75bd3a39bfcbc77",
   { "8.1.28", 210812860}: "d0795c0dffebea73",
   {"12.0.26", 211202668}: "f012987182d6f16c",
   {"12.0.27", 211202700}: "79b7e56e442e65ed",
   {"12.0.28", 211202876}: "439ba2e3622c344a",
   {"12.0.32", 211203249}: "60e1f010c4e7931e",
   {"12.0.33", 211203304}: "f0613d04b9ba4143",
   {"12.0.34", 211203457}: "843970cb0df053ba",
   {"12.0.36", 211203732}: "a674920042c954d9",
   {"12.0.40", 211204029}: "2c160dbae70b337f",
   {"12.0.44", 211204450}: "7297a39a244189d6",
}

func New_App_Token() (*App_Token, error) {
   app := app_details{"12.0.44", 211204450}
   return app_token_with(app_secrets[app])
}

func app_token_with(app_secret string) (*App_Token, error) {
   key, err := hex.DecodeString(secret_key)
   if err != nil {
      return nil, err
   }
   block, err := aes.NewCipher(key)
   if err != nil {
      return nil, err
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
   var token App_Token
   token.value = base64.StdEncoding.EncodeToString(dst)
   return &token, nil
}

const secret_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"

func pad(b []byte) []byte {
   length := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(length); length >= 1; length-- {
      b = append(b, high)
   }
   return b
}

type App_Token struct {
   value string
}

func location(content_ID string, query url.Values) (string, error) {
   req, err := http.NewRequest("GET", "http://link.theplatform.com", nil)
   if err != nil {
      return "", err
   }
   {
      var b []byte
      b = append(b, "/s/"...)
      b = append(b, cms_account_id...)
      b = append(b, "/media/guid/"...)
      b = strconv.AppendInt(b, aid, 10)
      b = append(b, '/')
      b = append(b, content_ID...)
      req.URL.Path = string(b)
   }
   req.URL.RawQuery = query.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusFound {
      return "", errors.New(res.Status)
   }
   return res.Header.Get("Location"), nil
}

const (
   aid = 2198311517
   cms_account_id = "dJ5BDC"
)

func DASH_CENC(content_ID string) (string, error) {
   query := url.Values{
      "assetTypes": {"DASH_CENC"},
      "formats": {"MPEG-DASH"},
   }
   return location(content_ID, query)
}

func Downloadable(content_ID string) (string, error) {
   query := url.Values{
      "assetTypes": {"Downloadable"},
      "formats": {"MPEG4"},
   }
   return location(content_ID, query)
}

