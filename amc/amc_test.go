package amc

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/widevine"
   "encoding/base64"
   "encoding/json"
   "os"
   "testing"
)

func Test_Login(t *testing.T) {
   home, err := mech.Home()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user(home + "/amc.json")
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Unauth()
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Login(u["username"], u["password"]); err != nil {
      t.Fatal(err)
   }
   if err := auth.Write_File(home + "/amc.json"); err != nil {
      t.Fatal(err)
   }
}

func Test_Refresh(t *testing.T) {
   home, err := mech.Home()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Read_Auth(home + "/amc.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Refresh(); err != nil {
      t.Fatal(err)
   }
   if err := auth.Write_File(home + "/amc.json"); err != nil {
      t.Fatal(err)
   }
}
func user(name string) (map[string]string, error) {
   b, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   var m map[string]string
   if err := json.Unmarshal(b, &m); err != nil {
      return nil, err
   }
   return m, nil
}

func Test_Post(t *testing.T) {
   home, err := mech.Home()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   test := tests[episode]
   pssh, err := base64.StdEncoding.DecodeString(test.pssh)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, pssh)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Read_Auth(home + "/amc.json")
   if err != nil {
      t.Fatal(err)
   }
   play, err := auth.Playback(test.address)
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Post(play)
   if err != nil {
      t.Fatal(err)
   }
   if keys.Content().String() != test.key {
      t.Fatal(keys)
   }
}

