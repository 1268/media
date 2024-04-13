package tubi

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
)

type content_management struct {
   Detailed_Type string
   Series_ID string
   Video_Resources []VideoResource
}

func (c *content_management) New(content_id int) error {
   req, err := http.NewRequest("GET", "https://uapi.adrise.tv/cms/content", nil)
   if err != nil {
      return err
   }
   req.URL.RawQuery = url.Values{
      "content_id": {strconv.Itoa(content_id)},
      "deviceId": {"!"},
      "platform": {"android"},
      "video_resources[]": {"dash_widevine"},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(c)
}

func (v VideoResource) RequestUrl() (string, bool) {
   return v.License_Server.URL, true
}

type VideoResource struct {
   License_Server struct {
      URL string
   }
   Manifest struct {
      URL string
   }
   Resolution string
}

func (VideoResource) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (VideoResource) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (VideoResource) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

func (c content_management) Resolution720p() (*VideoResource, bool) {
   for _, video := range c.Video_Resources {
      if video.Resolution == "VIDEO_RESOLUTION_720P" {
         return &video, true
      }
   }
   return nil, false
}
