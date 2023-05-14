package cbc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
)

type catalog_gem struct {
   Content []struct {
      Lineups []struct {
         Items []struct {
            ID_Media int `json:"idMedia"`
         }
      }
   }
   Structured_Metadata struct {
      Date_Created string `json:"dateCreated"`
      Episode_Number *int `json:"episodeNumber"`
      Name string
      Part_Of_Season *struct {
         Season_Number int `json:"seasonNumber"`
      } `json:"partofSeason"`
   } `json:"structuredMetadata"`
   Title string
}

func new_catalog_gem(link string) (*catalog_gem, error) {
   req := http.Get()
   req.URL.Scheme = "https"
   req.URL.Host = "services.radio-canada.ca"
   req.URL.Path = "/ott/catalog/v2/gem/show/" + link
   // you can also use `phone_android`, but it returns combined number and name:
   // 3. Beauty Hath Strange Power
   req.URL.RawQuery = "device=web"
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   gem := new(catalog_gem)
   if err := json.NewDecoder(res.Body).Decode(gem); err != nil {
      return nil, err
   }
   return gem, nil
}
