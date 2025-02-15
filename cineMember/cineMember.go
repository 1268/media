package cineMember

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

func (u *Article) Film() (*Asset, bool) {
   for _, asset := range u.Assets {
      if asset.LinkedType == "film" {
         asset.article = u
         return asset, true
      }
   }
   return nil, false
}

func (e *Entitlement) License(data []byte) (*http.Response, error) {
   return http.Post(
      e.KeyDeliveryUrl, "application/x-protobuf", bytes.NewReader(data),
   )
}

func (e *Entitlement) Mpd() (*http.Response, error) {
   return http.Get(e.Manifest)
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

func (u *User) Unmarshal(data []byte) error {
   return json.Unmarshal(data, u)
}

func (a *Play) Dash() (*Entitlement, bool) {
   for _, title := range a.Data.ArticleAssetPlay.Entitlements {
      if title.Protocol == "dash" {
         return &title, true
      }
   }
   return nil, false
}

func (a *Play) Unmarshal(data []byte) error {
   err := json.Unmarshal(data, a)
   if err != nil {
      return err
   }
   if len(a.Errors) >= 1 {
      return errors.New(a.Errors[0].Message)
   }
   return nil
}

func (a Address) String() string {
   return a[0]
}

func (a *Address) Set(data string) error {
   if !strings.HasPrefix(data, "https://") {
      return errors.New("must start with https://")
   }
   data = strings.TrimPrefix(data, "https://")
   data = strings.TrimPrefix(data, "www.")
   data = strings.TrimPrefix(data, "cinemember.nl")
   data = strings.TrimPrefix(data, "/nl")
   (*a)[0] = strings.TrimPrefix(data, "/")
   return nil
}

func (a Address) Article() (*Article, error) {
   data, err := json.Marshal(map[string]any{
      "query": query_article,
      "variables": map[string]string{
         "articleUrlSlug": a[0]
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
         Article Article
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data.Article, nil
}

func (User) Marshal(email, password string) ([]byte, error) {
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

const query_user = `
mutation UserAuthenticate($email: String, $password: String) {
   UserAuthenticate(email: $email, password: $password) {
      access_token
   }
}
`

// hard geo block
func (Play) Marshal(user1 User, asset *Asset) ([]byte, error) {
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
      "authorization": {"Bearer " + user1.Data.UserAuthenticate.AccessToken},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

const query_asset = `
mutation ArticleAssetPlay($article_id: Int, $asset_id: Int) {
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

type Address [1]string

type Entitlement struct {
   KeyDeliveryUrl string `json:"key_delivery_url"`
   Manifest string
   Protocol string
}

// UserAuthenticate
type User struct {
   Data struct {
      UserAuthenticate struct {
         AccessToken string `json:"access_token"`
      }
   }
}

type Article struct {
   Assets []*Asset
   Id     int
}

type Asset struct {
   Id         int
   LinkedType string `json:"linked_type"`
   article    *Article
}

// ArticleAssetPlay
type Play struct {
   Data struct {
      ArticleAssetPlay struct {
         Entitlements []Entitlement
      }
   }
   Errors []struct {
      Message string
   }
}
