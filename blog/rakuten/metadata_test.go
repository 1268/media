package rakuten

import (
   "fmt"
   "testing"
   "time"
)

func TestMetadata(t *testing.T) {
   for _, test := range web_tests {
      s, err := test.out.get_season()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n\n", s)
      time.Sleep(time.Second)
   }
}
