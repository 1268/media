package amc

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

func Test_Login(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   account, err := sign_in(home + "/mech/amc.txt")
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Unauth()
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Login(account[0], account[1]); err != nil {
      t.Fatal(err)
   }
   if err := auth.Write_File(home + "/mech/amc.json"); err != nil {
      t.Fatal(err)
   }
}

func Test_Refresh(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Read_Auth(home + "/mech/amc.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Refresh(); err != nil {
      t.Fatal(err)
   }
   if err := auth.Write_File(home + "/mech/amc.json"); err != nil {
      t.Fatal(err)
   }
}
