package nbc

import (
   "154.pages.dev/media"
   "fmt"
   "testing"
   "time"
)

var guids = []int64{
   // nbc.com/saturday-night-live/video/november-5-amy-schumer/9000258300
   9000258300,
   // nbc.com/pasion-de-gavilanes/video/una-verguenza/3760495
   3760495,
}

func Test_Meta(t *testing.T) {
   for _, guid := range guids {
      meta, err := New_Metadata(guid)
      if err != nil {
         t.Fatal(err)
      }
      name, err := media.Name(meta)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
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
