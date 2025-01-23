package kanopy

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strconv"
)

func (v video_items) item(id int64) (int, *video_item) {
   for key, value := range v.List {
      if value.Video != nil {
         if value.Video.VideoId == id {
            return key, &value
         }
      } else {
         if value.Playlist.VideoId == id {
            return key, &value
         }
      }
   }
   return -1, nil
}

type video_items struct {
   List []video_item
}

type video_item struct {
   Playlist *struct {
      Title string
      VideoId int64
   }
   Video *struct {
      Title string
      VideoId int64
   }
}

func (v *video_item) String() string {
   var b []byte
   if v.Video != nil {
      b = append(b, "title = "...)
      b = append(b, v.Video.Title...)
      b = append(b, "\nvideo id = "...)
      b = strconv.AppendInt(b, v.Video.VideoId, 10)
   } else {
      b = append(b, "title = "...)
      b = append(b, v.Playlist.Title...)
      b = append(b, "\nvideo id = "...)
      b = strconv.AppendInt(b, v.Playlist.VideoId, 10)
   }
   return string(b)
}

func (w *web_token) items(id int64) (*video_items, error) {
   req, err := http.NewRequest("", "https://www.kanopy.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/kapi/videos/")
      b = strconv.AppendInt(b, id, 10)
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
type video_response struct {
   AncestorVideoIds []int64
   ProductionYear int
   Title string
   VideoId int64
}

func (w *web_token) video(id int64) (*video_response, error) {
   req, err := http.NewRequest("", "https://www.kanopy.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/kapi/videos/" + strconv.FormatInt(id, 10)
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
   var video struct {
      Video video_response
   }
   err = json.NewDecoder(resp.Body).Decode(&video)
   if err != nil {
      return nil, err
   }
   return &video.Video, nil
}

type poster struct {
   video_manifest *video_manifest
   web_token *web_token
}

func (p poster) Wrap(data []byte) ([]byte, error) {
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/kapi/licenses/widevine/" + p.video_manifest.DrmLicenseId
   req.Header = http.Header{
      "authorization": {"Bearer " + p.web_token.Jwt},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type video_manifest struct {
   DrmLicenseId string
   ManifestType string
   Url string
}

type video_plays struct {
   Manifests []video_manifest
}

func (w *web_token) plays(
   member *membership, video_id int64,
) (*video_plays, error) {
   data, err := json.Marshal(map[string]int64{
      "domainId": member.DomainId,
      "userId": w.UserId,
      "videoId": video_id,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com/kapi/plays", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "authorization": {"Bearer " + w.Jwt},
      "content-type": {"application/json"},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   play := &video_plays{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

func (w *web_token) membership() (*membership, error) {
   req, err := http.NewRequest("", "https://www.kanopy.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/kapi/memberships"
   req.URL.RawQuery = "userId=" + strconv.FormatInt(w.UserId, 10)
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
   var member struct {
      List []membership
   }
   err = json.NewDecoder(resp.Body).Decode(&member)
   if err != nil {
      return nil, err
   }
   return &member.List[0], nil
}

type membership struct {
   DomainId int64
}

func (v *video_plays) dash() (*video_manifest, bool) {
   for _, manifest := range v.Manifests {
      if manifest.ManifestType == "dash" {
         return &manifest, true
      }
   }
   return nil, false
}

const x_version = "!/!/!/!"

const user_agent = "!"

// good for 10 years
type web_token struct {
   Jwt string
   UserId int64
}

func (web_token) marshal(email, password string) ([]byte, error) {
   value := map[string]any{
      "credentialType": "email",
      "emailUser": map[string]string{
         "email": email,
         "password": password,
      },
   }
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com/kapi/login", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type": {"application/json"},
      "user-agent": {user_agent},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

func (w *web_token) unmarshal(data []byte) error {
   return json.Unmarshal(data, w)
}
