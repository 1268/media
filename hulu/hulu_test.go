package hulu

import (
   "41.neocities.org/widevine"
   "bytes"
   "encoding/hex"
   "fmt"
   "os"
   "os/exec"
   "path"
   "strings"
   "testing"
   "time"
)

func TestAuthenticate(t *testing.T) {
   data, err := exec.Command("password", "hulu.com").Output()
   if err != nil {
      t.Fatal(err)
   }
   username, password, _ := strings.Cut(string(data), ":")
   data, err = new(Authenticate).Marshal(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.txt", data, os.ModePerm)
}

var tests = []struct {
   content string
   key_id  string
   url     string
}{
   {
      content: "episode",
      key_id:  "21b82dc2ebb24d5aa9f8631f04726650",
      url:     "hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d",
   },
   {
      content: "film",
      url:     "hulu.com/watch/f70dfd4d-dbfb-46b8-abb3-136c841bba11",
   },
}

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
      key_id, err := hex.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var pssh widevine.PsshData
      pssh.KeyIds = [][]byte{key_id}
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
         if bytes.Equal(container.Id(), key_id) {
            fmt.Printf("%x\n", container.Key(block))
         }
      }
      time.Sleep(time.Second)
   }
}
