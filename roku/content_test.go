package roku

import (
   "fmt"
   "testing"
)

func Test_Video(t *testing.T) {
   test := tests[key{episode, false}]
   con, err := New_Content(test.playback_ID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(con)
   fmt.Printf("%+v\n", con.DASH())
}
