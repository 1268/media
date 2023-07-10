package amc

import (
   "encoding/json"
   "os"
   "path/filepath"
)

func Set_Env() error {
   if os.Getenv("AMC_PLUS") != "" {
      return nil
   }
   v, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   v = filepath.Join(v, "amc-plus", "auth.json")
   return os.Setenv("AMC_PLUS", v)
}

type Auth_ID struct {
   Data struct {
      Access_Token string
      Refresh_Token string
   }
}

func (a Auth_ID) Marshal() ([]byte, error) {
   return json.MarshalIndent(a, "", " ")
}

func (a *Auth_ID) Unmarshal(b []byte) error {
   return json.Unmarshal(b, a)
}
