package cbc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "net/url"
   "strconv"
   "strings"
)

func (m metadata) name() (string, error) {
   var b []byte
   switch m.Type {
   case "Movie":
      year, _, found := strings.Cut(m.Date_Created, "-")
      if !found {
         return "", errors.New("invalid dateCreated")
      }
      b = append(b, m.Name...)
      b = append(b, sep_big...)
      b = append(b, year...)
   case "TVEpisode":
      b = append(b, m.Part_Of_Series.Name...)
      b = append(b, sep_big...)
      b = append(b, 'S')
      b = strconv.AppendInt(b, m.Part_Of_Season.Season_Number, 10)
      b = append(b, sep_small)
      b = append(b, 'E')
      b = strconv.AppendInt(b, *m.Episode_Number, 10)
      b = append(b, sep_big...)
      b = append(b, m.Name...)
   default:
      return "", errors.New(m.Type)
   }
   return string(b), nil
}

type lineup_item struct {
   URL string
   Formatted_ID_Media string `json:"formattedIdMedia"`
}

const (
   sep_big = " - "
   sep_small = ' '
)

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

type catalog_gem struct {
   Title string // The Fall
   Content []struct {
      Lineups []struct {
         Season_Number int `json:"seasonNumber"`
         Items []lineup_item
      }
   }
   Structured_Metadata metadata `json:"structuredMetadata"`
   Selected_URL string `json:"selectedUrl"`
}

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
   Type string `json:"@type"`
}

