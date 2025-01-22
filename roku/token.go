package roku

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
)

func (p *Playback) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      p.Drm.Widevine.LicenseServer, "application/x-protobuf",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type Playback struct {
   Drm struct {
      Widevine struct {
         LicenseServer string
      }
   }
   Url string
}

func (AccountToken) Marshal(
   auth *AccountAuth, code *AccountCode,
) ([]byte, error) {
   req, err := http.NewRequest("", "https://googletv.web.roku.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/v1/account/activation/" + code.Code
   req.Header = http.Header{
      "user-agent":           {user_agent},
      "x-roku-content-token": {auth.AuthToken},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

const user_agent = "trc-googletv; production; 0"

type AccountToken struct {
   Token string
}

func (a *AccountToken) Unmarshal(data []byte) error {
   return json.Unmarshal(data, a)
}
