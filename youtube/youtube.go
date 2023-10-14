package youtube

import (
   "net/url"
   "strings"
   "time"
)

func (p Player) Title() string {
   return p.Video_Details.Title
}

func (p Player) Author() string {
   return p.Video_Details.Author
}

type Player struct {
   Video_Details struct {
      Author string
      Length_Seconds int64 `json:"lengthSeconds,string"`
      Short_Description string `json:"shortDescription"`
      Title string
      Video_ID string `json:"videoId"`
      View_Count int64 `json:"viewCount,string"`
   } `json:"videoDetails"`
   Microformat struct {
      Player_Microformat_Renderer struct {
         Publish_Date string `json:"publishDate"`
      } `json:"playerMicroformatRenderer"`
   }
   Playability_Status struct {
      Status string
      Reason string
   } `json:"playabilityStatus"`
   Streaming_Data struct {
      Adaptive_Formats []Format `json:"adaptiveFormats"`
   } `json:"streamingData"`
}

func (i Image) URL(id string) *url.URL {
   return &url.URL{
      Scheme: "http",
      Host: "i.ytimg.com",
      Path: func() string {
         var b strings.Builder
         b.WriteString("/vi")
         if strings.HasSuffix(i.Name, ".webp") {
            b.WriteString("_webp")
         }
         b.WriteByte('/')
         b.WriteString(id)
         b.WriteByte('/')
         b.WriteString(i.Name)
         return b.String()
      }(),
   }
}

type Image struct {
   Crop bool
   Height int
   Name string
   Width int
}

var Images = []Image{
   {Width:120, Height:90, Name:"default.jpg"},
   {Width:120, Height:90, Name:"1.jpg"},
   {Width:120, Height:90, Name:"2.jpg"},
   {Width:120, Height:90, Name:"3.jpg"},
   {Width:120, Height:90, Name:"default.webp"},
   {Width:120, Height:90, Name:"1.webp"},
   {Width:120, Height:90, Name:"2.webp"},
   {Width:120, Height:90, Name:"3.webp"},
   {Width:320, Height:180, Name:"mq1.jpg", Crop:true},
   {Width:320, Height:180, Name:"mq2.jpg", Crop:true},
   {Width:320, Height:180, Name:"mq3.jpg", Crop:true},
   {Width:320, Height:180, Name:"mqdefault.jpg"},
   {Width:320, Height:180, Name:"mq1.webp", Crop:true},
   {Width:320, Height:180, Name:"mq2.webp", Crop:true},
   {Width:320, Height:180, Name:"mq3.webp", Crop:true},
   {Width:320, Height:180, Name:"mqdefault.webp"},
   {Width:480, Height:360, Name:"0.jpg"},
   {Width:480, Height:360, Name:"hqdefault.jpg"},
   {Width:480, Height:360, Name:"hq1.jpg"},
   {Width:480, Height:360, Name:"hq2.jpg"},
   {Width:480, Height:360, Name:"hq3.jpg"},
   {Width:480, Height:360, Name:"0.webp"},
   {Width:480, Height:360, Name:"hqdefault.webp"},
   {Width:480, Height:360, Name:"hq1.webp"},
   {Width:480, Height:360, Name:"hq2.webp"},
   {Width:480, Height:360, Name:"hq3.webp"},
   {Width:640, Height:480, Name:"sddefault.jpg"},
   {Width:640, Height:480, Name:"sd1.jpg"},
   {Width:640, Height:480, Name:"sd2.jpg"},
   {Width:640, Height:480, Name:"sd3.jpg"},
   {Width:640, Height:480, Name:"sddefault.webp"},
   {Width:640, Height:480, Name:"sd1.webp"},
   {Width:640, Height:480, Name:"sd2.webp"},
   {Width:640, Height:480, Name:"sd3.webp"},
   {Width:1280, Height:720, Name:"hq720.jpg"},
   {Width:1280, Height:720, Name:"maxresdefault.jpg"},
   {Width:1280, Height:720, Name:"maxres1.jpg"},
   {Width:1280, Height:720, Name:"maxres2.jpg"},
   {Width:1280, Height:720, Name:"maxres3.jpg"},
   {Width:1280, Height:720, Name:"hq720.webp"},
   {Width:1280, Height:720, Name:"maxresdefault.webp"},
   {Width:1280, Height:720, Name:"maxres1.webp"},
   {Width:1280, Height:720, Name:"maxres2.webp"},
   {Width:1280, Height:720, Name:"maxres3.webp"},
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
