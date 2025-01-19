package article

import (
   "41.neocities.org/media/cineMember"
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strconv"
)

type Asset struct {
   Id         int
   LinkedType string `json:"linked_type"`
   article    *Article
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
         canonical_title
         id
         metas(output: html) {
            ... on ArticleMeta {
               key
               value
            }
         }
      }
   }
}
`

func Marshal(url cineMember.Url) ([]byte, error) {
   var value struct {
      Query     string `json:"query"`
      Variables struct {
         ArticleUrlSlug string `json:"articleUrlSlug"`
      } `json:"variables"`
   }
   value.Variables.ArticleUrlSlug = url.String()
   value.Query = query_article
   data, err := json.Marshal(value)
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
   for _, as := range a.Assets {
      as.article = a
   }
   return nil
}

func (a *Article) Title() string {
   return a.CanonicalTitle
}

func (a *Article) Year() int {
   for _, meta := range a.Metas {
      if meta.Key == "year" {
         if v, err := strconv.Atoi(meta.Value); err == nil {
            return v
         }
      }
   }
   return 0
}

func (*Article) Episode() int {
   return 0
}

func (*Article) Season() int {
   return 0
}

func (*Article) Show() string {
   return ""
}

type Article struct {
   Assets         []*Asset
   CanonicalTitle string `json:"canonical_title"`
   Id             int
   Metas          []struct {
      Key   string
      Value string
   }
}

func (a *Article) Film() (*Asset, bool) {
   for _, as := range a.Assets {
      if as.LinkedType == "film" {
         return as, true
      }
   }
   return nil, false
}
