package hulu

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "bytes"
   "encoding/hex"
   "fmt"
   "os"
   "path"
   "testing"
   "time"
)

func TestLicense(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      data, err := os.ReadFile("authenticate.txt")
      if err != nil {
         t.Fatal(err)
      }
      var auth Authenticate
      err = auth.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      base := path.Base(test.url)
      link, err := auth.DeepLink(&EntityId{base})
      if err != nil {
         t.Fatal(err)
      }
      play, err := auth.Playlist(link)
      if err != nil {
         t.Fatal(err)
      }
      var pssh widevine.PsshData
      pssh.KeyId, err = hex.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      data, err = module.RequestBody()
      if err != nil {
         t.Fatal(err)
      }
      data, err = play.Wrap(data)
      if err != nil {
         t.Fatal(err)
      }
      var body widevine.ResponseBody
      err = body.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      block, err := module.Block(body)
      if err != nil {
         t.Fatal(err)
      }
      containers := body.Container()
      for {
         container, ok := containers()
         if !ok {
            break
         }
         if bytes.Equal(container.Id(), pssh.KeyId) {
            fmt.Printf("%x\n", container.Decrypt(block))
         }
      }
      time.Sleep(time.Second)
   }
}

func TestDetails(t *testing.T) {
   data, err := os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   var auth Authenticate
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      base := path.Base(test.url)
      link, err := auth.DeepLink(&EntityId{base})
      if err != nil {
         t.Fatal(err)
      }
      details, err := auth.Details(link)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", details)
      fmt.Printf("%q\n", text.Name(details))
      time.Sleep(time.Second)
   }
}
