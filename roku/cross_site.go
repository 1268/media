package roku

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/json"
   "io"
)

func (c Cross_Site) Playback(id string) (*Playback, error) {
   body := map[string]string{
      "mediaFormat": "mpeg-dash",
      "providerId": "rokuavod",
      "rokuId": id,
   }
   raw, err := json.MarshalIndent(body, "", " ")
   if err != nil {
      return nil, err
   }
   req := http.Post()
   req.Body_Bytes(raw)
   // we could use Request.AddCookie, but we would need to call it after this,
   // otherwise it would be clobbered
   req.Header = http.Header{
      "CSRF-Token": {c.token},
      "Content-Type": {"application/json"},
      "Cookie": {c.cookie.Raw},
   }
   req.URL.Host = "therokuchannel.roku.com"
   req.URL.Path = "/api/v3/playback"
   req.URL.Scheme = "https"
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(Playback)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

type Cross_Site struct {
   cookie *http.Cookie // has own String method
   token string
}

func New_Cross_Site() (*Cross_Site, error) {
   req := http.Get()
   req.URL.Scheme = "https"
   // this has smaller body than www.roku.com
   req.URL.Host = "therokuchannel.roku.com"
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var site Cross_Site
   for _, cook := range res.Cookies() {
      if cook.Name == "_csrf" {
         site.cookie = cook
      }
   }
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   sep := []byte("\tcsrf:")
   if err := json.Cut(data, sep, &site.token); err != nil {
      return nil, err
   }
   return &site, nil
}
