package cbc

import (
   "fmt"
   "os"
   "testing"
   "time"
)

var ids = []string{
   "downton-abbey/s01e05",
   "the-fall/s02e03",
   "the-witch/s01e01",
}

func Test_Asset(t *testing.T) {
   for _, id := range ids {
      asset, err := New_Asset(id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(asset.Name())
      time.Sleep(time.Second)
   }
}

func Test_Media(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   profile, err := Open_Profile(home + "/mech/cbc.json")
   if err != nil {
      t.Fatal(err)
   }
   asset, err := New_Asset(ids[0])
   if err != nil {
      t.Fatal(err)
   }
   media, err := profile.Media(asset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
