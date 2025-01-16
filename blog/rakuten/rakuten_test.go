package rakuten

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   http.DefaultClient.Transport = transport{}
   for _, test := range web_tests {
      class, _ := test.a.classification_id()
      var (
         content *gizmo_content
         err error
      )
      if test.a.season_id != "" {
         season, err := test.a.season(class)
         if err != nil {
            t.Fatal(err)
         }
         content, _ = test.a.content(season)
      } else {
         content, err = test.a.movie(class)
      }
      if err != nil {
         t.Fatal(err)
      }
      if content.String() == "" {
         t.Fatal(content)
      }
      time.Sleep(99 * time.Millisecond)
   }
}

func TestAddress(t *testing.T) {
   for _, test := range web_tests {
      t.Run("Set", func(t *testing.T) {
         var a address
         err := a.Set(test.address)
         if err != nil {
            t.Fatal(err)
         }
         if a != test.a {
            t.Fatal(test)
         }
      })
      t.Run("String", func(t *testing.T) {
         if test.a.String() != test.address_out {
            t.Fatal(test.a)
         }
      })
   }
   t.Run("classification_id", func(t *testing.T) {
      var web address
      _, ok := web.classification_id()
      if ok {
         t.Fatal(web)
      }
   })
}

var web_tests = []web_test{
   {
      a: address{
         market_code: "cz", content_id: "transvulcania-the-people-s-run",
      },
      language:    "SPA",
      address:     "rakuten.tv/cz/movies/transvulcania-the-people-s-run",
      address_out: "cz/movies/transvulcania-the-people-s-run",
   },
   {
      content_id:  "MGU1MTgwMDA2Y2Q1MDhlZWMwMGQ1MzVmZWM2YzQyMGQtbWMtMC0xNDEtMC0w",
      key_id:      "DlGAAGzVCO7ADVNf7GxCDQ==",
      address:     "rakuten.tv/fr/movies/infidele",
      language:    "ENG",
      address_out: "fr/movies/infidele",
      a: address{
         market_code: "fr", content_id: "infidele",
      },
   },
   {
      content_id:  "OWE1MzRhMWYxMmQ2OGUxYTIzNTlmMzg3MTBmZGRiNjUtbWMtMC0xNDctMC0w",
      key_id:      "mlNKHxLWjhojWfOHEP3bZQ==",
      language:    "ENG",
      address:     "rakuten.tv/se/movies/i-heart-huckabees",
      address_out: "se/movies/i-heart-huckabees",
      a: address{
         market_code: "se", content_id: "i-heart-huckabees",
      },
   },
   {
      a: address{
         market_code: "uk",
         season_id:   "hell-s-kitchen-usa-15",
         content_id:  "hell-s-kitchen-usa-15-1",
      },
      language:    "ENG",
      address:     "rakuten.tv/uk/player/episodes/stream/hell-s-kitchen-usa-15/hell-s-kitchen-usa-15-1",
      address_out: "uk/player/episodes/stream/hell-s-kitchen-usa-15/hell-s-kitchen-usa-15-1",
   },
}

type web_test struct {
   a           address
   address     string
   address_out string
   content_id  string
   key_id      string
   language    string
}

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   fmt.Println(req.URL)
   return http.DefaultTransport.RoundTrip(req)
}

type transport struct{}
