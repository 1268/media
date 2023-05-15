package cbc

import (
   "2a.pages.dev/mech"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Gem(t *testing.T) {
   for _, link := range links {
      gem, err := new_catalog_gem(link)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", gem.item())
      name, err := mech.Name(gem.Structured_Metadata)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}

var links = []string{
   // gem.cbc.ca/media/downton-abbey/s01e05
   "downton-abbey/s01e05",
   // gem.cbc.ca/the-fall/s02e03
   "the-fall/s02e03",
   // gem.cbc.ca/the-witch
   "the-witch",
}

func Test_Media(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   pro, err := Read_Profile(home + "/mech/cbc.json")
   if err != nil {
      t.Fatal(err)
   }
   gem, err := new_catalog_gem(links[0])
   if err != nil {
      t.Fatal(err)
   }
   media, err := pro.Media(gem.item())
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
