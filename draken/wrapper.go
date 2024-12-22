package draken

import (
   "bytes"
   "io"
   "net/http"
)

type Wrapper struct {
   AuthLogin *AuthLogin
   Playback *Playback
}

func (w Wrapper) Wrap(data []byte) ([]byte, error) {
   req, err := http.NewRequest(
      "POST", "https://client-api.magine.com/api/playback/v1/widevine/license",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   magine_accesstoken.set(req.Header)
   for key, value := range w.Playback.Headers {
      req.Header.Set(key, value)
   }
   req.Header.Set("authorization", "Bearer " + w.AuthLogin.Token)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}
