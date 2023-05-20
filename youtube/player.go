package youtube

import (
   "2a.pages.dev/mech"
   "strings"
   "time"
)

const sep_big = " - "

func (p Player) Name() string {
   var b strings.Builder
   b.WriteString(p.Video_Details.Author)
   b.WriteString(sep_big)
   b.WriteString(mech.Clean(p.Video_Details.Title))
   return b.String()
}

type Player struct {
   Microformat struct {
      Player_Microformat_Renderer struct {
         Publish_Date string `json:"publishDate"`
      } `json:"playerMicroformatRenderer"`
   }
   Playability_Status struct {
      Status string
      Reason string
   } `json:"playabilityStatus"`
   Video_Details struct {
      Author string
      Length_Seconds int64 `json:"lengthSeconds,string"`
      Short_Description string `json:"shortDescription"`
      Title string
      Video_ID string `json:"videoId"`
      View_Count int64 `json:"viewCount,string"`
   } `json:"videoDetails"`
   Streaming_Data struct {
      Adaptive_Formats Formats `json:"adaptiveFormats"`
   } `json:"streamingData"`
}

func (p Player) Time() (time.Time, error) {
   return time.Parse(time.DateOnly, p.Publish_Date())
}

func (p Player) Duration() time.Duration {
   return time.Duration(p.Video_Details.Length_Seconds) * time.Second
}

func (p Player) Publish_Date() string {
   return p.Microformat.Player_Microformat_Renderer.Publish_Date
}
