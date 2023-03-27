package youtube

import (
   "fmt"
   "testing"
   "time"
)

const prompt = `1. Go to
%v

2. Enter this code
%v
`

func Test_Code(t *testing.T) {
   code, err := New_Device_Code()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf(prompt, code.Verification_URL, code.User_Code)
   for range [9]bool{} {
      time.Sleep(9 * time.Second)
      tok, err := code.Token()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", tok)
      if tok.Access_Token != "" {
         break
      }
   }
}
