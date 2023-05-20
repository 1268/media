package cbc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
   "time"
)

func (m Metadata) Season() (int64, error) {
   return m.Part_Of_Season.Season_Number, nil
}

func (m Metadata) Episode() (int64, error) {
   return m.Episode_Number, nil
}

type Metadata struct {
   Part_Of_Series *struct {
      Name string // The Fall
   } `json:"partofSeries"`
   Part_Of_Season struct {
      Season_Number int64 `json:"seasonNumber"`
   } `json:"partofSeason"`
   Episode_Number int64 `json:"episodeNumber"`
   Name string
   Date_Created string `json:"dateCreated"` // 2014-01-01T00:00:00
}

func (m Metadata) Date() (time.Time, error) {
   return time.Parse("2006-01-02T15:04:05", m.Date_Created)
}

func (m Metadata) Series() string {
   if m.Part_Of_Series == nil {
      return ""
   }
   return m.Part_Of_Series.Name
}

func (m Metadata) Title() string {
   return m.Name
}

type Lineup_Item struct {
   URL string
   Formatted_ID_Media string `json:"formattedIdMedia"`
}

func New_Catalog_Gem(link string) (*Catalog_Gem, error) {
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
   gem := new(Catalog_Gem)
   if err := json.NewDecoder(res.Body).Decode(gem); err != nil {
      return nil, err
   }
   return gem, nil
}

func (c Catalog_Gem) Item() *Lineup_Item {
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

type Catalog_Gem struct {
   Content []struct {
      Lineups []struct {
         Items []Lineup_Item
      }
   }
   Selected_URL string `json:"selectedUrl"`
   Structured_Metadata Metadata `json:"structuredMetadata"`
}
