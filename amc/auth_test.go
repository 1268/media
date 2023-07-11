package amc

import (
   "154.pages.dev/media"
   "154.pages.dev/widevine"
   "encoding/base64"
   "os"
   "testing"
)

func Test_Login(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Unauth()
   if err != nil {
      t.Fatal(err)
   }
   u, err := media.User(home + "/amc-plus/user.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Login(u["username"], u["password"]); err != nil {
      t.Fatal(err)
   }
   {
      b, err := auth.Marshal()
      if err != nil {
         t.Fatal(err)
      }
      os.WriteFile(home + "/amc-plus/auth.json", b, 0666)
   }
}

func Test_Post(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/widevine/client_id.bin")
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
   var auth Auth_ID
   {
      b, err := os.ReadFile(home + "/amc-plus/auth.json")
      if err != nil {
         t.Fatal(err)
      }
      auth.Unmarshal(b)
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

