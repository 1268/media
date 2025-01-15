package rakuten

import "testing"

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
}

type web_test struct {
   a           address
   address     string
   address_out string
   content_id  string
   key_id      string
   language    string
}

var web_tests = []web_test{
   {
      a: address{
         market_code: "fr", content_id: "infidele",
      },
      content_id:  "MGU1MTgwMDA2Y2Q1MDhlZWMwMGQ1MzVmZWM2YzQyMGQtbWMtMC0xNDEtMC0w",
      key_id:      "DlGAAGzVCO7ADVNf7GxCDQ==",
      language:    "ENG",
      address:     "rakuten.tv/fr/movies/infidele",
      address_out: "fr/movies/infidele",
   },
   {
      a: address{
         market_code: "cz", content_id: "transvulcania-the-people-s-run",
      },
      language:    "SPA",
      address:     "rakuten.tv/cz/movies/transvulcania-the-people-s-run",
      address_out: "cz/movies/transvulcania-the-people-s-run",
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
