package rakuten

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

// geo block
func (o on_demand) streamings() ([]stream_info, error) {
   o.AudioQuality = "2.0"
   o.ContentType = "movies"
   o.DeviceIdentifier = "atvui40"
   o.DeviceStreamVideoQuality = "FHD"
   o.Player = "atvui40:DASH-CENC:WVM"
   o.SubtitleLanguage = "MIS"
   o.VideoType = "stream"
   o.DeviceSerial = "not implemented"
   // cz = fail
   // fr = pass
   o.AudioLanguage = "ENG"
   // cz = pass
   // fr = fail
   //o.AudioLanguage = "SPA"
   data, err := json.Marshal(o)
   if err != nil {
      return nil, err
   }
   resp, err := http.Post(
      "https://gizmo.rakuten.tv/v3/avod/streamings",
      "application/json", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   var value struct {
      Data struct {
         StreamInfos []stream_info `json:"stream_infos"`
      }
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   if err := value.Errors; len(err) >= 1 {
      return nil, errors.New(err[0].Message)
   }
   return value.Data.StreamInfos, nil
}

type stream_info struct {
   LicenseUrl   string `json:"license_url"`
   Url          string
   VideoQuality string `json:"video_quality"`
}

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
