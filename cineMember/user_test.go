package cineMember

import (
   "os"
   "strings"
   "testing"
)

func TestAuthenticate(t *testing.T) {
   username, password, ok := strings.Cut(os.Getenv("cineMember"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   data, err := (*OperationUser).Marshal(nil, username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("user.txt", data, os.ModePerm)
}
