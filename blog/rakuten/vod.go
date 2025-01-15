package rakuten

func (g *gizmo_content) video(
   classification_id int, language, quality string,
) on_demand {
   return on_demand{
      AudioLanguage: language,
      AudioQuality: "2.0",
      ClassificationId: classification_id,
      ContentId: g.Id,
      ContentType: g.Type,
      DeviceIdentifier: "atvui40",
      DeviceSerial: "not implemented",
      DeviceStreamVideoQuality: quality,
      Player: "atvui40:DASH-CENC:WVM",
      SubtitleLanguage: "MIS",
      VideoType: "stream",
   }
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
