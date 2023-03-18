package cbc

import (
   "encoding/json"
   "strings"
   "time"
)

type Asset struct {
   Title string
   Air_Date int64 `json:"airDate"`
   Apple_Content_ID string `json:"appleContentId"`
   Duration int64
   Play_Session struct {
      URL string
   } `json:"playSession"`
   Series string
}

func New_Asset(id string) (*Asset, error) {
   var b strings.Builder
   b.WriteString("https://services.radio-canada.ca/ott/cbc-api/v2/assets/")
   b.WriteString(id)
   res, err := Client.Get(b.String())
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   a := new(Asset)
   if err := json.NewDecoder(res.Body).Decode(a); err != nil {
      return nil, err
   }
   return a, nil
}

func (a Asset) Get_Duration() time.Duration {
   return time.Duration(a.Duration) * time.Second
}

func (a Asset) Get_Time() time.Time {
   return time.UnixMilli(a.Air_Date)
}

func (a Asset) String() string {
   var b strings.Builder
   b.WriteString("ID: ")
   b.WriteString(a.Apple_Content_ID)
   b.WriteString("\nSeries: ")
   b.WriteString(a.Series)
   b.WriteString("\nTitle: ")
   b.WriteString(a.Title)
   b.WriteString("\nDate: ")
   b.WriteString(a.Get_Time().String())
   b.WriteString("\nDuration: ")
   b.WriteString(a.Get_Duration().String())
   return b.String()
}
