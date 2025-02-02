package tv

import (
   "41.neocities.org/x/http"
   "log"
   "os"
   "os/exec"
   "strings"
   "testing"
)

func TestLogin(t *testing.T) {
   var transport http.Transport
   transport.ProxyFromEnvironment()
   transport.DefaultClient()
   log.SetFlags(log.Ltime)
   data, err := exec.Command("password", "sky.ch").Output()
   if err != nil {
      t.Fatal(err)
   }
   username, password, _ := strings.Cut(string(data), ":")
   var show show_login
   err = show.New()
   if err != nil {
      t.Fatal(err)
   }
   resp, err := show.login(username, password)
   if err != nil {
      t.Fatal(err)
   }
   resp.Write(os.Stdout)
}
