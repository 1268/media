package cineMember

import (
   "bytes"
   "io"
   "net/http"
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

func (u Url) String() string {
   return u.s
}

type Url struct {
   s string
}

func (u *Url) Set(data string) error {
   u.s = strings.TrimPrefix(data, "https://")
   u.s = strings.TrimPrefix(u.s, "www.")
   u.s = strings.TrimPrefix(u.s, "cinemember.nl")
   u.s = strings.TrimPrefix(u.s, "/nl")
   u.s = strings.TrimPrefix(u.s, "/")
   return nil
}
