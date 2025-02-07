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

func (VideoContent) Marshal(id int) ([]byte, error) {
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

func (r *Resolution) UnmarshalText(data []byte) error {
   var err error
   data1 := strings.TrimPrefix(string(data), "VIDEO_RESOLUTION_")
   (*r)[0], err = strconv.ParseInt(strings.TrimSuffix(data1, "P"), 10, 64)
   if err != nil {
      return err
   }
   return nil
}

func (r Resolution) MarshalText() ([]byte, error) {
   data := []byte("VIDEO_RESOLUTION_")
   data = strconv.AppendInt(data, r[0], 10)
   return append(data, 'P'), nil
}

type VideoContent struct {
   Children       []*VideoContent
   DetailedType   string `json:"detailed_type"`
   Id             int    `json:",string"`
   VideoResources []VideoResource `json:"video_resources"`
   parent         *VideoContent
}

type Resolution [1]int64

func (v *VideoContent) Resource() (*VideoResource, bool) {
   if len(v.VideoResources) == 0 {
      return nil, false
   }
   a := v.VideoResources[0]
   for _, b := range v.VideoResources {
      if b.Resolution[0] > a.Resolution[0] {
         a = b
      }
   }
   return &a, true
}
