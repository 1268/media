package kanopy

import (
   "encoding/json"
   "net/http"
   "strconv"
)

type video_items struct {
   List []struct {
      Video struct {
         VideoId int
      }
   }
}

func (w *web_token) items(video_id int64) (*video_items, error) {
   req, err := http.NewRequest("", "https://www.kanopy.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/kapi/videos/")
      b = strconv.AppendInt(b, video_id, 10)
      b = append(b, "/items"...)
      return string(b)
   }()
   req.Header = http.Header{
      "authorization": {"Bearer " + w.Jwt},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   items := &video_items{}
   err = json.NewDecoder(resp.Body).Decode(items)
   if err != nil {
      return nil, err
   }
   return items, nil
}
