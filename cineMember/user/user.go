package user

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
)

const query_user = `
mutation($email: String, $password: String) {
   UserAuthenticate(email: $email, password: $password) {
      access_token
   }
}
`

func marshal(email, password string) ([]byte, error) {
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
