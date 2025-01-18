package kanopy

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "os"
   "os/exec"
   "strings"
   "testing"
   "time"
)

func TestLogin(t *testing.T) {
   data, err := exec.Command("password", "kanopy.com").Output()
   if err != nil {
      t.Fatal(err)
   }
   email, password, _ := strings.Cut(string(data), ":")
   data, err = web_token{}.marshal(email, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", data, os.ModePerm)
}

func TestLicense(t *testing.T) {
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
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token web_token
   err = token.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   member, err := token.membership()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      plays, err := token.plays(member, test.video_id)
      if err != nil {
         t.Fatal(err)
      }
      manifest, ok := plays.dash()
      if !ok {
         t.Fatal("video_plays.dash")
      }
      var pssh widevine.PsshData
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
      data, err = module.RequestBody()
      if err != nil {
         t.Fatal(err)
      }
      _, err = poster{manifest, &token}.Wrap(data)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}

var tests = []struct{
   key_id string
   url string
   video_id int64
}{
   {
      key_id: "DUCS1DH4TB6Po1oEkG9xUA==",
      url: "kanopy.com/en/product/13808102",
      video_id: 13808102,
   },
   {
      url: "kanopy.com/en/product/14881167",
      video_id: 14881167,
   },
}
