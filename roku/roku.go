package roku

import (
   "154.pages.dev/encoding/json"
   "io"
   "net/http"
   "net/url"
)

func New_Cross_Site() (*Cross_Site, error) {
   // this has smaller body than www.roku.com
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "therokuchannel.roku.com",
   })
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
   {
      s, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      sep := []byte("\tcsrf:")
      if err := json.Cut(s, sep, &site.token); err != nil {
         return nil, err
      }
   }
   return &site, nil
}

func (c Cross_Site) Playback(id string) (*Playback, error) {
   body := func(r *http.Request) error {
      m := map[string]string{
         "mediaFormat": "mpeg-dash",
         "providerId": "rokuavod",
         "rokuId": id,
      }
      b, err := json.MarshalIndent(m, "", " ")
      if err != nil {
         return err
      }
      r.Body_Bytes(b)
      return nil
   }
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "therokuchannel.roku.com",
      Path: "/api/v3/playback",
   })
   // we could use Request.AddCookie, but we would need to call it after this,
   // otherwise it would be clobbered
   req.Header = http.Header{
      "CSRF-Token": {c.token},
      "Content-Type": {"application/json"},
      "Cookie": {c.cookie.Raw},
   }
   err := body(req)
   if err != nil {
      return nil, err
   }
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

type Playback struct {
   DRM struct {
      Widevine struct {
         License_Server string `json:"licenseServer"`
      }
   }
}

func (p Playback) Request_URL() string {
   return p.DRM.Widevine.License_Server
}

func (Playback) Request_Header() http.Header {
   return nil
}

func (Playback) Request_Body(buf []byte) ([]byte, error) {
   return buf, nil
}

func (Playback) Response_Body(buf []byte) ([]byte, error) {
   return buf, nil
}
