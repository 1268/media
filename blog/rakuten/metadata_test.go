package rakuten

import (
   "fmt"
   "testing"
   "time"
)

func TestMetadata(t *testing.T) {
   for _, test := range web_tests {
      meta, err := test.out.metadata()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", meta)
      time.Sleep(time.Second)
   }
}
