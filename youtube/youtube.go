package youtube

import (
   "net/url"
   "path"
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

// YouTube on TV
const (
   client_ID =
      "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68" +
      ".apps.googleusercontent.com"
   client_secret = "SboVhoG9s0rNafixCSGGKXAT"
)

func Video_ID(data string, v *string) error {
   ref, err := url.Parse(data)
   if err != nil {
      return err
   }
   *v = ref.Query().Get("v")
   if *v == "" {
      *v = path.Base(ref.Path)
   }
   return nil
}
