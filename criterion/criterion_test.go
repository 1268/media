package criterion

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "bytes"
   "encoding/hex"
   "fmt"
   "os"
   "strings"
   "testing"
)

func TestLicense(t *testing.T) {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token AuthToken
   err = token.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   item, err := token.Video(video_test.slug)
   if err != nil {
      t.Fatal(err)
   }
   files, err := token.Files(item)
   if err != nil {
      t.Fatal(err)
   }
   file, ok := files.Dash()
   if !ok {
      t.Fatal("VideoFiles.Dash")
   }
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
   var pssh widevine.PsshData
   pssh.KeyId, err = hex.DecodeString(video_test.key_id)
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
   data, err = file.Wrap(data)
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
}

func TestVideo(t *testing.T) {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token AuthToken
   err = token.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   item, err := token.Video(video_test.slug)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
   fmt.Printf("%q\n", text.Name(item))
}

var video_test = struct{
   key_id string
   slug string
   url string
}{
   key_id: "e4576465a745213f336c1ef1bf5d513e",
   slug: "my-dinner-with-andre",
   url: "criterionchannel.com/videos/my-dinner-with-andre",
}
func TestToken(t *testing.T) {
   username, password, ok := strings.Cut(os.Getenv("criterion"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   data, err := (*AuthToken).Marshal(nil, username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", data, os.ModePerm)
}
