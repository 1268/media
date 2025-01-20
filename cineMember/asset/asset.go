package asset

import (
   "41.neocities.org/media/cineMember"
   "41.neocities.org/media/cineMember/article"
   "41.neocities.org/media/cineMember/user"
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

// hard geo block
func Marshal(auth user.Authenticate, asset *article.UserAsset) ([]byte, error) {
   data, err := json.Marshal(map[string]any{
      "query": query_play,
      "variables": map[string]int{
         "article_id": asset.GetArticle().Id,
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

func (p *Play) Unmarshal(data []byte) error {
   err := json.Unmarshal(data, p)
   if err != nil {
      return err
   }
   if len(p.Errors) >= 1 {
      return errors.New(p.Errors[0].Message)
   }
   return nil
}

func (p *Play) Dash() (*cineMember.Entitlement, bool) {
   for _, title := range p.Data.ArticleAssetPlay.Entitlements {
      if title.Protocol == "dash" {
         return &title, true
      }
   }
   return nil, false
}
type Play struct {
   Data struct {
      ArticleAssetPlay struct {
         Entitlements []cineMember.Entitlement
      }
   }
   Errors []struct {
      Message string
   }
}
