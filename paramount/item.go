package paramount

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "net/url"
   "strings"
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

type Item struct {
   Episode_Num string `json:"episodeNum"`
   Label string
   // 2023-01-15T19:00:00-0800
   Media_Available_Date string `json:"mediaAvailableDate"`
   Season_Num string `json:"seasonNum"`
   Series_Title string `json:"seriesTitle"`
   Media_Type string `json:"mediaType"`
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
