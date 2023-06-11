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

// YouTube on TV
const (
   client_ID =
      "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68" +
      ".apps.googleusercontent.com"
   client_secret = "SboVhoG9s0rNafixCSGGKXAT"
)

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
