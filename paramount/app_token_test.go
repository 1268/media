package paramount

import (
   "encoding/json"
   "os"
   "testing"
   "time"
)

func Test_Secrets(t *testing.T) {
   test := tests[0]
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.SetEscapeHTML(false)
   for _, secret := range app_secrets {
      token, err := app_token_with(secret)
      if err != nil {
         t.Fatal(err)
      }
      sess, err := token.Session(test.content_ID)
      if err != nil {
         t.Fatal(err)
      }
      if err := enc.Encode(sess); err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
