package cbc

import (
   "io"
   "os"
   "testing"
)

func Test_Gem(t *testing.T) {
   res, err := gem("the-fall/s02e03")
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("gem.json", data, 0666)
}
