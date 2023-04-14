package paramount

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

func Test_Media(t *testing.T) {
   for _, test := range tests {
      ref := test.asset(test.content_ID)
      fmt.Println(ref)
      res, err := http.Get(ref)
      if err != nil {
         t.Fatal(err)
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != 200 && res.StatusCode != 302 {
         t.Fatal(res)
      }
      time.Sleep(time.Second)
   }
}
