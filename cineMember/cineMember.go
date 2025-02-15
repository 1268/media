package cineMember

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

const query_user = `
mutation($email: String, $password: String) {
   UserAuthenticate(email: $email, password: $password) {
      access_token
   }
}
`

type Authenticate struct {
   Data struct {
      UserAuthenticate struct {
         AccessToken string `json:"access_token"`
      }
   }
}

func (a *Authenticate) Unmarshal(data []byte) error {
   return json.Unmarshal(data, a)
}

func (Authenticate) Marshal(email, password string) ([]byte, error) {
   data, err := json.Marshal(map[string]any{
      "query": query_user,
      "variables": map[string]string{
         "email": email,
         "password": password,
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
func (a *AssetPlay) Dash() (*Entitlement, bool) {
   for _, title := range a.Data.ArticleAssetPlay.Entitlements {
      if title.Protocol == "dash" {
         return &title, true
      }
   }
   return nil, false
}

const query_asset = `
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

type AssetPlay struct {
   Data struct {
      ArticleAssetPlay struct {
         Entitlements []Entitlement
      }
   }
   Errors []struct {
      Message string
   }
}

func (a *AssetPlay) Unmarshal(data []byte) error {
   err := json.Unmarshal(data, a)
   if err != nil {
      return err
   }
   if len(a.Errors) >= 1 {
      return errors.New(a.Errors[0].Message)
   }
   return nil
}

// hard geo block
func (AssetPlay) Marshal(user Authenticate, asset *UserAsset) ([]byte, error) {
   data, err := json.Marshal(map[string]any{
      "query": query_asset,
      "variables": map[string]int{
         "article_id": asset.article.Id,
         "asset_id": asset.Id,
      },
   })
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
      "authorization": {"Bearer " + user.Data.UserAuthenticate.AccessToken},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}
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
