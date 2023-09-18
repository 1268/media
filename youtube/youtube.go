package youtube

import (
   "154.pages.dev/media"
   "net/url"
   "strings"
   "time"
)

func (p Player) Name() string {
   var b strings.Builder
   b.WriteString(p.Video_Details.Author)
   b.WriteString(sep_big)
   b.WriteString(media.Clean(p.Video_Details.Title))
   return b.String()
}

func (i Image) URL(id string) *url.URL {
   var u url.URL
   u.Scheme = "http"
   u.Host = "i.ytimg.com"
   {
      var b strings.Builder
      b.WriteString("/vi")
      if strings.HasSuffix(i.Name, ".webp") {
         b.WriteString("_webp")
      }
      b.WriteByte('/')
      b.WriteString(id)
      b.WriteByte('/')
      b.WriteString(i.Name)
      u.Path = b.String()
   }
   return &u
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

const sep_big = " - "

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
      Adaptive_Formats []Format `json:"adaptiveFormats"`
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

func (s Search) Items() []Item {
   var items []Item
   for _, sect := range s.Contents.Section_List_Renderer.Contents {
      if sect.Item_Section_Renderer != nil {
         for _, item := range sect.Item_Section_Renderer.Contents {
            if item.Video_With_Context_Renderer != nil {
               items = append(items, item)
            }
         }
      }
   }
   return items
}

type Search struct {
   Contents struct {
      Section_List_Renderer struct {
         Contents []struct {
            Item_Section_Renderer *struct {
               Contents []Item
            } `json:"itemSectionRenderer"`
         }
      } `json:"sectionListRenderer"`
   }
}

type Item struct {
   Video_With_Context_Renderer *struct {
      Video_ID string `json:"videoId"`
      Headline struct {
         Runs []struct {
            Text string
         }
      }
   } `json:"videoWithContextRenderer"`
}
