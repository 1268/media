package amc

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/widevine"
   "encoding/base64"
   "os"
   "strings"
   "testing"
)

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

func user_info(name string) ([]string, error) {
   text, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   return strings.Split(string(text), "\n"), nil
}

func Test_Login(t *testing.T) {
   home, err := mech.Home()
   if err != nil {
      t.Fatal(err)
   }
   user, err := user_info(home + "/amc.txt")
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Unauth()
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Login(user[0], user[1]); err != nil {
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
