package tubi

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (v *VideoContent) Series() bool {
   return v.DetailedType == "series"
}

func (v *VideoContent) Episode() bool {
   return v.DetailedType == "episode"
}

func (v *VideoContent) Video() (*VideoResource, bool) {
   if len(v.VideoResources) == 0 {
      return nil, false
   }
   a := v.VideoResources[0]
   for _, b := range v.VideoResources {
      if b.Resolution.Data > a.Resolution.Data {
         a = b
      }
   }
   return &a, true
}

func (v *VideoContent) Get(id int) (*VideoContent, bool) {
   if v.Id == id {
      return v, true
   }
   for _, child := range v.Children {
      if content, ok := child.Get(id); ok {
         return content, true
      }
   }
   return nil, false
}

func (v *VideoContent) set(parent *VideoContent) {
   v.parent = parent
   for _, child := range v.Children {
      child.set(v)
   }
}

func (v *VideoContent) Unmarshal(data []byte) error {
   err := json.Unmarshal(data, v)
   if err != nil {
      return err
   }
   v.set(nil)
   return nil
}

type VideoContent struct {
   Children       []*VideoContent
   DetailedType   string `json:"detailed_type"`
   EpisodeNumber  int    `json:"episode_number,string"`
   Id             int    `json:",string"`
   SeriesId       int    `json:"series_id,string"`
   Title          string
   VideoResources []VideoResource `json:"video_resources"`
   Year           int
   parent         *VideoContent
}

func (*VideoContent) Marshal(id int) ([]byte, error) {
   req, err := http.NewRequest("", "https://uapi.adrise.tv/cms/content", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "content_id": {strconv.Itoa(id)},
      "deviceId":   {"!"},
      "platform":   {"android"},
      "video_resources[]": {
         "dash",
         "dash_widevine",
      },
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}
func (v *VideoResource) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      v.LicenseServer.Url, "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type VideoResource struct {
   LicenseServer *struct {
      Url string
   } `json:"license_server"`
   Manifest struct {
      Url string
   }
   Resolution Resolution
   Type       string
}

type Resolution struct {
   Data int64
}

func (r *Resolution) UnmarshalText(text []byte) error {
   s := string(text)
   s = strings.TrimPrefix(s, "VIDEO_RESOLUTION_")
   s = strings.TrimSuffix(s, "P")
   var err error
   r.Data, err = strconv.ParseInt(s, 10, 64)
   if err != nil {
      return err
   }
   return nil
}

func (r Resolution) MarshalText() ([]byte, error) {
   b := []byte("VIDEO_RESOLUTION_")
   b = strconv.AppendInt(b, r.Data, 10)
   return append(b, 'P'), nil
}

func (n Namer) Show() string {
   if v := n.Content.parent; v != nil {
      return v.parent.Title
   }
   return ""
}

// S01:E03 - Hell Hath No Fury
func (n Namer) Title() string {
   if _, v, ok := strings.Cut(n.Content.Title, " - "); ok {
      return v
   }
   return n.Content.Title
}

type Namer struct {
   Content *VideoContent
}

func (n Namer) Episode() int {
   return n.Content.EpisodeNumber
}

func (n Namer) Year() int {
   return n.Content.Year
}

func (n Namer) Season() int {
   if n.Content.parent != nil {
      return n.Content.parent.Id
   }
   return 0
}
