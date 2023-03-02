package cbc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "strings"
   "time"
)

const forwarded_for = "99.224.0.0"

var Client = http.Default_Client

// gem.cbc.ca/media/downton-abbey/s01e05
func Get_ID(input string) string {
   _, after, found := strings.Cut(input, "/media/")
   if found {
      return after
   }
   return input
}

type Asset struct {
   Air_Date int64 `json:"airDate"`
   Apple_Content_ID string `json:"appleContentId"`
   Duration int64
   Play_Session struct {
      URL string
   } `json:"playSession"`
   Series string
   Title string
}

func New_Asset(id string) (*Asset, error) {
   var buf strings.Builder
   buf.WriteString("https://services.radio-canada.ca/ott/cbc-api/v2/assets/")
   buf.WriteString(id)
   res, err := Client.Get(buf.String())
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

type Media struct {
   Message *string
   URL *string
}

func (p Profile) Media(a *Asset) (*Media, error) {
   req, err := http.NewRequest("GET", a.Play_Session.URL, nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "X-Claims-Token": {p.Claims_Token},
      "X-Forwarded-For": {forwarded_for},
   }
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   med := new(Media)
   if err := json.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   if med.Message != nil {
      return nil, errors.New(*med.Message)
   }
   return med, nil
}
