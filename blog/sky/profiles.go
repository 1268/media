package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Accept-Encoding"] = []string{"compress"}
   req.Header["Accept-Language"] = []string{"en"}
   req.Header["Api-Version"] = []string{"1.0"}
   req.Header["Appversion"] = []string{"1.18.1.125"}
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxOTM4NDQxIiwiUHJmSWQiOiIyMTE5ODA5IiwiUm9sZXMiOiIiLCJCdW5kbGVzIjoie1wiU2t5U2hvd1wiOlwiU0hPV19QUkVNSVVNXCIsXCJTa3lTcG9ydFwiOlwiXCJ9IiwiZXhwIjoxNzM3OTQ1NjU2LCJpc3MiOiJodHRwczovL3d3dy5za3kuY2giLCJhdWQiOiJTa3kgVXNlcnMifQ.JByE_RXFAJ85ShKKGyyZYlDc0uGYwrQsMiz8amtOyoQ"}
   req.Header["Bundles"] = []string{""}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["Devicecode"] = []string{"ANDROID_INAPP"}
   req.Header["Devicename"] = []string{"Google AOSP on IA Emulator"}
   req.Header["Deviceplatform"] = []string{"ANDROID_INAPP"}
   req.Header["Environment"] = []string{"Sky"}
   req.Header["Macaddress"] = []string{"31f08176-dd2f-43a3-ae94-9e24190600fe"}
   req.Header["Osversion"] = []string{"28"}
   req.Header["Serialnumber"] = []string{"31f08176-dd2f-43a3-ae94-9e24190600fe"}
   req.Header["User-Agent"] = []string{"okhttp/5.0.0-alpha.10"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "gateway.prd.sky.ch"
   req.URL.Path = "/user/profiles"
   req.URL.Scheme = "https"
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
