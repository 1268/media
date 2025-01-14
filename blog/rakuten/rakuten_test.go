package rakuten

import (
   "fmt"
   "testing"
   "time"
)

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

var web_tests = []struct {
   in  string
   out address
}{
   {
      in: "rakuten.tv/uk/player/episodes/stream/hell-s-kitchen-usa-15/hell-s-kitchen-usa-15-1",
      out: address{
         market_code: "uk",
         season_id:   "hell-s-kitchen-usa-15",
         content_id:  "hell-s-kitchen-usa-15-1",
      },
   },
   {
      in:  "rakuten.tv/fr/movies/infidele",
      out: address{market_code: "fr", content_id: "infidele"},
   },
   {
      in: "rakuten.tv/cz/movies/transvulcania-the-people-s-run",
      out: address{
         market_code: "cz", content_id: "transvulcania-the-people-s-run",
      },
   },
}

func TestContent(t *testing.T) {
   for _, test := range web_tests {
      var content *gizmo_content
      if test.out.season_id != "" {
         season, err := test.out.season()
         if err != nil {
            t.Fatal(err)
         }
         var ok bool
         content, ok = season.content(&test.out)
         if !ok {
            t.Fatal(season)
         }
      } else {
         var err error
         content, err = test.out.movie()
         if err != nil {
            t.Fatal(err)
         }
      }
      fmt.Printf("%+v\n\n", content)
      time.Sleep(time.Second)
   }
}

func TestStreamFr(t *testing.T) {
   for _, test := range web_tests {
      if test.out.market_code == "fr" {
         var video on_demand
         video.ContentId = test.out.content_id
         var err error
         video.ClassificationId, err = test.out.classification_id()
         if err != nil {
            t.Fatal(err)
         }
         info, err := video.streamings()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%+v\n", info)
         time.Sleep(time.Second)
      }
   }
}
