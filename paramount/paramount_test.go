package paramount

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

func Test_Media(t *testing.T) {
   for _, test := range tests {
      ref := test.asset(test.guid)
      fmt.Println(ref)
      res, err := http.Get(ref)
      if err != nil {
         t.Fatal(err)
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != 200 && res.StatusCode != 302 {
         t.Fatal(res)
      }
      time.Sleep(time.Second)
   }
}

const (
   episode = iota
   movie
)

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

var tests = []struct{
   asset func(string)string // Downloadable
   content int // Movie
   guid string
   key string
   pssh string
}{
   { // paramountplus.com/shows/video/rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW
      asset: DASH_CENC,
      content: episode,
      guid: "rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW",
      key: "f335e480e47739dbcaae7b48faffc002",
      pssh: "AAAAWHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADgIARIQD3gqa9LyRm65UzN2yiD/XyIgcm4xenlpclZPUGpDbDhyeG9wV3JoVW1KRUlzM0djS1c4AQ==",
   },
   { // paramountplus.com/movies/video/tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD
      asset: DASH_CENC,
      content: movie,
      guid: "tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD",
   },
   { // paramountplus.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_
      asset: Downloadable,
      content: episode,
      guid: "YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_",
   },
}
