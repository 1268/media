package draken

import (
   "41.neocities.org/widevine"
   "bytes"
   "encoding/base64"
   "fmt"
   "os"
   "os/exec"
   "strings"
   "testing"
   "time"
)

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
   data, err := os.ReadFile("login.txt")
   if err != nil {
      t.Fatal(err)
   }
   var login AuthLogin
   err = login.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, film := range films {
      var movie FullMovie
      err = movie.New(film.custom_id)
      if err != nil {
         t.Fatal(err)
      }
      title, err := login.Entitlement(&movie)
      if err != nil {
         t.Fatal(err)
      }
      play, err := login.Playback(&movie, title)
      if err != nil {
         t.Fatal(err)
      }
      key_id, err := base64.StdEncoding.DecodeString(film.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var pssh widevine.PsshData
      pssh.KeyIds = [][]byte{key_id}
      pssh.ContentId, err = base64.StdEncoding.DecodeString(film.content_id)
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
      data, err = Wrapper{&login, play}.Wrap(data)
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
      time.Sleep(time.Second)
   }
}
func TestLogin(t *testing.T) {
   data, err := exec.Command("password", "drakenfilm.se").Output()
   if err != nil {
      t.Fatal(err)
   }
   username, password, _ := strings.Cut(string(data), ":")
   data, err = new(AuthLogin).Marshal(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.txt", data, os.ModePerm)
}

var films = []struct {
   content_id string
   custom_id  string
   key_id     string
   url        string
}{
   {
      content_id: "MjNkM2MxYjYtZTA0ZC00ZjMyLWIwYTYtOTgxYzU2MTliNGI0",
      custom_id:  "moon",
      key_id:     "74/ZQoQJukeOkUjy76DE+Q==",
      url:        "drakenfilm.se/film/moon",
   },
   {
      content_id: "MTcxMzkzNTctZWQwYi00YTE2LThiZTYtNjllNDE4YzRiYTQw",
      key_id:     "ToV4wH2nlVZE8QYLmLywDg==",
      custom_id:  "the-card-counter",
      url:        "drakenfilm.se/film/the-card-counter",
   },
}

func TestMovie(t *testing.T) {
   for _, film := range films {
      var movie FullMovie
      if err := movie.New(film.custom_id); err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", movie)
      time.Sleep(99 * time.Millisecond)
   }
}
