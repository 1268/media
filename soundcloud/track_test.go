package soundcloud

import (
   "fmt"
   "testing"
   "time"
)

type item_type struct {
   id int64
   address string
}

var items = []item_type{
   {936653761, "https://soundcloud.com/kino-scmusic/mqymd53jtwag"},
   {692707328, "https://soundcloud.com/kino-scmusic"},
}

func Test_Resolve(t *testing.T) {
   for _, item := range items {
      tracks, err := Resolve(item.address)
      if err != nil {
         t.Fatal(err)
      }
      for _, track := range tracks {
         fmt.Printf("%+v\n", track)
      }
      time.Sleep(time.Second)
   }
}

func Test_Track(t *testing.T) {
   track, err := New_Track(items[0].id)
   if err != nil {
      t.Fatal(err)
   }
   media, err := track.Progressive()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
