package tubi

import (
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strconv"
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
      if b.Resolution.Int64 > a.Resolution.Int64 {
         a = b
      }
   }
   return &a, true
}

func (v *VideoContent) Unmarshal() error {
   err := json.Unmarshal(v.Raw, v)
   if err != nil {
      return err
   }
   v.set(nil)
   return nil
}

func (v *VideoContent) New(id int) error {
   req, err := http.NewRequest("", "https://uapi.adrise.tv/cms/content", nil)
   if err != nil {
      return err
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
      return err
   }
   defer resp.Body.Close()
   v.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
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

type VideoContent struct {
   Children       []*VideoContent
   DetailedType   string `json:"detailed_type"`
   EpisodeNumber  int    `json:"episode_number,string"`
   Id             int    `json:",string"`
   Raw            []byte `json:"-"`
   SeriesId       int    `json:"series_id,string"`
   Title          string
   VideoResources []VideoResource `json:"video_resources"`
   Year           int
   parent         *VideoContent
}

func (v *VideoContent) set(parent *VideoContent) {
   v.parent = parent
   for _, child := range v.Children {
      child.set(v)
   }
}
