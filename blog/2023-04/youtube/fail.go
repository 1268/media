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
   req.Header["X-Goog-Visitor-Id"] = []string{"Cgtkb0NEOGQzRTNiayi5wvyhBg=="}
   req.Method = "POST"
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   req_body := protobuf.Message{
      2: protobuf.String("Xxk-ryO6J2I"), // videoId
      1:protobuf.Message{ // context
         1:protobuf.Message{ // client
            16: protobuf.Varint(3),
            17: protobuf.String("18.14.40"), // clientVersion
            18: protobuf.String("Android"), // clientName
         },
      },
   }.Marshal()
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
