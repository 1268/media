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
   val := make(url.Values)
   val["key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
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
      fmt.Println(
         "adaptive_formats",
         bytes.Contains(body, []byte(`"adaptiveFormats"`)),
         "view_count",
         bytes.Contains(body, []byte(`"viewCount"`)),
      )
      time.Sleep(time.Second)
   }
}

const req_body = `
{
   "context": {
      "client": {
         "hl": "en",
         "gl": "US",
         "remoteHost": "72.181.23.38",
         "deviceMake": "",
         "deviceModel": "",
         "visitorData": "CgtCYTlxQUNNUmJtOCi4qvGhBg%3D%3D",
         "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0,gzip(gfe)",
         "clientName": "ANDROID",
         "clientVersion": "18.14.40"
      }
   },
   "videoId": "coLCY15P6Bw",
   "contentCheckOk": false
}
`
