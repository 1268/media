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

const sleep = 99 * time.Millisecond

func Test(t *testing.T) {
   http.DefaultClient.Transport = transport{}
   t.Run("WebToken", TestWebToken)
   t.Run("Url", TestUrl)
   t.Run("VideoPlays", TestVideoPlays)
   t.Run("Wrapper", TestWrapper)
}

func TestUrl(t *testing.T) {
   data, err := os.ReadFile("ignore/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token WebToken
   token.Unmarshal(data)
   member, err := token.Membership()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      plays, err := token.Plays(member, test.video_id)
      if err != nil {
         t.Fatal(err)
      }
      manifest, ok := plays.Dash()
      if !ok {
         t.Fatal("VideoPlays.Dash")
      }
      data, err = manifest.Url.Get()
      if err != nil {
         t.Fatal(err)
      }
      if !strings.HasPrefix(string(data), `<?xml version="1.0"?>`) {
         t.Fatal(manifest.Url)
      }
      time.Sleep(sleep)
   }
}

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
      key_id:   "sYcEuBtnTH6Bqn65yIE0Ww==",
      url:      "kanopy.com/en/product/14881167",
      video_id: 14881167,
   },
}

func TestVideoPlays(t *testing.T) {
   data, err := os.ReadFile("ignore/token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token WebToken
   token.Unmarshal(data)
   member, err := token.Membership()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      play, err := token.Plays(member, test.video_id)
      if err != nil {
         t.Fatal(err)
      }
      _, ok := play.Dash()
      if !ok {
         t.Fatal("VideoPlays.Dash")
      }
      time.Sleep(sleep)
   }
   _, ok := VideoPlays{}.Dash()
   if ok {
      t.Fatal("VideoPlays.Dash")
   }
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
   var token WebToken
   token.Unmarshal(data)
   member, err := token.Membership()
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
      plays, err := token.Plays(member, test.video_id)
      if err != nil {
         t.Fatal(err)
      }
      manifest, ok := plays.Dash()
      if !ok {
         t.Fatal("VideoPlays.Dash")
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
      _, err = Wrapper{manifest, &token}.Wrap(data)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(sleep)
   }
}

func TestWebToken(t *testing.T) {
   var token WebToken
   t.Run("Marshal", func(t *testing.T) {
      data, err := exec.Command("password", "kanopy.com").Output()
      if err != nil {
         t.Fatal(err)
      }
      email, password, _ := strings.Cut(string(data), ":")
      data, err = token.Marshal(email, password)
      if err != nil {
         t.Fatal(err)
      }
      os.WriteFile("ignore/token.txt", data, os.ModePerm)
   })
   t.Run("Unmarshal", func(t *testing.T) {
      data, err := os.ReadFile("ignore/token.txt")
      if err != nil {
         t.Fatal(err)
      }
      err = token.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
   })
   var member *Membership
   t.Run("Membership", func(t *testing.T) {
      var err error
      member, err = token.Membership()
      if err != nil {
         t.Fatal(err)
      }
   })
   t.Run("Plays", func(t *testing.T) {
      for _, test := range tests {
         _, err := token.Plays(member, test.video_id)
         if err != nil {
            t.Fatal(err)
         }
         time.Sleep(sleep)
      }
   })
}
