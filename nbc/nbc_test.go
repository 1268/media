package nbc

import (
   "fmt"
   "testing"
   "time"
)

var guids = []int64{
   // nbc.com/saturday-night-live/video/november-5-amy-schumer/9000258300
   9000258300,
   // nbc.com/pasion-de-gavilanes/video/la-valentia-de-norma/9000221348
   9000221348,
}

func Test_Video(t *testing.T) {
   for _, guid := range guids {
      meta, err := New_Metadata(guid)
      if err != nil {
         t.Fatal(err)
      }
      vid, err := meta.Video()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", vid)
      time.Sleep(time.Second)
   }
}
