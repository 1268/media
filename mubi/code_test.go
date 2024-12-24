package mubi

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "bytes"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

func TestWrap(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   data, err := os.ReadFile(home + "/authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   var auth Authenticate
   err = auth.Unmarshal(data)
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
   data, err = auth.Wrap(data)
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
}

func TestFilm(t *testing.T) {
   for i, dogville := range dogvilles {
      var web Address
      err := web.Set(dogville)
      if err != nil {
         t.Fatal(err)
      }
      if i == 0 {
         film, err := web.Film()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Println(text.Name(&Namer{film}))
      }
      fmt.Println(web)
   }
}

// mubi.com/films/190/player
// mubi.com/films/dogville
var dogvilles = []string{
   "/films/dogville",
   "/en/us/films/dogville",
   "/us/films/dogville",
   "/en/films/dogville",
}

var test = struct{
   id int64
   key_id string
   url []string
}{
   id: 325455,
   key_id: "CA215A25BB2D43F0BD095FC671C984EE",
   url: []string{
      "mubi.com/films/325455/player",
      "mubi.com/films/passages-2022",
   },
}
func TestCode(t *testing.T) {
   var code LinkCode
   data, err := code.Marshal()
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
