package cbc

import (
   "encoding/json"
   "strings"
)

type Asset struct {
   Apple_Content_ID string `json:"appleContentId"`
   Play_Session struct {
      URL string
   } `json:"playSession"`
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
