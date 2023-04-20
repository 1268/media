package main

import (
   "2a.pages.dev/rosso/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "time"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Method = "POST"
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   req_body := protobuf.Message{
      2: protobuf.String("Xxk-ryO6J2I"),
      1: protobuf.Message{
         1: protobuf.Message{
            16: protobuf.Varint(3),
            17: protobuf.String("18.14.40"),
            64: protobuf.Varint(26),
         },
      },
   }.Marshal()
   req.Header["User-Agent"] = []string{"com.google.android.youtube/18.14.40"}
   for range [16]struct{}{} {
      req.Body = io.NopCloser(bytes.NewReader(req_body))
      res, err := new(http.Transport).RoundTrip(&req)
      if err != nil {
         panic(err)
      }
      res_body, err := io.ReadAll(res.Body)
      if err != nil {
         panic(err)
      }
      if err := res.Body.Close(); err != nil {
         panic(err)
      }
      res_mes, err := protobuf.Unmarshal(res_body)
      if err != nil {
         panic(err)
      }
      view_count, err := res_mes.Get(11).Get_String(32)
      if err != nil {
         panic(err)
      }
      adaptive_formats := res_mes.Get(4).Get_Messages(3)
      fmt.Println(view_count, len(adaptive_formats))
      time.Sleep(time.Second)
   }
}

