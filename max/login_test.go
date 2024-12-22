package max

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
   data, err := os.ReadFile("login.txt")
   if err != nil {
      t.Fatal(err)
   }
   var login LinkLogin
   err = login.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var web Address
      web.UnmarshalText([]byte(test.url))
      play, err := login.Playback(&web)
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

func TestRoutes(t *testing.T) {
   data, err := os.ReadFile("login.txt")
   if err != nil {
      t.Fatal(err)
   }
   var login LinkLogin
   err = login.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var web Address
      web.UnmarshalText([]byte(test.url))
      routes, err := login.Routes(&web)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", routes)
      fmt.Printf("%q\n", text.Name(routes))
      time.Sleep(time.Second)
   }
}

var tests = []struct {
   key_id string
   url        string
   video_type string
}{
   {
      key_id: "0102f83e95fac43cf1662dd8e5b08d90",
      url: "play.max.com/video/watch/c9e9bde1-1463-4c92-a25a-21451f3c5894/f1d899ac-1780-494a-a20d-caee55c9e262",
      video_type: "MOVIE",
   },
   {
      key_id: "0102d949c44f81b28fdb98d535c8bade",
      url:        "play.max.com/video/watch/d0938760-d3ca-4c59-aea2-74ecbed42d17/2e7d1db4-2fd7-47fb-a7c3-a65b7c2e5d6f",
      video_type: "EPISODE",
   },
}
func TestLogin(t *testing.T) {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token BoltToken
   token.St = string(data)
   data, err = (*LinkLogin).Marshal(nil, &token)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.txt", data, os.ModePerm)
}
