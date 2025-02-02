package web

import (
   "41.neocities.org/media/blog/sky/tv"
   "fmt"
   "os"
   "testing"
)

func TestWeb(t *testing.T) {
   data, err := os.ReadFile("../tv/session.txt")
   if err != nil {
      t.Fatal(err)
   }
   var session tv.Cookie
   err = session.Set(string(data))
   if err != nil {
      t.Fatal(err)
   }
   data, err = sky_player(session.Cookie)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(string(data))
}
