package twitter

import (
   "os"
   "testing"
)

func Test_Token(t *testing.T) {
   res, err := token()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
