package cineMember

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

func (o *OperationUser) Unmarshal(data []byte) error {
   return json.Unmarshal(data, o)
}

type OperationUser struct {
   Data struct {
      UserAuthenticate struct {
         AccessToken string `json:"access_token"`
      }
   }
}

func (OperationUser) Marshal(email, password string) ([]byte, error) {
   var value struct {
      Query     string `json:"query"`
      Variables struct {
         Email    string `json:"email"`
         Password string `json:"password"`
      } `json:"variables"`
   }
   value.Query = query_user
   value.Variables.Email = email
   value.Variables.Password = password
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

