package amc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "strings"
)

const sep_big = " - "

func (v Video_Player) Name() (string, error) {
   year, _, found := strings.Cut(v.Current_Video.Meta.Airdate, "-")
   if !found {
      return "", errors.New("year not found")
   }
   var b strings.Builder
   b.WriteString(v.Current_Video.Meta.Title)
   if v.Content_Type == "movie" {
      b.WriteString(sep_big)
      b.WriteString(year)
   }
   return b.String(), nil
}

type Video_Player struct {
   Content_Type string `json:"contentType"`
   Current_Video struct {
      Meta struct {
         Airdate string // 1996-01-01T00:00:00.000Z
         Title string
      }
   } `json:"currentVideo"`
}

type playback_request struct {
   Ad_Tags struct {
      Lat int `json:"lat"`
      Mode string `json:"mode"`
      PPID int `json:"ppid"`
      Player_Height int `json:"playerHeight"`
      Player_Width int `json:"playerWidth"`
      URL string `json:"url"`
   } `json:"adtags"`
}

func (Playback) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (p Playback) Request_Header() http.Header {
   token := p.head.Get("X-AMCN-BC-JWT")
   head := make(http.Header)
   head.Set("bcov-auth", token)
   return head
}

type Playback struct {
   head http.Header
   body struct {
      Data struct {
         Playback_JSON_Data struct {
            Sources []Source
         } `json:"playbackJsonData"`
      }
   }
}

func (p Playback) Request_URL() string {
   return p.Source().Key_Systems.Widevine.License_URL
}

type Source struct {
   Key_Systems *struct {
      Widevine struct {
         License_URL string
      } `json:"com.widevine.alpha"`
   }
   Src string
   Type string
}

func (p Playback) Source() *Source {
   for _, item := range p.body.Data.Playback_JSON_Data.Sources {
      if item.Type == "application/dash+xml" {
         return &item
      }
   }
   return nil
}

type Content struct {
   Data	struct {
      Children []struct {
         Properties json.RawMessage
         Type string
      }
   }
}

func (c Content) Video_Player() (*Video_Player, error) {
   for _, child := range c.Data.Children {
      if child.Type == "video-player-ap" {
         vid := new(Video_Player)
         err := json.Unmarshal(child.Properties, vid)
         if err != nil {
            return nil, err
         }
         return vid, nil
      }
   }
   return nil, errors.New("video-player-ap not present")
}
