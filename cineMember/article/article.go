package article

import (
   "41.neocities.org/media/cineMember"
   "bytes"
   "encoding/json"
   "io"
   "net/http"
)

func Marshal(url cineMember.Url) ([]byte, error) {
   data, err := json.Marshal(map[string]any{
      "query": query,
      "variables": map[string]string{
         "articleUrlSlug": url.String(),
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
   return io.ReadAll(resp.Body)
}

// NO ANONYMOUS QUERY
const query = `
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

type Article struct {
   Assets []*UserAsset
   Id     int
}

type UserAsset struct {
   Id         int
   LinkedType string `json:"linked_type"`
   article    *Article
}

func (u *UserAsset) GetArticle() *Article {
   return u.article
}

func (a *Article) Unmarshal(data []byte) error {
   var value struct {
      Data struct {
         Article Article
      }
   }
   err := json.Unmarshal(data, &value)
   if err != nil {
      return err
   }
   *a = value.Data.Article
   for _, asset := range a.Assets {
      asset.article = a
   }
   return nil
}

func (a *Article) Film() (*UserAsset, bool) {
   for _, asset := range a.Assets {
      if asset.LinkedType == "film" {
         return asset, true
      }
   }
   return nil, false
}
