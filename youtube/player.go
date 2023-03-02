package youtube

import (
   "strconv"
   "strings"
   "time"
)

func (p Player) Time() (time.Time, error) {
   return time.Parse(time.DateOnly, p.Publish_Date())
}

func (p Player) MarshalText() ([]byte, error) {
   b := []byte(p.Playability_Status.String())
   b = append(b, "\nVideo ID: "...)
   b = append(b, p.Video_Details.Video_ID...)
   b = append(b, "\nDuration: "...)
   b = append(b, p.Duration().String()...)
   b = append(b, "\nView Count: "...)
   b = strconv.AppendInt(b, p.Video_Details.View_Count, 10)
   b = append(b, "\nAuthor: "...)
   b = append(b, p.Video_Details.Author...)
   b = append(b, "\nTitle: "...)
   b = append(b, p.Video_Details.Title...)
   if p.Publish_Date() != "" {
      b = append(b, "\nPublish Date: "...)
      b = append(b, p.Publish_Date()...)
   }
   b = append(b, '\n')
   for _, form := range p.Streaming_Data.Adaptive_Formats {
      t, err := form.MarshalText()
      if err != nil {
         return nil, err
      }
      b = append(b, t...)
   }
   return b, nil
}

func (p Player) Duration() time.Duration {
   return time.Duration(p.Video_Details.Length_Seconds) * time.Second
}

func (p Player) Publish_Date() string {
   return p.Microformat.Player_Microformat_Renderer.Publish_Date
}

func (p Player) Name() string {
   var buf strings.Builder
   buf.WriteString(p.Video_Details.Author)
   buf.WriteByte('-')
   buf.WriteString(p.Video_Details.Title)
   return buf.String()
}

func (p Status) String() string {
   var buf strings.Builder
   buf.WriteString("Status: ")
   buf.WriteString(p.Status)
   if p.Reason != "" {
      buf.WriteString("\nReason: ")
      buf.WriteString(p.Reason)
   }
   return buf.String()
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
   Streaming_Data struct {
      Adaptive_Formats Formats `json:"adaptiveFormats"`
   } `json:"streamingData"`
   Video_Details struct {
      Author string
      Length_Seconds int64 `json:"lengthSeconds,string"`
      Short_Description string `json:"shortDescription"`
      Title string
      Video_ID string `json:"videoId"`
      View_Count int64 `json:"viewCount,string"`
   } `json:"videoDetails"`
}
