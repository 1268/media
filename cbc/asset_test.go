package cbc

import (
   "fmt"
   "os"
   "testing"
)

const downton = "downton-abbey/s01e05"

func Test_Media(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   profile, err := Open_Profile(home + "/mech/cbc.json")
   if err != nil {
      t.Fatal(err)
   }
   asset, err := New_Asset(downton)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(asset)
   media, err := profile.Media(asset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
