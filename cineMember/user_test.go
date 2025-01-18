package cineMember

import (
   "os"
   "os/exec"
   "strings"
   "testing"
)

func TestAuthenticate(t *testing.T) {
   data, err := exec.Command("password", "cinemember.nl").Output()
   if err != nil {
      t.Fatal(err)
   }
   username, password, _ := strings.Cut(string(data), ":")
   data, err = OperationUser{}.Marshal(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("user.txt", data, os.ModePerm)
}
