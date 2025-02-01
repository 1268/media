package tv

import (
   "os"
   "os/exec"
   "strings"
   "testing"
)

func TestLoginPage(t *testing.T) {
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
