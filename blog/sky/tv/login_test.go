package tv

import (
   "41.neocities.org/x/http"
   "log"
   "os"
   "os/exec"
   "strings"
   "testing"
)

func TestLoginPage(t *testing.T) {
   var transport http.Transport
   transport.ProxyFromEnvironment()
   transport.DefaultClient()
   log.SetFlags(log.Ltime)
   data, err := exec.Command("password", "sky.ch").Output()
   if err != nil {
      t.Fatal(err)
   }
   username, password, _ := strings.Cut(string(data), ":")
   var login login_page
   err = login.New()
   if err != nil {
      t.Fatal(err)
   }
   resp, err := login.login_page(username, password)
   if err != nil {
      t.Fatal(err)
   }
   resp.Write(os.Stdout)
}
