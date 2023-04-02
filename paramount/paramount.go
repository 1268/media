package paramount

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "net/url"
   "strconv"
   "strings"
)

func (a app_token) item(content_ID string) (*item, error) {
   req := http.Get()
   req.URL.Host = "www.paramountplus.com"
   req.URL.Path = "/apps-api/v2.0/androidphone/video/cid/" + content_ID + ".json"
   req.URL.RawQuery = "at=" + url.QueryEscape(a.at)
   req.URL.Scheme = "https"
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var video struct {
      Item_List []item `json:"itemList"`
   }
   if err := json.NewDecoder(res.Body).Decode(&video); err != nil {
      return nil, err
   }
   if len(video.Item_List) == 0 {
      return nil, errors.New("Item_List length is zero")
   }
   return &video.Item_List[0], nil
}

func (i item) Name() (string, bool) {
   var b strings.Builder
   if i.Media_Type == "Full Episode" {
      b.WriteString("-s")
      b.WriteString(i.Season_Num)
      b.WriteByte('e')
      b.WriteString(i.Episode_Num)
   }
   b.WriteString(mech.Clean(i.Label))
   year, _, ok := strings.Cut(i.Media_Available_Date, "-")
   if !ok {
      return "", false
   }
   b.WriteByte('-')
   b.WriteString(year)
   return b.String(), true
}

type item struct {
   Series_Title string `json:"seriesTitle"`
   Season_Num string `json:"seasonNum"`
   Episode_Num string `json:"episodeNum"`
   Label string
   // 2023-01-15T19:00:00-0800
   Media_Available_Date string `json:"mediaAvailableDate"`
   Media_Type string `json:"mediaType"`
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

func (a app_token) session(content_ID string) (*Session, error) {
   req := http.Get()
   req.URL = &url.URL{
      Scheme: "https",
      Host: "www.paramountplus.com",
      Path: "/apps-api/v3.0/androidphone/irdeto-control/anonymous-session-token.json",
      RawQuery: "at=" + url.QueryEscape(a.at),
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
   sess.URL += content_ID
   return sess, nil
}
