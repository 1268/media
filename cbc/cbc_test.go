package cbc

import (
   "os"
   "testing"
)

func Test_Profile(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   login, err := New_Login(email, password)
   if err != nil {
      t.Fatal(err)
   }
   web, err := login.Web_Token()
   if err != nil {
      t.Fatal(err)
   }
   top, err := web.Over_The_Top()
   if err != nil {
      t.Fatal(err)
   }
   profile, err := top.Profile()
   if err != nil {
      t.Fatal(err)
   }
   if err := profile.Create(home + "/mech/cbc.json"); err != nil {
      t.Fatal(err)
   }
}
