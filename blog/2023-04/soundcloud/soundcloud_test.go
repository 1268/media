package soundcloud

import (
   "fmt"
   "testing"
   "time"
)

const address = "https://soundcloud.com/kino-scmusic/mqymd53jtwag"

func Test_Resolve(t *testing.T) {
   for _, client_ID := range client_IDs {
      track, err := Resolve(address, client_ID)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(track)
      time.Sleep(time.Second)
   }
}
