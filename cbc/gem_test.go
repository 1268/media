package cbc

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
   "time"
)

var links = []string{
   // gem.cbc.ca/media/downton-abbey/s01e05
   "downton-abbey/s01e05",
   // gem.cbc.ca/the-fall/s02e03
   "the-fall/s02e03",
   // gem.cbc.ca/the-witch
   "the-witch",
}

func Test_Gem(t *testing.T) {
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   for _, link := range links {
      gem, err := new_catalog_gem(link)
      if err != nil {
         t.Fatal(err)
      }
      if err := enc.Encode(gem); err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", gem.item())
      //fmt.Println(asset.Name())
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

