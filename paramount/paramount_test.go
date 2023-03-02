package paramount

import (
   "2a.pages.dev/mech/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

const (
   dash_cenc = iota
   episode
   hls_clear
   movie
   stream_pack
)

type key struct {
   asset int
   content_type int
}

func Test_Preview(t *testing.T) {
   for _, test := range tests {
      preview, err := New_Preview(test.guid)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", preview)
      time.Sleep(time.Second)
   }
}

var tests = map[key]struct{
   guid string
   key string
   pssh string
} {
   // paramountplus.com/shows/video/rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW
   {dash_cenc, episode}: {
      guid: "rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW",
      key: "f335e480e47739dbcaae7b48faffc002",
      pssh: "AAAAWHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADgIARIQD3gqa9LyRm65UzN2yiD/XyIgcm4xenlpclZPUGpDbDhyeG9wV3JoVW1KRUlzM0djS1c4AQ==",
   },
   // paramountplus.com/movies/video/tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD
   {dash_cenc, movie}: {guid: "tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD"},
   // cbs.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_
   {hls_clear, episode}: {guid: "YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_"},
   // paramountplus.com/shows/video/622520382
   {stream_pack, episode}: {guid: "622520382"},
   // paramountplus.com/movies/video/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_
   {stream_pack, movie}: {guid: "wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_"},
}

func Test_Post(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   test := tests[key{dash_cenc, episode}]
   pssh, err := base64.StdEncoding.DecodeString(test.pssh)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, pssh)
   if err != nil {
      t.Fatal(err)
   }
   sess, err := New_Session(test.guid)
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Post(sess)
   if err != nil {
      t.Fatal(err)
   }
   if keys.Content().String() != test.key {
      t.Fatal(keys)
   }
}
