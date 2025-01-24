package kanopy

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "path"
   "strconv"
)

type wrapper struct {
   video_manifest *video_manifest
   web_token      *web_token
}

func (v video_plays) dash() (*video_manifest, bool) {
   for _, manifest := range v.Manifests {
      if manifest.ManifestType == "dash" {
         return &manifest, true
      }
   }
   return nil, false
}

func (w wrapper) Wrap(data []byte) ([]byte, error) {
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/kapi/licenses/widevine/" + w.video_manifest.DrmLicenseId
   req.Header = http.Header{
      "authorization": {"Bearer " + w.web_token.Jwt},
      "user-agent":    {user_agent},
      "x-version":     {x_version},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

const (
   user_agent = "!"
   x_version = "!/!/!/!"
)

type video_manifest struct {
   DrmLicenseId string
   ManifestType string
   Url          string
}

type video_plays struct {
   Manifests []video_manifest
}

// good for 10 years
type web_token struct {
   Jwt    string
   UserId int
}

func (web_token) marshal(email, password string) ([]byte, error) {
   data, err := json.Marshal(map[string]any{
      "credentialType": "email",
      "emailUser": map[string]string{
         "email":    email,
         "password": password,
      },
   })
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
      "user-agent":   {user_agent},
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

type membership struct {
   DomainId int
}

func (w *web_token) membership() (*membership, error) {
   req, _ := http.NewRequest("", "https://www.kanopy.com", nil)
   req.URL.Path = "/kapi/memberships"
   req.URL.RawQuery = "userId=" + strconv.Itoa(w.UserId)
   req.Header = http.Header{
      "authorization": {"Bearer " + w.Jwt},
      "user-agent":    {user_agent},
      "x-version":     {x_version},
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

func (a *address) Set(data string) error {
   var err error
   a.video_id, err = strconv.Atoi(path.Base(data))
   if err != nil {
      return err
   }
   return nil
}

func (a address) String() string {
   return strconv.Itoa(a.video_id)
}

type address struct {
   video_id int
}

func (w *web_token) plays(
   member *membership, web address,
) (*video_plays, error) {
   data, err := json.Marshal(map[string]int{
      "domainId": member.DomainId,
      "userId":   w.UserId,
      "videoId":  web.video_id,
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
      "content-type":  {"application/json"},
      "user-agent":    {user_agent},
      "x-version":     {x_version},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   play := &video_plays{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}
