package rakuten

import (
   "fmt"
   "testing"
   "time"
)

func TestMetadata(t *testing.T) {
   for _, test := range web_tests {
      b, err := test.out.bravo()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n\n", b)
      time.Sleep(time.Second)
   }
}
