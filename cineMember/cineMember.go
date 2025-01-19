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
