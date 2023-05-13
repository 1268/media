package cbc

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
   "time"
)

var ids = []string{
   // gem.cbc.ca/media/downton-abbey/s01e05
   "downton-abbey/s01e05",
   // gem.cbc.ca/media/the-fall/s02e03
   "the-fall/s02e03",
   // gem.cbc.ca/media/the-witch/s01e01
   "the-witch/s01e01",
}

func Test_Asset(t *testing.T) {
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   for _, id := range ids {
      asset, err := New_Asset(id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(asset.Name())
      if err := enc.Encode(asset); err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

func Test_Media(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   profile, err := Read_Profile(home + "/mech/cbc.json")
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
