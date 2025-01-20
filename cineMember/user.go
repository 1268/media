package cineMember

import (
   "bytes"
   "encoding/json"
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

type Address struct {
   s string
}

func (a Address) String() string {
   return a.s
}

func (a *Address) Set(data string) error {
   a.s = strings.TrimPrefix(data, "https://")
   a.s = strings.TrimPrefix(a.s, "www.")
   a.s = strings.TrimPrefix(a.s, "cinemember.nl")
   a.s = strings.TrimPrefix(a.s, "/nl")
   a.s = strings.TrimPrefix(a.s, "/")
   return nil
}
