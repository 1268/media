package kanopy

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strconv"
)

func (u Url) Mpd() (*http.Response, error) {
   req, err := http.NewRequest("", string(u), nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("user-agent", "Mozilla")
   return http.DefaultClient.Do(req)
}

type Url string

type VideoManifest struct {
   DrmLicenseId string
   ManifestType string
   Url          Url
}

func (w *WebToken) Plays(
   member *Membership, video_id int,
) (*VideoPlays, error) {
   data, err := json.Marshal(map[string]int{
      "domainId": member.DomainId,
      "userId":   w.UserId,
      "videoId":  video_id,
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
   play := &VideoPlays{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

type VideoPlays struct {
   ErrorMsgLong string `json:"error_msg_long"`
   Manifests []VideoManifest
}

func (v VideoPlays) Dash() (*VideoManifest, bool) {
   for _, manifest := range v.Manifests {
      if manifest.ManifestType == "dash" {
         return &manifest, true
      }
   }
   return nil, false
}

type Wrapper struct {
   Manifest *VideoManifest
   Token    *WebToken
}

func (w Wrapper) Wrap(data []byte) ([]byte, error) {
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/kapi/licenses/widevine/" + w.Manifest.DrmLicenseId
   req.Header = http.Header{
      "authorization": {"Bearer " + w.Token.Jwt},
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
   x_version  = "!/!/!/!"
)

// good for 10 years
type WebToken struct {
   Jwt    string
   UserId int
}

func (WebToken) Marshal(email, password string) ([]byte, error) {
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

func (w *WebToken) Unmarshal(data []byte) error {
   return json.Unmarshal(data, w)
}

type Membership struct {
   DomainId int
}

func (w *WebToken) Membership() (*Membership, error) {
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
   var value struct {
      List []Membership
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.List[0], nil
}
