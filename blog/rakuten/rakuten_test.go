package rakuten

import "testing"

func TestAddress(t *testing.T) {
   for _, test := range web_tests {
      var a address
      err := a.Set(test.address)
      if err != nil {
         t.Fatal(err)
      }
      if a != test.a {
         t.Fatal(test)
      }
   }
}

func (w *web_test) info() (*stream_info, error) {
   class, _ := w.a.classification_id()
   var content *gizmo_content
   if w.a.season_id != "" {
      season, err := w.a.season(class)
      if err != nil {
         return nil, err
      }
      content, _ = w.a.content(season)
   } else {
      var err error
      content, err = w.a.movie(class)
      if err != nil {
         return nil, err
      }
   }
   return content.hd(class, w.language).streamings()
}

type web_test struct {
   a address
   address string
   content_id string
   key_id     string
   language   string
}

var web_tests = []web_test{
   {
      a:        address{market_code: "fr", content_id: "infidele"},
      address:         "rakuten.tv/fr/movies/infidele",
      content_id: "MGU1MTgwMDA2Y2Q1MDhlZWMwMGQ1MzVmZWM2YzQyMGQtbWMtMC0xNDEtMC0w",
      key_id:     "DlGAAGzVCO7ADVNf7GxCDQ==",
      language:   "ENG",
   },
   {
      a: address{
         market_code: "cz", content_id: "transvulcania-the-people-s-run",
      },
      address:       "rakuten.tv/cz/movies/transvulcania-the-people-s-run",
      language: "SPA",
   },
   {
      a: address{
         market_code: "uk",
         season_id:   "hell-s-kitchen-usa-15",
         content_id:  "hell-s-kitchen-usa-15-1",
      },
      address:       "rakuten.tv/uk/player/episodes/stream/hell-s-kitchen-usa-15/hell-s-kitchen-usa-15-1",
      language: "ENG",
   },
}
