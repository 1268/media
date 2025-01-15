package rakuten

import (
   "41.neocities.org/widevine"
   "bytes"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestStreamFr(t *testing.T) {
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
      if test.out.market_code != "fr" {
         continue
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
      info, err := test.info()
      if err != nil {
         t.Fatal(err)
      }
      data, err = info.Wrap(data)
      if err != nil {
         t.Fatal(err)
      }
      var body widevine.ResponseBody
      err = body.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      block, err := module.Block(body)
      if err != nil {
         t.Fatal(err)
      }
      containers := body.Container()
      for {
         container, ok := containers()
         if !ok {
            break
         }
         if bytes.Equal(container.Id(), key_id) {
            fmt.Printf("%x\n", container.Key(block))
         }
      }
   }
}

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

type web_test struct {
   content_id string
   in       string
   key_id string
   language string
   out      address
}

var web_tests = []web_test{
   {
      in:       "rakuten.tv/fr/movies/infidele",
      content_id: "MGU1MTgwMDA2Y2Q1MDhlZWMwMGQ1MzVmZWM2YzQyMGQtbWMtMC0xNDEtMC0w",
      key_id:     "DlGAAGzVCO7ADVNf7GxCDQ==",
      language: "ENG",
      out:      address{market_code: "fr", content_id: "infidele"},
   },
   {
      in:       "rakuten.tv/cz/movies/transvulcania-the-people-s-run",
      language: "SPA",
      out: address{
         market_code: "cz", content_id: "transvulcania-the-people-s-run",
      },
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
