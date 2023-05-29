package cbc

import (
   "2a.pages.dev/mech"
   "fmt"
   "os"
   "testing"
   "time"
)

var links = []string{
   "https://gem.cbc.ca/downton-abbey/s01e05",
   "https://gem.cbc.ca/the-fall/s02e03",
   "https://gem.cbc.ca/the-witch",
}

func Test_Gem(t *testing.T) {
   for _, link := range links {
      gem, err := New_Catalog_Gem(link)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", gem.Item())
      name, err := mech.Name(gem.Structured_Metadata)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
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
   gem, err := New_Catalog_Gem(links[0])
   if err != nil {
      t.Fatal(err)
   }
   media, err := pro.Media(gem.Item())
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
