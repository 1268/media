package gem

import (
   "2a.pages.dev/mech"
   "fmt"
   "os"
   "strings"
   "testing"
)

func user(name string) (map[string]string, error) {
   b, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   var m map[string]string
   if err := json.Unmarshal(b, &m); err != nil {
      return nil, err
   }
   return m, nil
}

func Test_New_Profile(t *testing.T) {
   home, err := mech.Home()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user(home + "/gem.cbc.ca/user.json")
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
   home, err := mech.Home()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user(home + "/gem.cbc.ca/user.json")
   if err != nil {
      t.Fatal(err)
   }
   tok, err := New_Token(u[0], u[1])
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
