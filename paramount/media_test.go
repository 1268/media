package paramount

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

func Test_Media(t *testing.T) {
   for _, test := range tests {
      ref, err := test.asset(test.content_ID)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(ref)
      res, err := http.Get(ref)
      if err != nil {
         t.Fatal(err)
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != http.StatusOK {
         if res.StatusCode != http.StatusFound {
            t.Fatal(res)
         }
      }
      time.Sleep(time.Second)
   }
}
