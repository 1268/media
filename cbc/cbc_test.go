package cbc

import (
   "2a.pages.dev/mech"
   "fmt"
   "os"
   "strings"
   "testing"
)

func Test_New_Profile(t *testing.T) {
   home, err := mech.Home()
   if err != nil {
      t.Fatal(err)
   }
   user, err := user_info(home + "/cbc.txt")
   if err != nil {
      t.Fatal(err)
   }
   tok, err := New_Token(user[0], user[1])
   if err != nil {
      t.Fatal(err)
   }
   pro, err := tok.Profile()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", pro)
}

func user_info(name string) ([]string, error) {
   text, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   return strings.Split(string(text), "\n"), nil
}

func Test_Profile(t *testing.T) {
   home, err := mech.Home()
   if err != nil {
      t.Fatal(err)
   }
   user, err := user_info(home + "/cbc.txt")
   if err != nil {
      t.Fatal(err)
   }
   tok, err := New_Token(user[0], user[1])
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
