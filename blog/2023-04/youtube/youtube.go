package main

import (
   "2a.pages.dev/mech/youtube"
   "encoding/json"
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
   req.Header["Host"] = []string{"www.youtube.com"}
   req.Header["X-Goog-Api-Key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   for _, version := range client_versions {
      req_body := fmt.Sprintf(`{
       "contentCheckOk": true,
       "context": {
        "client": {
         "clientName": "ANDROID",
         "clientVersion": %q
        }
       },
       "videoId": "PCeQAvlCULo"
      }
      `, version)
      req.Body = io.NopCloser(strings.NewReader(req_body))
      res, err := new(http.Transport).RoundTrip(&req)
      if err != nil {
         panic(err)
      }
      var p youtube.Player
      if err := json.NewDecoder(res.Body).Decode(&p); err != nil {
         panic(err)
      }
      if err := res.Body.Close(); err != nil {
         panic(err)
      }
      fmt.Println(check(p), version)
      time.Sleep(99 * time.Millisecond)
   }
}

func check(p youtube.Player) string {
   if p.Playability_Status.Status == "" {
      return "status"
   }
   if p.Video_Details.Author == "" {
      return "author"
   }
   if p.Video_Details.Length_Seconds == 0 {
      return "length seconds"
   }
   if p.Video_Details.Short_Description == "" {
      return "description"
   }
   if p.Video_Details.Title == "" {
      return "title"
   }
   if p.Video_Details.Video_ID == "" {
      return "video ID"
   }
   if p.Video_Details.View_Count == 0 {
      return "view count"
   }
   if len(p.Streaming_Data.Adaptive_Formats) == 0 {
      return "adaptive formats"
   }
   if p.Streaming_Data.Adaptive_Formats[0].Bitrate == 0 {
      return "bitrate"
   }
   if p.Streaming_Data.Adaptive_Formats[0].Content_Length == 0 {
      return "content length"
   }
   if p.Streaming_Data.Adaptive_Formats[0].Height == 0 {
      return "height"
   }
   if p.Streaming_Data.Adaptive_Formats[0].MIME_Type == "" {
      return "MIME type"
   }
   if p.Streaming_Data.Adaptive_Formats[0].Quality_Label == "" {
      return "quality label"
   }
   if p.Streaming_Data.Adaptive_Formats[0].URL == "" {
      return "URL"
   }
   if p.Streaming_Data.Adaptive_Formats[0].Width == 0 {
      return "width"
   }
   return ""
}
