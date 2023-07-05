package amc

import (
   "2a.pages.dev/mech"
   "os"
   "testing"
)

func Test_Login(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Unauth()
   if err != nil {
      t.Fatal(err)
   }
   u, err := mech.User(home + "/amc-plus/user.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Login(u["username"], u["password"]); err != nil {
      t.Fatal(err)
   }
   if err := auth.Write_File(home + "/amc-plus/auth.json"); err != nil {
      t.Fatal(err)
   }
}
