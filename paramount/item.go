package paramount

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "net/url"
   "strings"
)

func (at app_token) item(content_ID string) (*item, error) {
   req := http.Get()
   req.URL.Host = "www.paramountplus.com"
   req.URL.Path = "/apps-api/v2.0/androidphone/video/cid/" + content_ID + ".json"
   req.URL.RawQuery = "at=" + url.QueryEscape(at.at)
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
