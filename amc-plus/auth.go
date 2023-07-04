package amc

import "encoding/json"

type Auth struct {
   Data struct {
      Access_Token string
      Refresh_Token string
   }
}

func (a Auth) Marshal() ([]byte, error) {
   return json.MarshalIndent(a, "", " ")
}

func (a *Auth) Unmarshal(b []byte) error {
   return json.Unmarshal(b, a)
}
