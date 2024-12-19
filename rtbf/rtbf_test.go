package rtbf

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "strings"
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
   data, err := os.ReadFile(home + "/rtbf.txt")
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
      var pssh widevine.PsshData
      pssh.KeyId, err = base64.StdEncoding.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      data, err = module.RequestBody()
      if err != nil {
         t.Fatal(err)
      }
      data, err = title.Wrap(data)
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
            t.Fatal("ResponseBody.Container")
         }
         if bytes.Equal(container.Id(), pssh.KeyId) {
            fmt.Printf("%x\n", container.Decrypt(block))
         }
      }
      time.Sleep(time.Second)
   }
}

var tests = []struct{
   key_id string
   path   string
   url    string
}{
   {
      key_id: "v6f4GpHISrWZcd6esgIvRw==",
      path: "/emission/la-proposition-27866",
      url: "auvio.rtbf.be/emission/la-proposition-27866",
   },
   {
      key_id: "10kWa4A9SOSzHFpq1n1zUQ==",
      path: "/media/grantchester-grantchester-s01-3194636",
      url:  "auvio.rtbf.be/media/grantchester-grantchester-s01-3194636",
   },
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

func TestPage(t *testing.T) {
   for _, test := range tests {
      page, err := Address{test.path}.Page()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", page)
      name := text.Name(&Namer{page})
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}

func TestAccountsLogin(t *testing.T) {
   username, password, ok := strings.Cut(os.Getenv("rtbf"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   data, err := (*AuvioLogin).Marshal(nil, username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.txt", data, os.ModePerm)
}
