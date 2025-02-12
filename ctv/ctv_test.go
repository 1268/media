package ctv

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = []struct{
   content_id string
   key_id string
   path string
   url string
}{
   {
      url: "ctv.ca/movies/fools-rush-in-57470",
      path: "/movies/fools-rush-in-57470",
      key_id: "A98dtspZsb9/z++3IHp0Dw==",
      content_id: "ZmYtOGYyNjEzYWUtNTIxNTAx",
   },
   {
      url: "ctv.ca/shows/friends/the-one-with-the-chicken-pox-s2e23",
      path: "/shows/friends/the-one-with-the-chicken-pox-s2e23",
   },
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
   var pssh widevine.PsshData
   pssh.ContentId, err = base64.StdEncoding.DecodeString(content_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      t.Fatal(err)
   }
   data, err := module.RequestBody()
   if err != nil {
      t.Fatal(err)
   }
   _, err = Wrapper{}.Wrap(data)
   if err != nil {
      t.Fatal(err)
   }
}

func TestMedia(t *testing.T) {
   for _, test_path := range test_paths {
      resolve, err := Address{test_path}.Resolve()
      if err != nil {
         t.Fatal(err)
      }
      axis, err := resolve.Axis()
      if err != nil {
         t.Fatal(err)
      }
      var media MediaContent
      data, err := media.Marshal(axis)
      if err != nil {
         t.Fatal(err)
      }
      err = media.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func TestManifest(t *testing.T) {
   for _, test_path := range test_paths {
      resolve, err := Address{test_path}.Resolve()
      if err != nil {
         t.Fatal(err)
      }
      axis, err := resolve.Axis()
      if err != nil {
         t.Fatal(err)
      }
      var media MediaContent
      data, err := media.Marshal(axis)
      if err != nil {
         t.Fatal(err)
      }
      err = media.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      manifest, err := axis.Manifest(&media)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(string(manifest))
      time.Sleep(time.Second)
   }
}
