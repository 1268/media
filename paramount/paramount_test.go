package paramount

import (
   "fmt"
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
}

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
