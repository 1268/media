package youtube

import (
   "fmt"
   "testing"
)

func Test_Format(t *testing.T) {
   var r Request
   r.Android()
   r.Video_ID = androids[0]
   play, err := r.Player(nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play)
}
