package phone

import (
   "os"
   "testing"
)

func TestMovie(t *testing.T) {
   resp, err := movie()
   if err != nil {
      t.Fatal(err)
   }
   resp.Write(os.Stdout)
}
