package kanopy

import (
   "fmt"
   "net/http"
   "os"
   "os/exec"
   "strings"
   "testing"
)

type transport struct{}

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   fmt.Println(req.URL)
   return http.DefaultTransport.RoundTrip(req)
}

func TestMain(m *testing.M) {
   http.DefaultClient.Transport = transport{}
   m.Run()
}

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
