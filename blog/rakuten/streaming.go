package rakuten

import (
   "io"
   "net/http"
   "net/url"
   "strings"
)

type on_demand struct {
   AudioLanguage            string `json:"audio_language"`
   AudioQuality             string `json:"audio_quality"`
   ClassificationId         int    `json:"classification_id"`
   ContentId                string `json:"content_id"`
   ContentType              string `json:"content_type"`
   DeviceIdentifier         string `json:"device_identifier"`
   DeviceSerial             string `json:"device_serial"`
   DeviceStreamVideoQuality string `json:"device_stream_video_quality"`
   Player                   string `json:"player"`
   SubtitleLanguage         string `json:"subtitle_language"`
   VideoType                string `json:"video_type"`
}

func (o *on_demand) streamings() (*http.Response, error) {
   o.AudioLanguage = "SPA"
   o.AudioQuality = "2.0"
   o.ClassificationId = 272
   o.ContentId = "transvulcania-the-people-s-run"
   o.ContentType = "movies"
   o.DeviceIdentifier = "atvui40"
   o.DeviceSerial = "not implemented"
   o.DeviceStreamVideoQuality = "FHD"
   o.Player = "atvui40:DASH-CENC:WVM"
   o.SubtitleLanguage = "MIS"
   o.VideoType = "stream"
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.URL = &url.URL{}
   req.URL.Host = "gizmo.rakuten.tv"
   req.URL.Path = "/v3/avod/streamings"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/json"}
   return http.DefaultClient.Do(&req)
}
