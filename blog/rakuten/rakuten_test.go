package rakuten

import (
   "41.neocities.org/platform/mullvad"
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/base64"
   "errors"
   "log"
   "net/http"
   "os"
   "testing"
)

type content_class struct {
   g     *gizmo_content
   class int
}

type transport struct{}

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   log.Print(req.URL)
   return http.DefaultTransport.RoundTrip(req)
}

type web_test struct {
   address    string
   content_id string
   key_id     string
   language   string
   location   string
}

var web_tests = []web_test{
   {
      address:    "rakuten.tv/fr/movies/infidele",
      content_id: "MGU1MTgwMDA2Y2Q1MDhlZWMwMGQ1MzVmZWM2YzQyMGQtbWMtMC0xNDEtMC0w",
      key_id:     "DlGAAGzVCO7ADVNf7GxCDQ==",
      language:   "ENG",
      location:   "fr",
   },
   {
      language: "ENG",
      address:  "rakuten.tv/uk/player/episodes/stream/hell-s-kitchen-usa-15/hell-s-kitchen-usa-15-1",
      location: "gb",
   },
   {
      address:  "rakuten.tv/cz/movies/transvulcania-the-people-s-run",
      language: "SPA",
      location: "cz",
   },
   {
      content_id: "OWE1MzRhMWYxMmQ2OGUxYTIzNTlmMzg3MTBmZGRiNjUtbWMtMC0xNDctMC0w",
      key_id:     "mlNKHxLWjhojWfOHEP3bZQ==",
      language:   "ENG",
      address:    "rakuten.tv/se/movies/i-heart-huckabees",
      location:   "se",
   },
}

///

func TestAddress(t *testing.T) {
   for _, test := range web_tests {
      var web address
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
      var web address
      _, ok := web.classification_id()
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
      t.Run("fhd", func(t *testing.T) {
         _, err = content.g.fhd(content.class, test.language).streamings()
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
         t.Run("hd", func(t *testing.T) {
            _, err = content.g.hd(content.class, test.language).streamings()
            if err != nil {
               t.Fatal(err)
            }
         })
      }()
   }
}

func TestMain(m *testing.M) {
   http.DefaultClient.Transport = transport{}
   m.Run()
}

func TestNamer(t *testing.T) {
   for _, test := range web_tests {
      content, err := test.content()
      if err != nil {
         t.Fatal(err)
      }
      if text.Name(namer{content.g}) == "" {
         t.Fatal(content)
      }
   }
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
      if test.key_id == "" {
         continue
      }
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
         info, err := content.g.hd(content.class, test.language).streamings()
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
   var web address
   web.Set(w.address)
   var content content_class
   content.class, _ = web.classification_id()
   if web.season_id != "" {
      season, err := web.season(content.class)
      if err != nil {
         return nil, err
      }
      _, ok := season.content(&address{})
      if ok {
         return nil, errors.New("gizmo_season.content")
      }
      content.g, _ = season.content(&web)
   } else {
      var err error
      content.g, err = web.movie(content.class)
      if err != nil {
         return nil, err
      }
   }
   return &content, nil
}
