package paramount

import (
   "testing"
   "time"
)

func Test_Secrets(t *testing.T) {
   for _, secret := range app_secrets {
      token, err := app_token_with(secret)
      if err != nil {
         t.Fatal(err)
      }
      if _, err := token.Item(tests[0].content_ID); err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
