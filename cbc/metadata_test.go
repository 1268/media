package gem

import (
   "154.pages.dev/http"
   "154.pages.dev/stream"
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

func Test_Stream(t *testing.T) {
   http.No_Location()
   http.Verbose()
   for _, link := range links {
      gem, err := New_Catalog_Gem(link)
      if err != nil {
         t.Fatal(err)
      }
      item := gem.Item()
      fmt.Printf("%+v\n", item)
      name, err := stream.Format_Film(gem.Structured_Metadata)
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
   pro, err := Read_Profile(home + "/cbc/profile.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, link := range links {
      gem, err := New_Catalog_Gem(link)
      if err != nil {
         t.Fatal(err)
      }
      media, err := pro.Media(gem.Item())
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", media)
      time.Sleep(time.Second)
   }
}
