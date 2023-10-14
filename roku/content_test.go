package roku

import (
   "154.pages.dev/http"
   "154.pages.dev/stream"
   "fmt"
   "testing"
   "time"
)

func Test_Content(t *testing.T) {
   http.No_Location()
   http.Verbose()
   for _, test := range tests {
      con, err := New_Content(test.playback_ID)
      if err != nil {
         t.Fatal(err)
      }
      name, err := stream.Format_Film(con)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}

const (
   episode = iota
   movie
)

var tests = map[int]struct {
   key string
   playback_ID string
   pssh string
} {
   // therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76
   episode: {
      key: "e258b67d75420066c8424bd142f84565",
      playback_ID: "105c41ea75775968b670fbb26978ed76",
      pssh: "AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQvfpNbNs5cC5baB+QYX+afhoKaW50ZXJ0cnVzdCIBKg==",
   },
   // therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
   movie: {
      key: "13d7c7cf295444944b627ef0ad2c1b3c",
      playback_ID: "597a64a4a25c5bf6af4a8c7053049a6f",
   },
}
