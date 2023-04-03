package paramount

import (
   "fmt"
   "testing"
   "time"
)

func Test_Secrets(t *testing.T) {
   for _, secret := range app_secrets {
      token, err := app_token_with(secret)
      if err != nil {
         t.Fatal(err)
      }
      item, err := token.Item(tests[0].content_ID)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(item)
      time.Sleep(99 * time.Millisecond)
   }
}
