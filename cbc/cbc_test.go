package cbc

import (
   "os"
   "strings"
   "testing"
)

func sign_in(name string) ([]string, error) {
   data, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   return strings.Split(string(data), "\n"), nil
}

func Test_Profile(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   account, err := sign_in(home + "/mech/cbc.txt")
   if err != nil {
      t.Fatal(err)
   }
   login, err := New_Login(account[0], account[1])
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
   if err := profile.Write_File(home + "/mech/cbc.json"); err != nil {
      t.Fatal(err)
   }
}
