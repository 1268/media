package itv

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "path"
   "testing"
   "time"
)

var tests = []struct{
   content_id string
   key_id string
   legacy_id LegacyId
   url string
}{
   {
      content_id: "MS05MzU5LTAwMDEtMDA0XzM0",
      key_id: "CqlTfeulS5mlu9fu81QtfQ==",
      legacy_id: LegacyId{"1", "9359", "0001"},
      url: "itv.com/watch/quantum-of-solace/1a9359",
   },
   {
      content_id: "MTAtMzkxNS0wMDAyLTAwMV8zNA==",
      key_id: "zCXIAYrkT9+eG6gbjNG1Qw==",
      legacy_id: LegacyId{"10", "3915", "0002"},
      url: "itv.com/watch/community/10a3915/10a3915a0002",
   },
   {
      content_id: "MTAtMzkxOC0wMDAxLTAwMV8zNA==",
      key_id: "znjzKgOaRBqJMBDGiUDN8g==",
      legacy_id: LegacyId{"10", "3918", "0001"},
      url: "itv.com/watch/joan/10a3918/10a3918a0001",
   },
}

func TestPlaylist(t *testing.T) {
   for _, test := range tests {
      discovery, err := test.legacy_id.Discovery()
      if err != nil {
         t.Fatal(err)
      }
      play, err := discovery.Playlist()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(play.Resolution720())
      time.Sleep(time.Second)
   }
}

func TestLegacyId(t *testing.T) {
   for _, test := range tests {
      var id LegacyId
      err := id.Set(path.Base(test.url))
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(id)
   }
}

func TestDiscovery(t *testing.T) {
   for _, test := range tests {
      discovery, err := test.legacy_id.Discovery()
      if err != nil {
         t.Fatal(err)
      }
      name := text.Name(Namer{discovery})
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
      key, err := module.Key(Client{}, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
