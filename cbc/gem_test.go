package gem

import (
   "154.pages.dev/media"
   "encoding/json"
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
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   pro, err := Read_Profile(home + "/cbc.json")
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   for _, link := range links {
      gem, err := New_Catalog_Gem(link)
      if err != nil {
         t.Fatal(err)
      }
      enc.Encode(gem.Item())
      name, err := media.Name(gem.Structured_Metadata)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      media, err := pro.Media(gem.Item())
      if err != nil {
         t.Fatal(err)
      }
      enc.Encode(media)
      time.Sleep(time.Second)
   }
}
