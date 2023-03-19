package paramount

import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

func session_secret(guid, secret string) (*Session, error) {
   token, err := new_token(secret)
   if err != nil {
      return nil, err
   }
   var buf strings.Builder
   buf.WriteString("https://www.paramountplus.com/apps-api/v3.0/androidphone")
   buf.WriteString("/irdeto-control/anonymous-session-token.json")
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "at=" + url.QueryEscape(token)
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   sess := new(Session)
   if err := json.NewDecoder(res.Body).Decode(sess); err != nil {
      return nil, err
   }
   sess.URL += guid
   return sess, nil
}

func new_token(secret string) (string, error) {
   key, err := hex.DecodeString(secret_key)
   if err != nil {
      return "", err
   }
   block, err := aes.NewCipher(key)
   if err != nil {
      return "", err
   }
   var src []byte
   src = append(src, '|')
   src = append(src, secret...)
   src = pad(src)
   var iv [aes.BlockSize]byte
   cipher.NewCBCEncrypter(block, iv[:]).CryptBlocks(src, src)
   var dst []byte
   dst = append(dst, 0, aes.BlockSize)
   dst = append(dst, iv[:]...)
   dst = append(dst, src...)
   return base64.StdEncoding.EncodeToString(dst), nil
}

func New_Session(guid string) (*Session, error) {
   return session_secret(guid, app_secrets["12.0.28"])
}

type Session struct {
   URL string
   LS_Session string
}

func (s Session) Request_URL() string {
   return s.URL
}

func (s Session) Request_Header() http.Header {
   head := make(http.Header)
   head.Set("Authorization", "Bearer " + s.LS_Session)
   return head
}

func (Session) Request_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (Session) Response_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

// older versions fail
var app_secrets = map[string]string{
   "12.0.28": "439ba2e3622c344a",
   "12.0.27": "79b7e56e442e65ed",
   "12.0.26": "f012987182d6f16c",
   "8.1.28": "d0795c0dffebea73",
   "8.1.26": "a75bd3a39bfcbc77",
   "8.1.23": "c0966212aa651e8b",
   "8.1.22": "ddca2f16bfa3d937",
   "8.1.20": "817774cbafb2b797",
   "8.1.18": "1705732089ff4d60",
   "8.1.16": "add603b54be2fc3c",
   "8.1.14": "",
   "8.1.12": "",
   "8.1.10": "",
   "8.0.54": "",
   "8.0.52": "",
   "8.0.50": "",
   "8.0.48": "",
   "8.0.46": "",
   "8.0.44": "",
   "8.0.42": "",
   "8.0.40": "",
   "8.0.38": "",
   "8.0.36": "",
   "8.0.34": "",
   "8.0.32": "",
   "8.0.30": "",
   "8.0.28": "",
   "8.0.26": "",
   "8.0.24": "",
   "8.0.22": "",
   "8.0.20": "",
   "8.0.16": "",
   "8.0.14": "",
   "8.0.12": "",
   "8.0.10": "",
   "8.0.00": "5d1d865f929d3daa",
   "7.3.58": "4be3d46aecbcd26d",
   "4.8.6": "a958002817953588",
   "4.6.0": "",
   "4.3.5": "",
   "4.3.1": "",
   "4.3.0": "f33bfa06992390a5",
   "4.1.8": "4c5bafd363ca317a",
   "4.1.5": "",
   "4.1.0": "118b561316186716",
   /*
   rg cbs.?appsecret
   */
}
