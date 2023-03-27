package youtube

import (
   "fmt"
   "testing"
)

func Test_Format(t *testing.T) {
   play, err := Android().Player(androids[0], nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play)
}
