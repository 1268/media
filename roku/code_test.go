package roku

import (
   "fmt"
   "os"
   "testing"
)

func TestCode(t *testing.T) {
   // AccountAuth
   var auth AccountAuth
   data, err := auth.Marshal(nil)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("auth.txt", data, os.ModePerm)
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   // AccountCode
   var code AccountCode
   data, err = code.Marshal(&auth)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.txt", data, os.ModePerm)
   err = code.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
}
