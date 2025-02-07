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

type Content struct {
   Children       []*Content
   DetailedType   string `json:"detailed_type"`
   Id             int    `json:",string"`
   SeriesId       int    `json:"series_id,string"`
   VideoResources []VideoResource `json:"video_resources"`
   parent         *Content
}

type Resolution [1]int64

func (r *Resolution) UnmarshalText(data []byte) error {
   var err error
   data1 := strings.TrimPrefix(string(data), "VIDEO_RESOLUTION_")
   (*r)[0], err = strconv.ParseInt(strings.TrimSuffix(data1, "P"), 10, 64)
   if err != nil {
      return err
   }
   return nil
}

func (c *Content) Series() bool {
   return c.DetailedType == "series"
}

func (c *Content) Episode() bool {
   return c.DetailedType == "episode"
}

func (r Resolution) MarshalText() ([]byte, error) {
   data := []byte("VIDEO_RESOLUTION_")
   data = strconv.AppendInt(data, r[0], 10)
   return append(data, 'P'), nil
}

func (c *Content) set(parent *Content) {
   c.parent = parent
   for _, child := range c.Children {
      child.set(c)
   }
}

func (c *Content) Unmarshal(data []byte) error {
   err := json.Unmarshal(data, c)
   if err != nil {
      return err
   }
   c.set(nil)
   return nil
}

func (c *Content) Resource() (*VideoResource, bool) {
   if len(c.VideoResources) == 0 {
      return nil, false
   }
   a := c.VideoResources[0]
   for _, b := range c.VideoResources {
      if b.Resolution[0] > a.Resolution[0] {
         a = b
      }
   }
   return &a, true
}

func (Content) Marshal(id int) ([]byte, error) {
   req, _ := http.NewRequest("", "https://uapi.adrise.tv/cms/content", nil)
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

func (c *Content) Get(id int) (*Content, bool) {
   if c.Id == id {
      return c, true
   }
   for _, child := range c.Children {
      content0, ok := child.Get(id)
      if ok {
         return content0, true
      }
   }
   return nil, false
}

func (v *VideoResource) License(data []byte) (*http.Response, error) {
   return http.Post(
      v.LicenseServer.Url, "application/x-protobuf", bytes.NewReader(data),
   )
}

func (v *VideoResource) Mpd() (*http.Response, error) {
   return http.Get(v.Manifest.Url)
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
