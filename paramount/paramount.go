package paramount

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/rosso/http"
   "errors"
   "strconv"
   "strings"
)

const (
   aid = 2198311517
   sid = "dJ5BDC"
)

func media(content_ID string) string {
   var b []byte
   b = append(b, "http://link.theplatform.com/s/"...)
   b = append(b, sid...)
   b = append(b, "/media/guid/"...)
   b = strconv.AppendInt(b, aid, 10)
   b = append(b, '/')
   b = append(b, content_ID...)
   return string(b)
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

func (i Item) String() string {
   var b strings.Builder
   if i.Media_Type == "Full Episode" {
      b.WriteString("series title: ")
      b.WriteString(i.Series_Title)
      b.WriteString("\nseason num: ")
      b.WriteString(i.Season_Num)
      b.WriteString("\nepisode num: ")
      b.WriteString(i.Episode_Num)
      b.WriteByte('\n')
   }
   b.WriteString("label: ")
   b.WriteString(i.Label)
   b.WriteString("\nmedia available date: ")
   b.WriteString(i.Media_Available_Date)
   return b.String()
}

type Item struct {
   Episode_Num string `json:"episodeNum"`
   Label string
   // 2023-01-15T19:00:00-0800
   Media_Available_Date string `json:"mediaAvailableDate"`
   Media_Type string `json:"mediaType"`
   Season_Num string `json:"seasonNum"`
   Series_Title string `json:"seriesTitle"`
}

const (
   sep_big = " - "
   sep_small = ' '
)

func (i Item) Name() (string, error) {
   var b strings.Builder
   if i.Media_Type == "Full Episode" {
      b.WriteString(i.Series_Title)
      b.WriteString(sep_big)
      b.WriteByte('S')
      b.WriteString(i.Season_Num)
      b.WriteByte(sep_small)
      b.WriteByte('E')
      b.WriteString(i.Episode_Num)
      b.WriteString(sep_big)
   }
   b.WriteString(mech.Clean(i.Label))
   if i.Media_Type == "Movie" {
      year, _, found := strings.Cut(i.Media_Available_Date, "-")
      if !found {
         return "", errors.New("year not found")
      }
      b.WriteString(sep_big)
      b.WriteString(year)
   }
   return b.String(), nil
}

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
