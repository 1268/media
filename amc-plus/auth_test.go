package amc

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Refresh(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Read_Auth(home + "/amc-plus/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Refresh(); err != nil {
      t.Fatal(err)
   }
   if err := auth.Write_File(home + "/amc-plus/auth.json"); err != nil {
      t.Fatal(err)
   }
}

func Test_Login(t *testing.T) {
   auth, err := Unauth()
   if err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := mech.User(home + "/amc-plus/user.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Login(u["username"], u["password"]); err != nil {
      t.Fatal(err)
   }
   if err := auth.Write_File(home + "/amc-plus/auth.json"); err != nil {
      t.Fatal(err)
   }
}

func Test_Content(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Read_Auth(home + "/amc-plus/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      con, err := auth.Content(test.address)
      if err != nil {
         t.Fatal(err)
      }
      vid, err := con.Video()
      if err != nil {
         t.Fatal(err)
      }
      name, err := mech.Name(vid)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
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
   auth, err := Read_Auth(home + "/amc-plus/auth.json")
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
