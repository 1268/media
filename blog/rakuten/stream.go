package rakuten

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

// geo block
func (a *address) streamings(
   content *gizmo_content, language string,
) ([]stream_info, error) {
   classification, err := a.classification_id()
   if err != nil {
      return nil, err
   }
   data, err := json.Marshal(map[string]string{
      "audio_language":              language,
      "audio_quality":               "2.0",
      "classification_id":           classification,
      "content_id":                  content.Id,
      "content_type":                content.Type,
      "device_identifier":           "atvui40",
      "device_serial":               "not implemented",
      "device_stream_video_quality": "FHD",
      "player":                      "atvui40:DASH-CENC:WVM",
      "subtitle_language":           "MIS",
      "video_type":                  "stream",
   })
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
