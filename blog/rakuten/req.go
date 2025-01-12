package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.URL = &url.URL{}
   req.URL.Host = "gizmo.rakuten.tv"
   req.URL.Path = "/v3/avod/streamings"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/json"}
   value := url.Values{}
   value["device_identifier"] = []string{"web"}
   req.URL.RawQuery = value.Encode()
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

var body = strings.NewReader(`
{
   "audio_language": "SPA",
   "audio_quality": "2.0",
   "classification_id": "272",
   "content_id": "transvulcania-the-people-s-run",
   "content_type": "movies",
   "device_serial": "not implemented",
   "device_stream_video_quality": "HD",
   "player": "web:DASH-CENC:WVM",
   "video_type": "stream",
   "subtitle_language": "MIS"
}
`)
