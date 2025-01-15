package rakuten

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
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

func (g *gizmo_content) String() string {
   var b []byte
   for _, stream := range g.ViewOptions.Private.Streams {
      for _, language := range stream.AudioLanguages {
         b = append(b, "\naudio language = "...)
         b = append(b, language.Id...)
      }
   }
   b = append(b, "\nid = "...)
   b = append(b, g.Id...)
   b = append(b, "\nnumber = "...)
   b = strconv.AppendInt(b, int64(g.Number), 10)
   b = append(b, "\nseason number = "...)
   b = strconv.AppendInt(b, int64(g.SeasonNumber), 10)
   b = append(b, "\ntitle = "...)
   b = append(b, g.Title...)
   b = append(b, "\ntv show = "...)
   b = append(b, g.TvShowTitle...)
   b = append(b, "\ntype = "...)
   b = append(b, g.Type...)
   b = append(b, "\nyear = "...)
   b = strconv.AppendInt(b, int64(g.Year), 10)
   return string(b)
}
