package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
   "time"
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
   req.Header["X-Goog-Api-Key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   for i := 1; i <= 9; i++ {
      req.Body = io.NopCloser(strings.NewReader(req_body))
      res, err := new(http.Transport).RoundTrip(&req)
      if err != nil {
         panic(err)
      }
      body, err := io.ReadAll(res.Body)
      if err != nil {
         panic(err)
      }
      if err := res.Body.Close(); err != nil {
         panic(err)
      }
      if bytes.Contains(body, []byte(adaptive)) {
         fmt.Println("pass")
      } else {
         panic(adaptive)
      }
      time.Sleep(time.Second)
   }
}

const adaptive = `"adaptiveFormats"`

const req_body = `
{
   "contentCheckOk": true,
   "context": {
      "client": {
         "clientName": "ANDROID",
         "clientVersion": "18.14.40"
      }
   },
   "params": "8AEB",
   "videoId": "K8YwZkdCcGw"
}
`
