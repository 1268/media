package cineMember

import (
   "fmt"
   "os"
   "reflect"
   "strings"
   "testing"
)

func TestAuthenticate(t *testing.T) {
   username, password, ok := strings.Cut(os.Getenv("cine_member"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   data, err := (*OperationUser).Marshal(nil, username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("user.txt", data, os.ModePerm)
}
