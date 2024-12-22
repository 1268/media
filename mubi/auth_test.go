package mubi

import (
   "os"
   "testing"
)

func TestAuthenticate(t *testing.T) {
   data, err := os.ReadFile("code.txt")
   if err != nil {
      t.Fatal(err)
   }
   var code LinkCode
   err = code.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   data, err = (*Authenticate).Marshal(nil, &code)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.txt", data, os.ModePerm)
}
