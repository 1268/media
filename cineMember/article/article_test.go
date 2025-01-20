package article

import (
   "41.neocities.org/media/cineMember"
   "fmt"
   "testing"
)

const test = "cinemember.nl/films/american-hustle"

func Test(t *testing.T) {
   var url cineMember.Url
   err := url.Set(test)
   if err != nil {
      t.Fatal(err)
   }
   data, err := Marshal(url)
   if err != nil {
      t.Fatal(err)
   }
   var art Article
   err = art.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", art)
}
