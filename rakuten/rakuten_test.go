package rakuten

import (
   "41.neocities.org/platform/mullvad"
   "41.neocities.org/widevine"
   "41.neocities.org/x/http"
   "encoding/base64"
   "errors"
   "os"
   "testing"
)

var web_tests = []web_test{
   {
      address:  "rakuten.tv/cz/movies/transvulcania-the-people-s-run",
      content_id: "MzE4ZjdlY2U2OWFmY2ZlM2U5NmRlMzFiZTZiNzcyNzItbWMtMC0xNjQtMC0w",
      key_id: "MY9+zmmvz+PpbeMb5rdycg==",
      language: "SPA",
      location: "cz",
   },
   {
      address:    "rakuten.tv/fr/movies/infidele",
      content_id: "MGU1MTgwMDA2Y2Q1MDhlZWMwMGQ1MzVmZWM2YzQyMGQtbWMtMC0xNDEtMC0w",
      key_id:     "DlGAAGzVCO7ADVNf7GxCDQ==",
      language:   "ENG",
      location:   "fr",
   },
   {
      address:    "rakuten.tv/nl/movies/a-knight-s-tale",
      content_id: "MGJlNmZmYWRhMzY2NjNhMGExNzMwODYwN2U3Y2ZjYzYtbWMtMC0xMzctMC0w",
      key_id: "C+b/raNmY6Chcwhgfnz8xg==",
      language:   "ENG",
      location:   "nl",
   },
   {
      address:  "rakuten.tv/pl/movies/ad-astra",
      content_id: "YTk1MjMzMDI1NWFiOWJmZmIxYTAwZTk3ZDA1ZTBhZjItbWMtMC0xMzctMC0w",
      key_id: "qVIzAlWrm/+xoA6X0F4K8g==",
      language: "ENG",
      location: "pl",
   },
   {
      address:    "rakuten.tv/se/movies/i-heart-huckabees",
      content_id: "OWE1MzRhMWYxMmQ2OGUxYTIzNTlmMzg3MTBmZGRiNjUtbWMtMC0xNDctMC0w",
      key_id:     "mlNKHxLWjhojWfOHEP3bZQ==",
      language:   "ENG",
      location:   "se",
   },
   {
      address:  "rakuten.tv/uk/player/episodes/stream/hell-s-kitchen-usa-15/hell-s-kitchen-usa-15-1",
      content_id: "YmI5NGE0YTA0MTdkMjYyY2MzMGMyZjIzODExNmQ2NzktbWMtMC0xMzktMC0w",
      key_id: "u5SkoEF9JizDDC8jgRbWeQ==",
      language: "ENG",
      location: "gb",
   },
}

type content_class struct {
   g     *GizmoContent
   class int
}

func TestAddress(t *testing.T) {
   for _, test := range web_tests {
      var web Address
      t.Run("Set", func(t *testing.T) {
         err := web.Set(test.address)
         if err != nil {
            t.Fatal(err)
         }
      })
      t.Run("String", func(t *testing.T) {
         if web.String() == "" {
            t.Fatal(test)
         }
      })
   }
   t.Run("classification_id", func(t *testing.T) {
      var web Address
      _, ok := web.ClassificationId()
      if ok {
         t.Fatal(web)
      }
   })
}

func TestContent(t *testing.T) {
   for _, test := range web_tests {
      content, err := test.content()
      if err != nil {
         t.Fatal(err)
      }
      t.Run("String", func(t *testing.T) {
         if content.g.String() == "" {
            t.Fatal(content.g)
         }
      })
      t.Run("Fhd", func(t *testing.T) {
         _, err = content.g.Fhd(content.class, test.language).Streamings()
         if err == nil {
            t.Fatal(content.g)
         }
      })
      func() {
         err := mullvad.Connect(test.location)
         if err != nil {
            t.Fatal(err)
         }
         defer mullvad.Disconnect()
         t.Run("Hd", func(t *testing.T) {
            _, err = content.g.Hd(content.class, test.language).Streamings()
            if err != nil {
               t.Fatal(err)
            }
         })
      }()
   }
}

type web_test struct {
   address    string
   content_id string
   key_id     string
   language   string
   location   string
}
func TestMain(m *testing.M) {
   http.Transport{}.DefaultClient()
   m.Run()
}

func TestStreamInfo(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range web_tests {
      content, err := test.content()
      if err != nil {
         t.Fatal(err)
      }
      func() {
         err := mullvad.Connect(test.location)
         if err != nil {
            t.Fatal(err)
         }
         defer mullvad.Disconnect()
         info, err := content.g.Hd(content.class, test.language).Streamings()
         if err != nil {
            t.Fatal(err)
         }
         var pssh widevine.PsshData
         pssh.ContentId, err = base64.StdEncoding.DecodeString(test.content_id)
         if err != nil {
            t.Fatal(err)
         }
         key_id, err := base64.StdEncoding.DecodeString(test.key_id)
         if err != nil {
            t.Fatal(err)
         }
         pssh.KeyIds = [][]byte{key_id}
         var module widevine.Cdm
         err = module.New(private_key, client_id, pssh.Marshal())
         if err != nil {
            t.Fatal(err)
         }
         data, err := module.RequestBody()
         if err != nil {
            t.Fatal(err)
         }
         _, err = info.Wrap(data)
         if err != nil {
            t.Fatal(err)
         }
      }()
   }
}

func (w *web_test) content() (*content_class, error) {
   var web Address
   web.Set(w.address)
   var content content_class
   content.class, _ = web.ClassificationId()
   if web.SeasonId != "" {
      season, err := web.Season(content.class)
      if err != nil {
         return nil, err
      }
      _, ok := season.Content(&Address{})
      if ok {
         return nil, errors.New("gizmo_season.content")
      }
      content.g, _ = season.Content(&web)
   } else {
      var err error
      content.g, err = web.Movie(content.class)
      if err != nil {
         return nil, err
      }
   }
   return &content, nil
}
