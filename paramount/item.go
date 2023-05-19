package paramount

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "net/url"
   "strconv"
   "time"
)

func (at App_Token) Item(content_ID string) (*Item, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "www.paramountplus.com",
      Path: "/apps-api/v2.0/androidphone/video/cid/" + content_ID + ".json",
      RawQuery: "at=" + url.QueryEscape(at.value),
   })
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var video struct {
      Item_List []Item `json:"itemList"`
   }
   if err := json.NewDecoder(res.Body).Decode(&video); err != nil {
      return nil, err
   }
   if len(video.Item_List) == 0 {
      return nil, errors.New("Item_List length is zero")
   }
   return &video.Item_List[0], nil
}

func (i Item) Series() string {
   return i.Series_Title
}

func (i Item) Title() string {
   return i.Label
}

func (i Item) Date() (time.Time, error) {
   return time.Parse(time.RFC3339, i.Media_Available_Date)
}

type Item struct {
   Series_Title string `json:"seriesTitle"`
   Season_Num string `json:"seasonNum"`
   Episode_Num string `json:"episodeNum"`
   Label string
   // 2023-01-15T19:00:00-0800
   Media_Available_Date string `json:"mediaAvailableDate"`
}

func (i Item) Season() (int64, error) {
   return strconv.ParseInt(i.Season_Num, 10, 64)
}

func (i Item) Episode() (int64, error) {
   return strconv.ParseInt(i.Episode_Num, 10, 64)
}
