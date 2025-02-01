package main

import (
   "net/http"
   "net/url"
   "os"
)

// no geo block
func main() {
   var req http.Request
   req.Header = http.Header{}
   req.URL = &url.URL{}
   req.URL.Host = "clientapi.prd.sky.ch"
   req.URL.Path = "/stream/2035/MOVIE"
   req.URL.Scheme = "https"
   req.Header["Devicecode"] = []string{"ANDROID_INAPP"}
   // ONLY LAST 5 MINUTE
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxOTM4NDQxIiwiUHJmSWQiOiIyMTE5ODA5IiwiUm9sZXMiOiIiLCJCdW5kbGVzIjoie1wiU2t5U2hvd1wiOlwiU0hPV19QUkVNSVVNXCIsXCJTa3lTcG9ydFwiOlwiXCJ9IiwiZXhwIjoxNzM4MzgzNzA5LCJpc3MiOiJodHRwczovL3d3dy5za3kuY2giLCJhdWQiOiJTa3kgVXNlcnMifQ.oPKRBArSOdo1MnWkSkNx-91hPmf-upxiKnnbFU8gASU"}
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
