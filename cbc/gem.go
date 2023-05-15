package cbc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
   "time"
)

type metadata struct {
   Part_Of_Series *struct {
      Name string // The Fall
   } `json:"partofSeries"`
   Part_Of_Season *struct {
      Season_Number int64 `json:"seasonNumber"`
   } `json:"partofSeason"`
   Episode_Number *int64 `json:"episodeNumber"`
   Name string
   Date_Created string `json:"dateCreated"` // 2014-01-01T00:00:00
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

func (m metadata) Date() (time.Time, error) {
   return time.Parse("2006-01-02T15:04:05", m.Date_Created)
}

func (m metadata) Episode() int64 {
   if m.Episode_Number == nil {
      return 0
   }
   return *m.Episode_Number
}

type catalog_gem struct {
   Content []struct {
      Lineups []struct {
         Items []lineup_item
      }
   }
   Selected_URL string `json:"selectedUrl"`
   Structured_Metadata metadata `json:"structuredMetadata"`
}

func (m metadata) Season() int64 {
   if m.Part_Of_Season == nil {
      return 0
   }
   return m.Part_Of_Season.Season_Number
}

func (m metadata) Series() string {
   if m.Part_Of_Series == nil {
      return ""
   }
   return m.Part_Of_Series.Name
}

func (m metadata) Title() string {
   return m.Name
}
