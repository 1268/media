package tubi

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = []struct {
   content_id int
   key_id     string
   url        string
}{
   {
      content_id: 678877,
      key_id:     "xwRcbVkPTQif73pLihtGkw==",
      url:        "tubitv.com/movies/678877",
   },
   {
      content_id: 200042567,
      key_id:     "Ndopo1ozQ8iSL75MAfbL6A==",
      url:        "tubitv.com/tv-shows/200042567",
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
      content := &VideoContent{}
      data, err := content.Marshal(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      err = content.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      if content.Episode() {
         data, err = content.Marshal(content.SeriesId)
         if err != nil {
            t.Fatal(err)
         }
         err = content.Unmarshal(data)
         if err != nil {
            t.Fatal(err)
         }
         var ok bool
         content, ok = content.Get(test.content_id)
         if !ok {
            t.Fatal("VideoContent.Get")
         }
      }
      video, ok := content.Resource()
      if !ok {
         t.Fatal("VideoContent.Resource")
      }
      key_id, err := base64.StdEncoding.DecodeString(test.key_id)
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
      _, err = video.Wrap(data)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func TestResolution(t *testing.T) {
   for _, test := range tests {
      content := &VideoContent{}
      data, err := content.Marshal(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      err = content.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      if content.Episode() {
         data, err = content.Marshal(content.SeriesId)
         if err != nil {
            t.Fatal(err)
         }
         err = content.Unmarshal(data)
         if err != nil {
            t.Fatal(err)
         }
         var ok bool
         content, ok = content.Get(test.content_id)
         if !ok {
            t.Fatal("VideoContent.Get")
         }
      }
      fmt.Println(test.url)
      for _, v := range content.VideoResources {
         fmt.Println(v.Resolution, v.Type)
      }
      time.Sleep(time.Second)
   }
}
