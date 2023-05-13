package cbc

import (
   "fmt"
   "os"
   "strings"
   "testing"
)

func Test_New_Profile(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   account, err := sign_in(home + "/mech/cbc.txt")
   if err != nil {
      t.Fatal(err)
   }
   tok, err := New_Token(account[0], account[1])
   if err != nil {
      t.Fatal(err)
   }
   pro, err := tok.Profile()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", pro)
}

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
   tok, err := New_Token(account[0], account[1])
   if err != nil {
      t.Fatal(err)
   }
   pro, err := tok.Profile()
   if err != nil {
      t.Fatal(err)
   }
   if err := pro.Write_File(home + "/mech/cbc.json"); err != nil {
      t.Fatal(err)
   }
}
