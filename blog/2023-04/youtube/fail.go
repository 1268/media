package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
   req.Header["X-Goog-Api-Key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if bytes.Contains(body, []byte(`"viewCount"`)) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

var req_body = strings.NewReader(`
{
   "context": {
      "client": {
         "clientName": "ANDROID",
         "clientVersion": "18.14.40"
      }
   },
   "params": "8AEB",
   "videoId": "k5dX9sjXYVk"
}
`)
