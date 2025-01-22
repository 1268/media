package nbc

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "bytes"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestWrap(t *testing.T) {
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
   var proxy DrmProxy
   proxy.New()
   for _, test := range tests {
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
      data, err := module.RequestBody()
      if err != nil {
         t.Fatal(err)
      }
      data, err = proxy.Wrap(data)
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
var tests = []struct{
   url string
   program string
   id int
   lock bool
   key_id string
}{
   {
      id: 9000283422,
      key_id: "0552e44842654a4e81b326004be47be0",
      program: "episode",
      url: "nbc.com/saturday-night-live/video/may-18-jake-gyllenhaal/9000283438",
   },
   {
      id: 9000283435,
      key_id: "a48d84f23ec74aa1ba8b1d4c863ac02b",
      lock: true,
      program: "episode",
      url: "nbc.com/saturday-night-live/video/march-30-ramy-youssef/9000283435",
   },
   {
      id: 2957739,
      key_id: "e416811be8c44b8e9e598ea7b22e57cc",
      lock: true,
      program: "movie",
      url: "nbc.com/2-fast-2-furious/video/2-fast-2-furious/2957739",
   },
}

func TestMetadata(t *testing.T) {
   for _, test := range tests {
      var meta Metadata
      err := meta.New(test.id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", text.Name(&meta))
      time.Sleep(time.Second)
   }
}

func TestOnDemand(t *testing.T) {
   for _, test := range tests {
      var meta Metadata
      err := meta.New(test.id)
      if err != nil {
         t.Fatal(err)
      }
      video, err := meta.OnDemand()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", video)
      time.Sleep(time.Second)
   }
}
