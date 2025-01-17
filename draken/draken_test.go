package draken

import (
   "41.neocities.org/text"
   "fmt"
   "os"
   "strings"
   "testing"
   "time"
)

var films = []struct {
   content_id string
   custom_id string
   key_id     string
   url        string
}{
   {
      content_id: "MjNkM2MxYjYtZTA0ZC00ZjMyLWIwYTYtOTgxYzU2MTliNGI0",
      custom_id: "moon",
      key_id: "74/ZQoQJukeOkUjy76DE+Q==",
      url: "drakenfilm.se/film/moon",
   },
   {
      content_id: "MTcxMzkzNTctZWQwYi00YTE2LThiZTYtNjllNDE4YzRiYTQw",
      key_id:     "ToV4wH2nlVZE8QYLmLywDg==",
      custom_id:        "the-card-counter",
      url:        "drakenfilm.se/film/the-card-counter",
   },
}

func TestLogin(t *testing.T) {
   username, password, ok := strings.Cut(os.Getenv("draken"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   data, err := (*AuthLogin).Marshal(nil, username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.txt", data, os.ModePerm)
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
