package kanopy

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "net/http"
   "os"
   "os/exec"
   "strings"
   "testing"
   "time"
)

var tests = []struct {
   key_id   string
   url      string
   video_id int
}{
   {
      key_id:   "DUCS1DH4TB6Po1oEkG9xUA==",
      url:      "kanopy.com/en/product/13808102",
      video_id: 13808102,
   },
   {
      url:      "kanopy.com/en/product/14881167",
      video_id: 14881167,
   },
}

func TestMain(m *testing.M) {
   http.DefaultClient.Transport = transport{}
   m.Run()
}

func TestVideoPlays(t *testing.T) {
   data, err := os.ReadFile("ignore/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token web_token
   token.unmarshal(data)
   member, err := token.membership()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      play, err := token.plays(member, test.video_id)
      if err != nil {
         t.Fatal(err)
      }
      _, ok := play.dash()
      if !ok {
         t.Fatal("video_plays.dash")
      }
      time.Sleep(99 * time.Millisecond)
   }
   _, ok := video_plays{}.dash()
   if ok {
      t.Fatal("video_plays.dash")
   }
}

func TestWebToken(t *testing.T) {
   var token web_token
   t.Run("marshal", func(t *testing.T) {
      data, err := exec.Command("password", "kanopy.com").Output()
      if err != nil {
         t.Fatal(err)
      }
      email, password, _ := strings.Cut(string(data), ":")
      data, err = token.marshal(email, password)
      if err != nil {
         t.Fatal(err)
      }
      os.WriteFile("ignore/token.txt", data, os.ModePerm)
   })
   t.Run("unmarshal", func(t *testing.T) {
      data, err := os.ReadFile("ignore/token.txt")
      if err != nil {
         t.Fatal(err)
      }
      err = token.unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
   })
   var member *membership
   t.Run("membership", func(t *testing.T) {
      var err error
      member, err = token.membership()
      if err != nil {
         t.Fatal(err)
      }
   })
   t.Run("plays", func(t *testing.T) {
      for _, test := range tests {
         _, err := token.plays(member, test.video_id)
         if err != nil {
            t.Fatal(err)
         }
         time.Sleep(99 * time.Millisecond)
      }
   })
}

type transport struct{}

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   fmt.Println(req.URL)
   return http.DefaultTransport.RoundTrip(req)
}

func TestWrapper(t *testing.T) {
   data, err := os.ReadFile("ignore/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token web_token
   token.unmarshal(data)
   member, err := token.membership()
   if err != nil {
      t.Fatal(err)
   }
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
   for _, test := range tests {
      if test.key_id == "" {
         continue
      }
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
      _, err = wrapper{manifest, &token}.Wrap(data)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
