package user

import "encoding/json"

type user struct {
   Data struct {
      UserAuthenticate struct {
         AccessToken string `json:"access_token"`
      }
   }
}

func marshal(email, password string) ([]byte, error) {
   return nil, nil
}

func (u *user) unmarshal(data []byte) error {
   return json.Unmarshal(data, u)
}
