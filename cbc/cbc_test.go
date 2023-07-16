package gem

import (
   "154.pages.dev/media"
   "fmt"
   "os"
   "testing"
)

func Test_New_Profile(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := media.User(home + "/cbc/user.json")
   if err != nil {
      t.Fatal(err)
   }
   tok, err := New_Token(u["username"], u["password"])
   if err != nil {
      t.Fatal(err)
   }
   pro, err := tok.Profile()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", pro)
}

func Test_Profile(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := media.User(home + "/cbc/user.json")
   if err != nil {
      t.Fatal(err)
   }
   tok, err := New_Token(u["username"], u["password"])
   if err != nil {
      t.Fatal(err)
   }
   pro, err := tok.Profile()
   if err != nil {
      t.Fatal(err)
   }
   if err := pro.Write_File(home + "/cbc.json"); err != nil {
      t.Fatal(err)
   }
}
