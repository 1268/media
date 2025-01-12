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
      out: address{
         market_code: "cz", content_id: "transvulcania-the-people-s-run",
      },
   },
   {
      in:  "rakuten.tv/fr/movies/infidele",
      out: address{market_code: "fr", content_id: "infidele"},
   },
   {
      in: "rakuten.tv/uk/player/episodes/stream/hell-s-kitchen-usa-15/hell-s-kitchen-usa-15-1",
      out: address{
         market_code: "uk",
         season:      "hell-s-kitchen-usa-15",
         content_id:     "hell-s-kitchen-usa-15-1",
      },
   },
}

func TestStreamFr(t *testing.T) {
   for _, test := range web_tests {
      if test.out.market_code == "fr" {
         info, err := on_demand{ContentId: test.out.content_id}.streamings()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%+v\n", info)
         time.Sleep(time.Second)
      }
   }
}

func TestStreamCz(t *testing.T) {
   for _, test := range web_tests {
      if test.out.market_code == "cz" {
         info, err := on_demand{ContentId: test.out.content_id}.streamings()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%+v\n", info)
         time.Sleep(time.Second)
      }
   }
}
