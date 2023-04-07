package youtube

import (
   "2a.pages.dev/mech"
   "strconv"
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

func (p Status) String() string {
   var b strings.Builder
   b.WriteString("status: ")
   b.WriteString(p.Status)
   if p.Reason != "" {
      b.WriteString("\nreason: ")
      b.WriteString(p.Reason)
   }
   return b.String()
}

func (p Player) String() string {
   var b []byte
   b = append(b, p.Playability_Status.String()...)
   b = append(b, "\nduration: "...)
   b = append(b, p.Duration().String()...)
   if p.Publish_Date() != "" {
      b = append(b, "\npublish date: "...)
      b = append(b, p.Publish_Date()...)
   }
   b = append(b, "\nauthor: "...)
   b = append(b, p.Video_Details.Author...)
   b = append(b, "\ntitle: "...)
   b = append(b, p.Video_Details.Title...)
   b = append(b, "\nvideo ID: "...)
   b = append(b, p.Video_Details.Video_ID...)
   b = append(b, "\nview count: "...)
   b = strconv.AppendInt(b, p.Video_Details.View_Count, 10)
   for _, form := range p.Streaming_Data.Adaptive_Formats {
      b = append(b, '\n')
      b = append(b, form.String()...)
   }
   return string(b)
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

type Status struct {
   Status string
   Reason string
}

type Player struct {
   Microformat struct {
      Player_Microformat_Renderer struct {
         Publish_Date string `json:"publishDate"`
      } `json:"playerMicroformatRenderer"`
   }
   Playability_Status Status `json:"playabilityStatus"`
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
