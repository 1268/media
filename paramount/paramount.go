package paramount

import (
   "2a.pages.dev/rosso/http"
   "crypto/aes"
   "encoding/json"
   "strconv"
   "strings"
   "time"
)

func DASH(guid string) string {
   var b strings.Builder
   b.WriteString(media(guid))
   b.WriteByte('?')
   b.WriteString("assetTypes=DASH_CENC")
   b.WriteByte('&')
   b.WriteString("formats=MPEG-DASH")
   return b.String()
}

func HLS_Clear(guid string) string {
   var b strings.Builder
   b.WriteString(media(guid))
   b.WriteByte('?')
   b.WriteString("assetTypes=HLS_CLEAR")
   b.WriteByte('&')
   b.WriteString("formats=MPEG4,M3U")
   return b.String()
}

func Stream_Pack(guid string) string {
   var b strings.Builder
   b.WriteString(media(guid))
   b.WriteByte('?')
   b.WriteString("assetTypes=StreamPack")
   b.WriteByte('&')
   b.WriteString("formats=MPEG4,M3U")
   return b.String()
}

func (p Preview) Name() string {
   var b []byte
   b = append(b, p.Title...)
   if p.Season_Number >= 1 {
      b = append(b, '-')
      b = strconv.AppendInt(b, p.Season_Number, 10)
      b = append(b, '-')
      b = append(b, p.Episode_Number...)
   }
   b = append(b, '-')
   b = append(b, p.Time().Format("2006")...)
   return string(b)
}

func (p Preview) Time() time.Time {
   return time.UnixMilli(p.Pub_Date)
}

type Preview struct {
   Episode_Number string `json:"cbs$EpisodeNumber"`
   GUID string
   Season_Number int64 `json:"cbs$SeasonNumber"`
   Pub_Date int64 `json:"pubDate"`
   Title string
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
