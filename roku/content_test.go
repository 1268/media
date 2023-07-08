package roku

import (
   "encoding/json"
   "fmt"
   "mechanize.pages.dev"
   "os"
   "testing"
   "time"
)

func Test_Content(t *testing.T) {
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   for _, test := range tests {
      con, err := New_Content(test.playback_ID)
      if err != nil {
         t.Fatal(err)
      }
      name, err := mechanize.Name(con)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      if err := enc.Encode(con.DASH()); err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
