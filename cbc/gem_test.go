package cbc

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
   "time"
)

var links = []string{
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
      time.Sleep(time.Second)
   }
}
