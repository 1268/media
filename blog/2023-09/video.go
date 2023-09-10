package main

import (
   "net/http"
   "os"
)

var parts = []string{
   "http://redirector.us-east-1.prod-a.boltdns.net/v1/6245817279001/65fb076c-f9eb-4771-bcc7-64a44a7ad4c9/xcf/312f192d-08e0-4d4e-96ab-dddc92524bcd/init.m4f",
   "http://redirector.us-east-1.prod-a.boltdns.net/v1/6245817279001/65fb076c-f9eb-4771-bcc7-64a44a7ad4c9/xcf/312f192d-08e0-4d4e-96ab-dddc92524bcd/segment0.m4f",
}

func main() {
   file, err := os.Create("enc.mp4")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   for _, part := range parts {
      res, err := http.Get(part)
      if err != nil {
         panic(err)
      }
      file.ReadFrom(res.Body)
      res.Body.Close()
   }
}
