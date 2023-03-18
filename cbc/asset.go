package cbc

import (
   "encoding/json"
   "strconv"
   "strings"
)

func (a Asset) Name() string {
   var b []byte
   b = append(b, a.Series...)
   if a.Episode >= 1 {
      b = append(b, '-')
      b = append(b, a.Title...)
      b = append(b, '-')
      b = strconv.AppendInt(b, a.Season, 10)
      b = append(b, '-')
      b = strconv.AppendInt(b, a.Episode, 10)
   }
   b = append(b, '-')
   b = append(b, a.Credits.Release_Date...)
   return string(b)
}

type Asset struct {
   Play_Session struct {
      URL string
   } `json:"playSession"`
   Series string
   Title string
   Season int64
   Episode int64
   Credits struct {
      Release_Date string `json:"releaseDate"`
   }
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
