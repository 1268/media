package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   "https://schnaussandmunk.bandcamp.com",
   "https://schnaussandmunk.bandcamp.com/album/passage-2",
   "https://schnaussandmunk.bandcamp.com/music",
   "https://schnaussandmunk.bandcamp.com/track/amaris-2",
}

func Test_Param(t *testing.T) {
   for _, test := range tests {
      param, err := New_Params(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", param)
      time.Sleep(99 * time.Millisecond)
   }
}
