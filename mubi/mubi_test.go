package mubi

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func Test(t *testing.T) {
   data, err := os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   var auth Authenticate
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, test1 := range tests {
      var secure SecureUrl
      data, err = secure.Marshal(&auth, &Film{Id: test1.id})
      if err != nil {
         t.Fatal(err)
      }
      err = secure.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", secure)
      time.Sleep(time.Second)
   }
}

var tests = []struct{
   id     int64
   key_id string
   url  string
}{
   {
      id: 391623,
      key_id: "",
      url: "mubi.com/films/bird",
   },
   {
      id:     325455,
      key_id: "CA215A25BB2D43F0BD095FC671C984EE",
      url: "mubi.com/films/passages-2022",
   },
}
