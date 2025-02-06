package cineMember

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

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
