package cineMember

import (
   "41.neocities.org/media/cineMember/user"
   "41.neocities.org/text"
   "fmt"
   "os"
   "testing"
)

// cinemember.nl/films/american-hustle
var american_hustle = Address{"films/american-hustle"}

func TestArticle(t *testing.T) {
   var article OperationArticle
   data, err := article.Marshal(&american_hustle)
   if err != nil {
      t.Fatal(err)
   }
   err = article.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", article)
   fmt.Printf("%q\n", text.Name(&article))
}
