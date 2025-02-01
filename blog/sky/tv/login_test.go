package tv

import (
   "fmt"
   "testing"
)

func TestLoginPage(t *testing.T) {
   var login login_page
   err := login.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", login.section)
}
