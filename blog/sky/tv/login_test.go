package tv

import (
   "41.neocities.org/x/http"
   "log"
   "os"
   "os/exec"
   "strings"
   "testing"
)

func TestWrite(t *testing.T) {
   var port http.Transport
   port.ProxyFromEnvironment()
   port.DefaultClient()
   log.SetFlags(log.Ltime)
   data, err := exec.Command("password", "sky.ch").Output()
   if err != nil {
      t.Fatal(err)
   }
   username, password, _ := strings.Cut(string(data), ":")
   var login1 login
   err = login1.New()
   if err != nil {
      t.Fatal(err)
   }
   cookies1, err := login1.login(username, password)
   if err != nil {
      t.Fatal(err)
   }
   session, ok := cookies1.session_id()
   if !ok {
      t.Fatal("session_id")
   }
   os.WriteFile("session.txt", []byte(session.String()), os.ModePerm)
}

func TestRead(t *testing.T) {
   data, err := os.ReadFile("session.txt")
   if err != nil {
      t.Fatal(err)
   }
   var session cookie
   err = session.Set(string(data))
   if err != nil {
      t.Fatal(err)
   }
   log.Print(session)
}
