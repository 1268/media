package paramount

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/rosso/http"
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "net/url"
   "strconv"
   "strings"
   "time"
)

type Preview struct {
   Series_Title string `json:"cbs$SeriesTitle"`
   Season_Number int64 `json:"cbs$SeasonNumber"`
   Episode_Number string `json:"cbs$EpisodeNumber"`
   Title string
   Pub_Date int64 `json:"pubDate"`
}

func session_secret(guid, secret string) (*Session, error) {
   token, err := new_token(secret)
   if err != nil {
      return nil, err
   }
   req := http.Get()
   req.URL = &url.URL{
      Scheme: "https",
      Host: "www.paramountplus.com",
      Path: "/apps-api/v3.0/androidphone/irdeto-control/anonymous-session-token.json",
      RawQuery: "at=" + url.QueryEscape(token),
   }
   res, err := http.Default_Client.Do(req)
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

func New_Session(guid string) (*Session, error) {
   return session_secret(guid, app_secrets["12.0.40"])
}

var app_secrets = map[string]string{
   "12.0.40": "2c160dbae70b337f",
   "12.0.36": "a674920042c954d9",
   "12.0.34": "843970cb0df053ba",
   "12.0.33": "f0613d04b9ba4143",
   "12.0.32": "60e1f010c4e7931e",
   "12.0.28": "439ba2e3622c344a",
   "12.0.27": "79b7e56e442e65ed",
   "12.0.26": "f012987182d6f16c",
   "08.1.28": "d0795c0dffebea73",
   "08.1.26": "a75bd3a39bfcbc77",
   "08.1.23": "c0966212aa651e8b",
   "08.1.22": "ddca2f16bfa3d937",
   "08.1.20": "817774cbafb2b797",
   "08.1.18": "1705732089ff4d60",
   "08.1.16": "add603b54be2fc3c",
   "08.1.14": "acacc94f4214ddc1",
   "08.1.12": "3395c01da67a358b",
   "08.1.10": "8150802ffdeffaf0",
   "08.0.54": "6c70b33080758409",
   "08.0.52": "5fcf8c6937dba442",
   "08.0.50": "2e6cd61ba21dd0c4",
   "08.0.48": "00a7ea18c54d674c",
   "08.0.46": "88065c1d30bc15ed",
   "08.0.44": "9c5b3eda87e38402",
   "08.0.42": "c824c27d68eb9fc3",
   "08.0.40": "d08c12908070b2ac",
   "08.0.38": "423187842fdd7eac",
   "08.0.36": "6dfcc58b09fca975",
   "08.0.34": "0f84a8e9f62594ad",
   "08.0.32": "262d30953b16032b",
   "08.0.30": "90946a66385ceeb5",
   "08.0.28": "1fc4f2e07173b30c",
   "08.0.26": "860c7062bb69759d",
   "08.0.24": "2b7feb264967d94f",
   "08.0.22": "36a841291cfecc4e",
   "08.0.20": "003ff1f049feb54a",
   "08.0.16": "79e71194ad8b97d4",
   "08.0.14": "f3577b860abfa76d",
   "08.0.12": "20021bb2eda91db4",
   "08.0.10": "685c401ff9a4a2d9",
   "08.0.00": "5d1d865f929d3daa",
   "07.3.58": "4be3d46aecbcd26d",
   "04.8.06": "a958002817953588",
}

func (p Preview) Name() string {
   var b []byte
   b = append(b, mech.Clean(p.Title)...)
   if p.Season_Number >= 1 {
      b = append(b, "-s"...)
      b = strconv.AppendInt(b, p.Season_Number, 10)
      b = append(b, "e"...)
      b = append(b, p.Episode_Number...)
   }
   b = append(b, '-')
   b = append(b, p.Time().Format("2006")...)
   return string(b)
}

func (p Preview) Time() time.Time {
   return time.UnixMilli(p.Pub_Date)
}

const secret_key = "302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"

func pad(b []byte) []byte {
   length := aes.BlockSize - len(b) % aes.BlockSize
   for high := byte(length); length >= 1; length-- {
      b = append(b, high)
   }
   return b
}

const (
   aid = 2198311517
   sid = "dJ5BDC"
)

func New_Preview(guid string) (*Preview, error) {
   req, err := http.Get_URL(media(guid))
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "format=preview"
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   prev := new(Preview)
   if err := json.NewDecoder(res.Body).Decode(prev); err != nil {
      return nil, err
   }
   return prev, nil
}

func media(guid string) string {
   var b []byte
   b = append(b, "http://link.theplatform.com/s/"...)
   b = append(b, sid...)
   b = append(b, "/media/guid/"...)
   b = strconv.AppendInt(b, aid, 10)
   b = append(b, '/')
   b = append(b, guid...)
   return string(b)
}

func DASH_CENC(guid string) string {
   var b strings.Builder
   b.WriteString(media(guid))
   b.WriteByte('?')
   b.WriteString("assetTypes=DASH_CENC")
   b.WriteByte('&')
   b.WriteString("formats=MPEG-DASH")
   return b.String()
}

func Downloadable(guid string) string {
   var b strings.Builder
   b.WriteString(media(guid))
   b.WriteByte('?')
   b.WriteString("assetTypes=Downloadable")
   b.WriteByte('&')
   b.WriteString("formats=MPEG4")
   return b.String()
}
