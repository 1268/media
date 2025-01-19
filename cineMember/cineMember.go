package cineMember

import (
   "41.neocities.org/media/cineMember/user"
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strconv"
   "strings"
)

// hard geo block
func (OperationPlay) Marshal(
   auth user.Authenticate, asset *ArticleAsset,
) ([]byte, error) {
   var value struct {
      Query     string `json:"query"`
      Variables struct {
         ArticleId int `json:"article_id"`
         AssetId   int `json:"asset_id"`
      } `json:"variables"`
   }
   value.Query = query_play
   value.Variables.AssetId = asset.Id
   value.Variables.ArticleId = asset.article.Id
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.audienceplayer.com/graphql/2/user",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "authorization": {"Bearer " + auth.Data.UserAuthenticate.AccessToken},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

func (o *OperationPlay) Unmarshal(data []byte) error {
   err := json.Unmarshal(data, o)
   if err != nil {
      return err
   }
   if v := o.Errors; len(v) >= 1 {
      return errors.New(v[0].Message)
   }
   return nil
}

func (o *OperationPlay) Dash() (*Entitlement, bool) {
   for _, title := range o.Data.ArticleAssetPlay.Entitlements {
      if title.Protocol == "dash" {
         return &title, true
      }
   }
   return nil, false
}
const query_play = `
mutation($article_id: Int, $asset_id: Int) {
   ArticleAssetPlay(article_id: $article_id asset_id: $asset_id) {
      entitlements {
         ... on ArticleAssetPlayEntitlement {
            key_delivery_url
            manifest
            protocol
         }
      }
   }
}
`

type Entitlement struct {
   KeyDeliveryUrl string `json:"key_delivery_url"`
   Manifest string
   Protocol string
}

func (e *Entitlement) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      e.KeyDeliveryUrl, "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type OperationPlay struct {
   Data struct {
      ArticleAssetPlay struct {
         Entitlements []Entitlement
      }
   }
   Errors []struct {
      Message string
   }
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

func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   s = strings.TrimPrefix(s, "cinemember.nl")
   s = strings.TrimPrefix(s, "/nl")
   a.Path = strings.TrimPrefix(s, "/")
   return nil
}

type Address struct {
   Path string
}

func (a *Address) String() string {
   return a.Path
}

type ArticleAsset struct {
   Id         int
   LinkedType string `json:"linked_type"`
   article    *OperationArticle
}

func (*OperationArticle) Marshal(web *Address) ([]byte, error) {
   var value struct {
      Query     string `json:"query"`
      Variables struct {
         ArticleUrlSlug string `json:"articleUrlSlug"`
      } `json:"variables"`
   }
   value.Variables.ArticleUrlSlug = web.Path
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

func (o *OperationArticle) Film() (*ArticleAsset, bool) {
   for _, asset := range o.Assets {
      if asset.LinkedType == "film" {
         return asset, true
      }
   }
   return nil, false
}

func (o *OperationArticle) Title() string {
   return o.CanonicalTitle
}

func (o *OperationArticle) Year() int {
   for _, meta := range o.Metas {
      if meta.Key == "year" {
         if v, err := strconv.Atoi(meta.Value); err == nil {
            return v
         }
      }
   }
   return 0
}

func (*OperationArticle) Episode() int {
   return 0
}

func (*OperationArticle) Season() int {
   return 0
}

func (*OperationArticle) Show() string {
   return ""
}

type OperationArticle struct {
   Assets         []*ArticleAsset
   CanonicalTitle string `json:"canonical_title"`
   Id             int
   Metas          []struct {
      Key   string
      Value string
   }
}

func (o *OperationArticle) Unmarshal(data []byte) error {
   var value struct {
      Data struct {
         Article OperationArticle
      }
   }
   err := json.Unmarshal(data, &value)
   if err != nil {
      return err
   }
   *o = value.Data.Article
   for _, asset := range o.Assets {
      asset.article = o
   }
   return nil
}
