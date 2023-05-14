package cbc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "fmt"
   "net/url"
   "strings"
)

func (m metadata) name() (string, error) {
   var s []string
   if m.Part_Of_Series != nil {
      s = append(s, m.Part_Of_Series.Name)
   }
   if m.Part_Of_Season != nil {
      s = append(s, fmt.Sprint("S", m.Part_Of_Season.Season_Number))
   }
   if m.Episode_Number != nil {
      s = append(s, fmt.Sprint("E", *m.Episode_Number))
      s = append(s, m.Name)
   } else {
      s = append(s, m.Name)
      year, _, found := strings.Cut(m.Date_Created, "-")
      if !found {
         return "", fmt.Errorf("invalid dateCreated")
      }
      s = append(s, year)
   }
   return strings.Join(s, sep_big), nil
}

const sep_big = " - "

type metadata struct {
   Part_Of_Series *struct {
      Name string
   } `json:"partofSeries"`
   Part_Of_Season *struct {
      Season_Number int64 `json:"seasonNumber"`
   } `json:"partofSeason"`
   Episode_Number *int64 `json:"episodeNumber"`
   Name string
   Date_Created string `json:"dateCreated"` // 2015-01-01T00:00:00
}

type catalog_gem struct {
   Selected_URL string `json:"selectedUrl"`
   Content []struct {
      Lineups []struct {
         Items []lineup_item
      }
   }
   Structured_Metadata metadata `json:"structuredMetadata"`
}

type lineup_item struct {
   URL string
   Formatted_ID_Media string `json:"formattedIdMedia"`
}

func new_catalog_gem(link string) (*catalog_gem, error) {
   // you can also use `phone_android`, but it returns combined number and name:
   // 3. Beauty Hath Strange Power
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "services.radio-canada.ca",
      Path: "/ott/catalog/v2/gem/show/" + link,
      RawQuery: "device=web",
   })
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

func (c catalog_gem) item() *lineup_item {
   for _, content := range c.Content {
      for _, lineup := range content.Lineups {
         for _, item := range lineup.Items {
            if item.URL == c.Selected_URL {
               return &item
            }
         }
      }
   }
   return nil
}

