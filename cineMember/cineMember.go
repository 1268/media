package cineMember

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

func (e *Entitlement) License(data []byte) (*http.Response, error) {
   return http.Post(
      e.KeyDeliveryUrl, "application/x-protobuf", bytes.NewReader(data),
   )
}

func (e *Entitlement) Mpd() (*http.Response, error) {
   return http.Get(e.Manifest)
}

type Entitlement struct {
   KeyDeliveryUrl string `json:"key_delivery_url"`
   Manifest string
   Protocol string
}

func (a *Address) Set(data string) error {
   if !strings.HasPrefix(data, "https://") {
      return errors.New("must start with https://")
   }
   a.s = strings.TrimPrefix(data, "https://")
   a.s = strings.TrimPrefix(a.s, "www.")
   a.s = strings.TrimPrefix(a.s, "cinemember.nl")
   a.s = strings.TrimPrefix(a.s, "/nl")
   a.s = strings.TrimPrefix(a.s, "/")
   return nil
}

type Address struct {
   s string
}

func (a Address) String() string {
   return a.s
}

// NO ANONYMOUS QUERY
const query_article = `
query Article($articleUrlSlug: String) {
   Article(full_url_slug: $articleUrlSlug) {
      ... on Article {
         assets {
            ... on Asset {
               id
               linked_type
            }
         }
         id
      }
   }
}
`

type UserAsset struct {
   Id         int
   LinkedType string `json:"linked_type"`
   article    *UserArticle
}

func (a Address) Article() (*UserArticle, error) {
   data, err := json.Marshal(map[string]any{
      "query": query_article,
      "variables": map[string]string{
         "articleUrlSlug": a.String(),
      },
   })
   if err != nil {
      return nil, err
   }
   resp, err := http.Post(
      "https://api.audienceplayer.com/graphql/2/user",
      "application/json", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var value struct {
      Data struct {
         Article UserArticle
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data.Article, nil
}

type UserArticle struct {
   Assets []*UserAsset
   Id     int
}

func (u *UserArticle) Film() (*UserAsset, bool) {
   for _, asset := range u.Assets {
      if asset.LinkedType == "film" {
         asset.article = u
         return asset, true
      }
   }
   return nil, false
}
