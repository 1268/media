package cineMember

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
)

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

type UserArticle struct {
   Assets []*UserAsset
   Id     int
}

func (u *UserArticle) Film() (*UserAsset, bool) {
   for _, asset := range u.Assets {
      if asset.LinkedType == "film" {
         return asset, true
      }
   }
   return nil, false
}

func (u *UserArticle) Unmarshal(data []byte) error {
   var value struct {
      Data struct {
         Article UserArticle
      }
   }
   err := json.Unmarshal(data, &value)
   if err != nil {
      return err
   }
   *u = value.Data.Article
   for _, asset := range u.Assets {
      asset.article = u
   }
   return nil
}

func (UserArticle) Marshal(web Address) ([]byte, error) {
   data, err := json.Marshal(map[string]any{
      "query": query_article,
      "variables": map[string]string{
         "articleUrlSlug": web.String(),
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

type UserAsset struct {
   Id         int
   LinkedType string `json:"linked_type"`
   article    *UserArticle
}
