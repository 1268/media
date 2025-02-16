package rtbf

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = []struct {
   key_id string
   path   string
   url    string
}{
   {
      key_id: "Ma5jT/1dR8K/ljWx/1Pb4A==",
      path:   "/media/titanic-3286058",
      url:    "auvio.rtbf.be/media/titanic-3286058",
   },
   {
      key_id: "xESyRLihQMacu++BvoakfA==",
      path:   "/media/agatha-christie-pourquoi-pas-evans-agatha-christie-pourquoi-pas-evans-3280380",
      url:    "auvio.rtbf.be/media/agatha-christie-pourquoi-pas-evans-agatha-christie-pourquoi-pas-evans-3280380",
   },
}

func TestPage(t *testing.T) {
   for _, test := range tests {
      page, err := Address{test.path}.Page()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", page)
      time.Sleep(time.Second)
   }
}

func TestWebToken(t *testing.T) {
   data, err := os.ReadFile("login.txt")
   if err != nil {
      t.Fatal(err)
   }
   var login AuvioLogin
   err = login.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   token, err := login.Token()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", token)
}

func TestEntitlement(t *testing.T) {
   data, err := os.ReadFile("login.txt")
   if err != nil {
      t.Fatal(err)
   }
   var login AuvioLogin
   err = login.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   token, err := login.Token()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := token.Auth()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      page, err := Address{test.path}.Page()
      if err != nil {
         t.Fatal(err)
      }
      asset_id, ok := page.GetAssetId()
      if !ok {
         t.Fatal("AuvioPage.GetAssetId")
      }
      title, err := auth.Entitlement(asset_id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", title)
      fmt.Println(title.Dash())
      time.Sleep(time.Second)
   }
}
