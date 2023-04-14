package paramount

import (
   "strconv"
   "strings"
)

func media(content_ID string) string {
   var b []byte
   b = append(b, "http://link.theplatform.com/s/"...)
   b = append(b, cms_account_id...)
   b = append(b, "/media/guid/"...)
   b = strconv.AppendInt(b, aid, 10)
   b = append(b, '/')
   b = append(b, content_ID...)
   return string(b)
}

const (
   aid = 2198311517
   cms_account_id = "dJ5BDC"
)

func DASH_CENC(content_ID string) string {
   var b strings.Builder
   b.WriteString(media(content_ID))
   b.WriteByte('?')
   b.WriteString("assetTypes=DASH_CENC")
   b.WriteByte('&')
   b.WriteString("formats=MPEG-DASH")
   return b.String()
}

func Downloadable(content_ID string) string {
   var b strings.Builder
   b.WriteString(media(content_ID))
   b.WriteByte('?')
   b.WriteString("assetTypes=Downloadable")
   b.WriteByte('&')
   b.WriteString("formats=MPEG4")
   return b.String()
}
