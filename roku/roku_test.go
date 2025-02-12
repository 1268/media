package roku

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestTokenWrite(t *testing.T) {
   // AccountAuth
   data, err := os.ReadFile("auth.txt")
   if err != nil {
      t.Fatal(err)
   }
   var auth AccountAuth
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   // AccountCode
   data, err = os.ReadFile("code.txt")
   if err != nil {
      t.Fatal(err)
   }
   var code AccountCode
   err = code.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   // AccountToken
   data, err = (*AccountToken).Marshal(nil, &auth, &code)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", data, os.ModePerm)
}

func TestTokenRead(t *testing.T) {
   // AccountToken
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token AccountToken
   err = token.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   // AccountAuth
   var auth AccountAuth
   data, err = auth.Marshal(&token)
   if err != nil {
      t.Fatal(err)
   }
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", auth)
}
func TestCode(t *testing.T) {
   // AccountAuth
   var auth AccountAuth
   data, err := auth.Marshal(nil)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("auth.txt", data, os.ModePerm)
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   // AccountCode
   var code AccountCode
   data, err = code.Marshal(&auth)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.txt", data, os.ModePerm)
   err = code.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
}
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
   for _, test := range tests {
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
      key_id, err := base64.StdEncoding.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var pssh widevine.PsshData
      pssh.ContentId, err = base64.StdEncoding.DecodeString(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
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
      _, err = play.Wrap(data)
      if err != nil {
         t.Fatal(err)
      }
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
