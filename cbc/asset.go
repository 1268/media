package cbc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "strconv"
)

func New_Asset(id string) (*Asset, error) {
   req := http.Get()
   req.URL.Scheme = "https"
   req.URL.Host = "services.radio-canada.ca"
   req.URL.Path = "/ott/cbc-api/v2/assets/" + id
   client := http.Default_Client
   client.Status = 426
   res, err := client.Do(req)
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

const (
   sep_big = " - "
   sep_small = ' '
)

func (a Asset) Name() string {
   var b []byte
   b = append(b, a.Series...)
   if a.Episode >= 1 {
      b = append(b, sep_big...)
      b = append(b, 'S')
      b = strconv.AppendInt(b, a.Season, 10)
      b = append(b, sep_small)
      b = append(b, 'E')
      b = strconv.AppendInt(b, a.Episode, 10)
      b = append(b, sep_big...)
      b = append(b, a.Title...)
   } else {
      b = append(b, sep_big...)
      b = append(b, a.Credits.Release_Date...)
   }
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
