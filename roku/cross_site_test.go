package roku

import (
   "fmt"
   "testing"
   "time"
)

func Test_Playback(t *testing.T) {
   site, err := New_Cross_Site()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      play, err := site.Playback(test.playback_ID)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", play)
      time.Sleep(time.Second)
   }
}
