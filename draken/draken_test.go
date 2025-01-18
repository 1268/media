package draken

import (
   "41.neocities.org/text"
   "fmt"
   "os"
   "os/exec"
   "strings"
   "testing"
   "time"
)

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
      name := text.Name(&Namer{movie})
      fmt.Printf("%q\n", name)
      time.Sleep(99 * time.Millisecond)
   }
}
