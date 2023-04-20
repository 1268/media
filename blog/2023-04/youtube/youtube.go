package main

import (
   "io"
   "net/url"
   "strings"
   "time"
   "net/http"
)

func main() {
   req := new(http.Request)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.Header = make(http.Header)
   req.Header["User-Agent"] = []string{"com.google.android.youtube/18.14.40"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   req_body := `
   {
      "context": {
         "client": {
            "androidSdkVersion": 26,
            "clientName": "ANDROID",
            "clientVersion": "18.14.40"
         }
      },
      "videoId": "Xxk-ryO6J2I"
   }
   `
   for range [16]struct{}{} {
      req.Body = io.NopCloser(strings.NewReader(req_body))
      res, err := new(http.Transport).RoundTrip(req)
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
      text := string(body)
      println(
         "adaptiveFormats", strings.Contains(text, `"adaptiveFormats"`),
         "viewCount", strings.Contains(text, `"viewCount"`),
      )
      time.Sleep(time.Second)
   }
}
