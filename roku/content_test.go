package roku

import (
   "2a.pages.dev/mech"
   "fmt"
   "testing"
   "time"
)

func Test_Content(t *testing.T) {
   for _, test := range tests {
      con, err := New_Content(test.playback_ID)
      if err != nil {
         t.Fatal(err)
      }
      name, err := mech.Name(con)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      fmt.Printf("%+v\n", con.DASH())
      time.Sleep(time.Second)
   }
}
