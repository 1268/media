package roku

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   for _, test := range tests {
      var home HomeScreen
      err := home.New(test.id)
      if err != nil {
         t.Fatal(err)
      }
      name := text.Name(&Namer{home})
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
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
      var pssh widevine.Pssh
      pssh.ContentId, err = base64.StdEncoding.DecodeString(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      pssh.KeyId, err = base64.StdEncoding.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Module
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      var auth AccountAuth
      data, err := auth.Marshal(nil)
      if err != nil {
         t.Fatal(err)
      }
      err = auth.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      play, err := auth.Playback(test.id)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(play, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}

var tests = map[string]struct {
   content_id string
   id         string
   key_id     string
   url        string
}{
   "episode": {
      content_id: "Kg==",
      id:         "105c41ea75775968b670fbb26978ed76",
      key_id:     "vfpNbNs5cC5baB+QYX+afg==",
      url:        "therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76",
   },
   "movie": {
      content_id: "Kg==",
      id:         "597a64a4a25c5bf6af4a8c7053049a6f",
      key_id:     "KDOa149zRSDaJObgVz05Lg==",
      url:        "therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f",
   },
}
