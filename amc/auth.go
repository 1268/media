package amc

import "encoding/json"

type Auth_ID struct {
   Data struct {
      Access_Token string
      Refresh_Token string
   }
}

func (a Auth_ID) Marshal() ([]byte, error) {
   return json.MarshalIndent(a, "", " ")
}

func (a *Auth_ID) Unmarshal(text []byte) error {
   return json.Unmarshal(text, a)
}
