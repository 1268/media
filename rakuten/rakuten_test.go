package rakuten

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "bytes"
   "encoding/base64"
   "errors"
   "fmt"
   "os"
   "strings"
   "testing"
   "time"
)

var tests = []movie_test{
   {
      content_id: "MGU1MTgwMDA2Y2Q1MDhlZWMwMGQ1MzVmZWM2YzQyMGQtbWMtMC0xNDEtMC0w",
      key_id:     "DlGAAGzVCO7ADVNf7GxCDQ==",
      url:        "rakuten.tv/fr/movies/infidele",
   },
   {
      content_id: "OWE1MzRhMWYxMmQ2OGUxYTIzNTlmMzg3MTBmZGRiNjUtbWMtMC0xNDctMC0w",
      key_id:     "mlNKHxLWjhojWfOHEP3bZQ==",
      url:        "rakuten.tv/se/movies/i-heart-huckabees",
   },
}

func (m *movie_test) license() ([]byte, error) {
   var web Address
   web.Set(m.url)
   info, err := web.Hd().Info()
   if err != nil {
      return nil, err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      return nil, err
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      return nil, err
   }
   key_id, err := base64.StdEncoding.DecodeString(m.key_id)
   if err != nil {
      return nil, err
   }
   var pssh widevine.PsshData
   pssh.KeyIds = [][]byte{key_id}
   pssh.ContentId, err = base64.StdEncoding.DecodeString(m.content_id)
   if err != nil {
      return nil, err
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      return nil, err
   }
   data, err := module.RequestBody()
   if err != nil {
      return nil, err
   }
   data, err = info.Wrap(data)
   if err != nil {
      return nil, err
   }
   var body widevine.ResponseBody
   err = body.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   block, err := module.Block(body)
   if err != nil {
      return nil, err
   }
   containers := body.Container()
   for {
      container, ok := containers()
      if !ok {
         return nil, errors.New("ResponseBody.Container")
      }
      if bytes.Equal(container.Id(), key_id) {
         return container.Key(block), nil
      }
   }
}

type movie_test struct {
   content_id string
   key_id     string
   url        string
}

func TestMovie(t *testing.T) {
   for _, test := range tests {
      var web Address
      err := web.Set(test.url)
      if err != nil {
         t.Fatal(err)
      }
      movie, err := web.Movie()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", movie)
      name := text.Name(movie)
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}

func TestFr(t *testing.T) {
   for _, test := range tests {
      if strings.Contains(test.url, "/fr/") {
         var web Address
         web.Set(test.url)
         stream, err := web.Fhd().Info()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%+v\n", stream)
         time.Sleep(time.Second)
      }
   }
}

func TestSe(t *testing.T) {
   for _, test := range tests {
      if strings.Contains(test.url, "/se/") {
         var web Address
         web.Set(test.url)
         stream, err := web.Fhd().Info()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%+v\n", stream)
         time.Sleep(time.Second)
      }
   }
}

func TestLicenseFr(t *testing.T) {
   for _, test := range tests {
      if strings.Contains(test.url, "/fr/") {
         key, err := test.license()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%x\n", key)
         time.Sleep(time.Second)
      }
   }
}

func TestLicenseSe(t *testing.T) {
   for _, test := range tests {
      if strings.Contains(test.url, "/se/") {
         key, err := test.license()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%x\n", key)
         time.Sleep(time.Second)
      }
   }
}
