package http

import (
   "net/http"
   "io"
   "strings"
   "testing"
)

func post() ([]byte, error) {
   resp, err := http.Post(
      "http://httpbingo.org/post", "", strings.NewReader("hello world"),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

func TestTransport(t *testing.T) {
   http.DefaultClient.Transport = transport{}
   _, err := post()
   if err != nil {
      t.Fatal(err)
   }
   _, err = post()
   if err != nil {
      t.Fatal(err)
   }
   _, err = http.Head("http://hello.invalid")
   if err == nil {
      t.Fatal(".invalid")
   }
}
