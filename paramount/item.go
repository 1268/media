package paramount

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

func (at App_Token) Item(content_ID string) (*Item, error) {
   req, err := http.NewRequest("GET", "https://www.paramountplus.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/apps-api/v2.0/androidphone/video/cid/" + content_ID + ".json"
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
   var video struct {
      Item_List []Item `json:"itemList"`
   }
   if err := json.NewDecoder(res.Body).Decode(&video); err != nil {
      return nil, err
   }
   if len(video.Item_List) == 0 {
      return nil, errors.New("itemList length is zero")
   }
   return &video.Item_List[0], nil
}

type Item struct {
   Series_Title string `json:"seriesTitle"`
   Season_Num string `json:"seasonNum"`
   Episode_Num string `json:"episodeNum"`
   Label string
   Air_Date_ISO string `json:"_airDateISO"`
}

func (i Item) Date() (time.Time, error) {
   return time.Parse("2006-01-02T15:04:05-07:00", i.Air_Date_ISO)
}

// some items stupidly have the show and episode title combined:
// paramountplus.com/shows/video/H87tz3NIw_Ymtcj4zZlWUivmzAPBnMYZ
func (i Item) Series() string {
   before, _, _ := strings.Cut(i.Series_Title, " - ")
   return before
}

func (i Item) Season() (int64, error) {
   return strconv.ParseInt(i.Season_Num, 10, 64)
}

func (i Item) Episode() (int64, error) {
   return strconv.ParseInt(i.Episode_Num, 10, 64)
}

func (i Item) Title() string {
   return i.Label
}
