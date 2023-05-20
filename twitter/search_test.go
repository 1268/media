package twitter

import (
   "2a.pages.dev/rosso/http"
   "fmt"
   "testing"
)

func Test_Search(t *testing.T) {
   http.Default_Client.Log_Level = 2
   s, err := New_Search("filter:spaces")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(s)
}
