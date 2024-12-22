package paramount

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "bytes"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestWidevine(t *testing.T) {
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
   var app AppToken
   err = app.ComCbsApp()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      session, err := app.Session(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      var pssh widevine.PsshData
      pssh.ContentId = []byte(test.content_id)
      pssh.KeyId, err = base64.StdEncoding.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      data, err := module.RequestBody()
      if err != nil {
         t.Fatal(err)
      }
      data, err = session.Wrap(data)
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
         if bytes.Equal(container.Id(), pssh.KeyId) {
            fmt.Printf("%x\n", container.Decrypt(block))
         }
      }
      time.Sleep(time.Second)
   }
}

var tests = []struct{
   content_id string
   location string
   url string
   key_id string
}{
   {
      content_id: "Oo75PgAbcmt9xqqn1AMoBAfo190Cfhqi",
      key_id: "3RyyVzthSSOklAXiQ2vyRw==",
      location: "US",
      url: "paramountplus.com/movies/video/Oo75PgAbcmt9xqqn1AMoBAfo190Cfhqi",
   },
   {
      content_id: "esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
      key_id: "H94BVNcqT0WRKzTwzgd36w==",
      location: "US",
      url: "paramountplus.com/shows/video/esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
   },
   {
      content_id: "rZ59lcp4i2fU4dAaZJ_iEgKqVg_ogrIf",
      key_id: "Sryog4HeT2CLHx38NftIMA==",
      location: "US",
      url: "cbs.com/shows/video/rZ59lcp4i2fU4dAaZJ_iEgKqVg_ogrIf",
   },
   {
      content_id: "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
      key_id: "BsO37qHORXefruKryNAaVQ==",
      location: "GB",
      url: "paramountplus.com/movies/video/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
   },
   {
      content_id: "ssc3CuuS4mrQ7EyVXILH0FEQSi5yBAsA",
      key_id: "",
      location: "GB",
      url: "paramountplus.com/shows/video/ssc3CuuS4mrQ7EyVXILH0FEQSi5yBAsA",
   },
   {
      content_id: "WNujiS5PHkY5wN9doNY6MSo_7G8uBUcX",
      key_id: "bsT01+Q1Ta+39TayayKhBg==",
      location: "AU",
      url: "paramountplus.com/shows/video/WNujiS5PHkY5wN9doNY6MSo_7G8uBUcX",
   },
}

func TestMpdUs(t *testing.T) {
   var token AppToken
   err := token.ComCbsApp()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      if test.location == "usa" {
         var item VideoItem
         data, err := item.Marshal(token, test.content_id)
         if err != nil {
            t.Fatal(err)
         }
         err = item.Unmarshal(data)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", item.Mpd())
         time.Sleep(time.Second)
      }
   }
}

func TestMpdFr(t *testing.T) {
   var token AppToken
   err := token.ComCbsCa()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      if test.location == "france" {
         var item VideoItem
         data, err := item.Marshal(token, test.content_id)
         if err != nil {
            t.Fatal(err)
         }
         err = item.Unmarshal(data)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", item.Mpd())
         time.Sleep(time.Second)
      }
   }
}

func TestItemUsa(t *testing.T) {
   var token AppToken
   err := token.ComCbsApp()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      if test.location == "usa" {
         var item VideoItem
         data, err := item.Marshal(token, test.content_id)
         if err != nil {
            t.Fatal(err)
         }
         err = item.Unmarshal(data)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", text.Name(&item))
         time.Sleep(time.Second)
      }
   }
}
