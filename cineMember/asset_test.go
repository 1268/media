package cineMember

import (
   "fmt"
   "os"
   "testing"
)

const test_url = "cinemember.nl/films/american-hustle"

func TestAsset(t *testing.T) {
   var web Address
   err := web.Set(test_url)
   if err != nil {
      t.Fatal(err)
   }
   var article UserArticle
   data, err := article.Marshal(web)
   if err != nil {
      t.Fatal(err)
   }
   err = article.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   asset, ok := article.Film()
   if !ok {
      t.Fatal("OperationArticle.Film")
   }
   data, err = os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   var user Authenticate
   err = user.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   var play AssetPlay
   data, err = play.Marshal(user, asset)
   if err != nil {
      t.Fatal(err)
   }
   err = play.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play.Dash())
}
