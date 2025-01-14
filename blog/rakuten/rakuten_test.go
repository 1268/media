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
      in: "rakuten.tv/cz/movies/transvulcania-the-people-s-run",
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
         content_id:  "hell-s-kitchen-usa-15-1",
      },
   },
}

func TestStreamCz(t *testing.T) {
   for _, test := range web_tests {
      if test.out.market_code == "cz" {
         var video on_demand
         video.ClassificationId = classification_id[test.out.market_code]
         video.ContentId = test.out.content_id
         info, err := video.streamings()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%+v\n", info)
         time.Sleep(time.Second)
      }
   }
}

func TestStreamFr(t *testing.T) {
   for _, test := range web_tests {
      if test.out.market_code == "fr" {
         var video on_demand
         video.ClassificationId = classification_id[test.out.market_code]
         video.ContentId = test.out.content_id
         info, err := video.streamings()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%+v\n", info)
         time.Sleep(time.Second)
      }
   }
}
func TestMetadata(t *testing.T) {
   for _, test := range web_tests {
      s, err := test.out.get_season()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n\n", s)
      time.Sleep(time.Second)
   }
}
func TestAddress(t *testing.T) {
   for _, test := range web_tests {
      var out address
      err := out.Set(test.in)
      if err != nil {
         t.Fatal(err)
      }
      if out != test.out {
         t.Fatal(test)
      }
   }
}
