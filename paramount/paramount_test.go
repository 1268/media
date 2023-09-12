package paramount

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "net/http"
   "os"
   "testing"
   "time"
)

func Test_Post(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   test := tests[0]
   pssh, err := base64.StdEncoding.DecodeString(test.pssh)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, pssh)
   if err != nil {
      t.Fatal(err)
   }
   token, err := New_App_Token()
   if err != nil {
      t.Fatal(err)
   }
   sess, err := token.Session(test.content_ID)
   if err != nil {
      t.Fatal(err)
   }
   key, err := mod.Key(sess)
   if err != nil {
      t.Fatal(err)
   }
   if fmt.Sprintf("%x", key) != test.key {
      t.Fatal(key)
   }
}

func Test_Secrets(t *testing.T) {
   for _, secret := range app_secrets {
      token, err := app_token_with(secret)
      if err != nil {
         t.Fatal(err)
      }
      if _, err := token.Item(tests[0].content_ID); err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}

const (
   episode = iota
   movie
)
var tests = []struct{
   asset func(string)(string,error) // Downloadable
   content int // Movie
   content_ID string
   key string
   pssh string
}{
   {
      // paramountplus.com/shows/video/rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW
      // SEAL Team Season 1 Episode 1: Tip of the Spear
      asset: DASH_CENC,
      content: episode,
      content_ID: "rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW",
      key: "f335e480e47739dbcaae7b48faffc002",
      pssh: "AAAAWHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADgIARIQD3gqa9LyRm65UzN2yiD/XyIgcm4xenlpclZPUGpDbDhyeG9wV3JoVW1KRUlzM0djS1c4AQ==",
   }, {
      // paramountplus.com/movies/video/tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD
      // The SpongeBob Movie: Sponge on the Run
      asset: DASH_CENC,
      content: movie,
      content_ID: "tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD",
   }, {
      // paramountplus.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_
      // 60 Minutes Season 55 Episode 18: 1/15/2023: Star Power, Hide and Seek,
      // The Guru
      asset: Downloadable,
      content: episode,
      content_ID: "YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_",
   },
}

func Test_Media(t *testing.T) {
   for _, test := range tests {
      ref, err := test.asset(test.content_ID)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(ref)
      res, err := http.Get(ref)
      if err != nil {
         t.Fatal(err)
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != http.StatusOK {
         if res.StatusCode != http.StatusFound {
            t.Fatal(res)
         }
      }
      time.Sleep(time.Second)
   }
}

