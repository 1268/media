package main

import (
   "2a.pages.dev/rosso/protobuf"
   "encoding/base64"
   "fmt"
)

func main() {
   b, err := base64.StdEncoding.DecodeString("8AEB")
   if err != nil {
      panic(err)
   }
   m, err := protobuf.Unmarshal(b)
   if err != nil {
      panic(err)
   }
   fmt.Println(m) // map[30:1]
}
