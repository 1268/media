package youtube

import (
   "2a.pages.dev/rosso/http"
   "fmt"
   "testing"
)

func Test_Format(t *testing.T) {
   play, err := Android().Player(http.Default_Client, androids[0], nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play)
}
