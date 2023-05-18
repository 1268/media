package amc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "net/url"
   "strings"
)

func (a Auth) Playback(ref string) (*Playback, error) {
   path_body := func(r *http.Request) error {
      _, nID, found := strings.Cut(ref, "--")
      if !found {
         return errors.New("nid not found")
      }
      r.URL.Path += nID
      var p playback_request
      p.Ad_Tags.Mode = "on-demand"
      p.Ad_Tags.URL = "-"
      b, err := json.MarshalIndent(p, "", " ")
      if err != nil {
         return err
      }
      r.Body_Bytes(b)
      return nil
   }
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "gw.cds.amcn.com",
      Path: "/playback-id/api/v1/playback/",
   })
   err := path_body(req)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Data.Access_Token},
      "Content-Type": {"application/json"},
      "X-Amcn-Device-Ad-ID": {"-"},
      "X-Amcn-Language": {"en"},
      "X-Amcn-Network": {"amcplus"},
      "X-Amcn-Platform": {"web"},
      "X-Amcn-Service-ID": {"amcplus"},
      "X-Amcn-Tenant": {"amcn"},
      "X-Ccpa-Do-Not-Sell": {"doNotPassData"},
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var play Playback
   if err := json.NewDecoder(res.Body).Decode(&play.body); err != nil {
      return nil, err
   }
   play.head = res.Header
   return &play, nil
}

func (Playback) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (p Playback) Request_Header() http.Header {
   token := p.head.Get("X-AMCN-BC-JWT")
   head := make(http.Header)
   head.Set("bcov-auth", token)
   return head
}

type Playback struct {
   head http.Header
   body struct {
      Data struct {
         Playback_JSON_Data struct {
            Sources []Source
         } `json:"playbackJsonData"`
      }
   }
}

func (p Playback) Request_URL() string {
   return p.Source().Key_Systems.Widevine.License_URL
}

func (p Playback) Source() *Source {
   for _, item := range p.body.Data.Playback_JSON_Data.Sources {
      if item.Type == "application/dash+xml" {
         return &item
      }
   }
   return nil
}
