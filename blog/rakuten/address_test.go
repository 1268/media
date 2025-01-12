package rakuten

import "testing"

var web_tests = []struct{
   in string
   out address
}{
   {
      in: "rakuten.tv/fr/movies/infidele",
      out: address{market_code: "fr", movie: "infidele"},
   },
   {
      in: "rakuten.tv/uk/player/episodes/stream/hell-s-kitchen-usa-15/hell-s-kitchen-usa-15-1",
      out: address{
         market_code: "uk",
         season: "hell-s-kitchen-usa-15",
         episode: "hell-s-kitchen-usa-15-1",
      },
   },
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
