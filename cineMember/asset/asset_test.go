package asset

import (
   "41.neocities.org/media/cineMember"
   "41.neocities.org/media/cineMember/article"
   "41.neocities.org/media/cineMember/user"
   "fmt"
   "os"
   "testing"
)

const test_url = "cinemember.nl/films/american-hustle"

func Test(t *testing.T) {
   var url cineMember.Url
   err := url.Set(test_url)
   if err != nil {
      t.Fatal(err)
   }
   data, err := article.Marshal(url)
   if err != nil {
      t.Fatal(err)
   }
   var art article.Article
   err = art.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   asset, ok := art.Film()
   if !ok {
      t.Fatal("OperationArticle.Film")
   }
   data, err = os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   var auth user.Authenticate
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   data, err = Marshal(auth, asset)
   if err != nil {
      t.Fatal(err)
   }
   var p Play
   err = p.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(p.Dash())
}
