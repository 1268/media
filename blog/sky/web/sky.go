package main

import (
   "net/http"
   "net/url"
   "fmt"
)

// geo block
func main() {
   var req http.Request
   req.Header = http.Header{}
   req.URL = &url.URL{}
   req.URL.Host = "www.sky.ch"
   req.URL.Scheme = "https"
   req.Header["Cookie"] = []string{"_ASP.NET_SessionId_=gqqn5xtzsrzu3qzw4e5hrgz0"}
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   fmt.Println(resp.Header)
}
