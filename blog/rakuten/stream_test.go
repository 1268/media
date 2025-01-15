package rakuten

import (
   "fmt"
   "testing"
)

func (w *web_test) info() ([]stream_info, error) {
   var content *gizmo_content
   if w.out.season_id != "" {
      season, err := w.out.season()
      if err != nil {
         return nil, err
      }
      content, _ = w.out.content(season)
   } else {
      var err error
      content, err = w.out.movie()
      if err != nil {
         return nil, err
      }
   }
   return w.out.streamings(content, w.language)
}

type web_test struct {
   in       string
   language string
   out      address
}

var web_tests = []web_test{
   {
      in:       "rakuten.tv/cz/movies/transvulcania-the-people-s-run",
      language: "SPA",
      out: address{
         market_code: "cz", content_id: "transvulcania-the-people-s-run",
      },
   },
   {
      in:       "rakuten.tv/fr/movies/infidele",
      language: "ENG",
      out:      address{market_code: "fr", content_id: "infidele"},
   },
   {
      in:       "rakuten.tv/uk/player/episodes/stream/hell-s-kitchen-usa-15/hell-s-kitchen-usa-15-1",
      language: "ENG",
      out: address{
         market_code: "uk",
         season_id:   "hell-s-kitchen-usa-15",
         content_id:  "hell-s-kitchen-usa-15-1",
      },
   },
}

func TestStreamUk(t *testing.T) {
   for _, test := range web_tests {
      if test.out.market_code == "uk" {
         info, err := test.info()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%+v\n", info)
      }
   }
}

func TestStreamCz(t *testing.T) {
   for _, test := range web_tests {
      if test.out.market_code == "cz" {
         info, err := test.info()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%+v\n", info)
      }
   }
}

func TestStreamFr(t *testing.T) {
   for _, test := range web_tests {
      if test.out.market_code == "fr" {
         info, err := test.info()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%+v\n", info)
      }
   }
}
