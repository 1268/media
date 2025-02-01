package main

import (
   "encoding/base64"
   "encoding/json"
   "fmt"
   "os"
   "strings"
)

func main() {
   data, err := os.ReadFile("sky.json")
   if err != nil {
      panic(err)
   }
   var value struct {
      RefreshToken string
   }
   err = json.Unmarshal(data, &value)
   if err != nil {
      panic(err)
   }
   data, err = base64.StdEncoding.DecodeString(
      strings.Split(value.RefreshToken, ".")[1],
   )
   var value1 struct {
      MacAddress string
   }
   err = json.Unmarshal(data, &value1)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", value1)
}
