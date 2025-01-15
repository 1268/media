package rakuten

import (
   "fmt"
   "testing"
   "time"
)

func (w *web_test) info() (*stream_info, error) {
   class, _ := w.out.classification_id()
   var content *gizmo_content
   if w.out.season_id != "" {
      season, err := w.out.season(class)
      if err != nil {
         return nil, err
      }
      content, _ = w.out.content(season)
   } else {
      var err error
      content, err = w.out.movie(class)
      if err != nil {
         return nil, err
      }
   }
   return content.hd(class, w.language).streamings()
}

func TestContent(t *testing.T) {
   for _, test := range web_tests {
      class, ok := test.out.classification_id()
      if !ok {
         t.Fatal(test.out)
      }
      var content *gizmo_content
      if test.out.season_id != "" {
         season, err := test.out.season(class)
         if err != nil {
            t.Fatal(err)
         }
         var ok bool
         content, ok = test.out.content(season)
         if !ok {
            t.Fatal(season)
         }
      } else {
         var err error
         content, err = test.out.movie(class)
         if err != nil {
            t.Fatal(err)
         }
      }
      fmt.Print(content, "\n\n")
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
