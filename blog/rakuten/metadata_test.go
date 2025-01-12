package rakuten

import (
   "fmt"
   "testing"
   "time"
)

var web_tests = []struct {
   in  string
   out address
}{
   {
      in:  "rakuten.tv/cz/movies/transvulcania-the-people-s-run",
      out: address{market_code: "cz", movie: "transvulcania-the-people-s-run"},
   },
   {
      in:  "rakuten.tv/fr/movies/infidele",
      out: address{market_code: "fr", movie: "infidele"},
   },
   {
      in: "rakuten.tv/uk/player/episodes/stream/hell-s-kitchen-usa-15/hell-s-kitchen-usa-15-1",
      out: address{
         market_code: "uk",
         season:      "hell-s-kitchen-usa-15",
         episode:     "hell-s-kitchen-usa-15-1",
      },
   },
}

func TestStreaming(t *testing.T) {
   for _, test := range web_tests {
      movie, err := test.out.gizmo_movie()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", movie)
      time.Sleep(time.Second)
   }
}
