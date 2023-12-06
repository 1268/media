package amc

import (
   "154.pages.dev/stream"
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Content(t *testing.T) {
   var auth Auth_ID
   {
      s, err := os.UserHomeDir()
      if err != nil {
         t.Fatal(err)
      }
      b, err := os.ReadFile(s + "/amc/auth.json")
      if err != nil {
         t.Fatal(err)
      }
      auth.Unmarshal(b)
   }
   for _, test := range tests {
      con, err := auth.Content(Path{test.path})
      if err != nil {
         t.Fatal(err)
      }
      vid, err := con.Video()
      if err != nil {
         t.Fatal(err)
      }
      name, err := stream.Format_Film(vid)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}

func Test_Refresh(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var auth Auth_ID
   {
      b, err := os.ReadFile(home + "/amc/auth.json")
      if err != nil {
         t.Fatal(err)
      }
      auth.Unmarshal(b)
   }
   if err := auth.Refresh(); err != nil {
      t.Fatal(err)
   }
   {
      b, err := auth.Marshal()
      if err != nil {
         t.Fatal(err)
      }
      os.WriteFile(home + "/amc/auth.json", b, 0666)
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
   u, err := user(home + "/amc/user.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Login(u["username"], u["password"]); err != nil {
      t.Fatal(err)
   }
   text, err := auth.Marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile(home + "/amc/auth.json", text, 0666)
}

func Test_Key(t *testing.T) {
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
   test := tests[0]
   pssh, err := base64.StdEncoding.DecodeString(test.pssh)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, nil, pssh)
   if err != nil {
      t.Fatal(err)
   }
   var auth Auth_ID
   {
      b, err := os.ReadFile(home + "/amc/auth.json")
      if err != nil {
         t.Fatal(err)
      }
      auth.Unmarshal(b)
   }
   play, err := auth.Playback(Path{test.path})
   if err != nil {
      t.Fatal(err)
   }
   key, err := mod.Key(play)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
