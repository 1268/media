package cbc

import (
   "fmt"
   "os"
   "strings"
   "testing"
)

func Test_Profile(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   account, err := sign_in(home + "/mech/cbc.txt")
   if err != nil {
      t.Fatal(err)
   }
   tok, err := new_token(account[0], account[1])
   if err != nil {
      t.Fatal(err)
   }
   pro, err := tok.profile()
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
