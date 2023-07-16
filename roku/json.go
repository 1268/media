package roku

import (
   "bytes"
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

func (c Cross_Site) Playback(id string) (*Playback, error) {
   body, err := func() ([]byte, error) {
      m := map[string]string{
         "mediaFormat": "mpeg-dash",
         "providerId": "rokuavod",
         "rokuId": id,
      }
      return json.MarshalIndent(m, "", " ")
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://therokuchannel.roku.com/api/v3/playback",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   // we could use Request.AddCookie, but we would need to call it after this,
   // otherwise it would be clobbered
   req.Header = http.Header{
      "CSRF-Token": {c.token},
      "Content-Type": {"application/json"},
      "Cookie": {c.cookie.Raw},
   }
   res, err := new(http.Transport).RoundTrip(req)
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

func New_Content(id string) (*Content, error) {
   req, err := http.NewRequest(
      "GET", "https://therokuchannel.roku.com/api/v2/homescreen/content", nil,
   )
   if err != nil {
      return nil, err
   }
   {
      include := []string{
         "episodeNumber",
         "releaseDate",
         "seasonNumber",
         "series.title",
         "title",
         "viewOptions",
      }
      expand := url.URL{
         Scheme: "https",
         Host: "content.sr.roku.com",
         Path: "/content/v1/roku-trc/" + id,
         RawQuery: url.Values{
            "expand": {"series"},
            "include": {strings.Join(include, ",")},
         }.Encode(),
      }
      homescreen := url.PathEscape(expand.String())
      req.URL = req.URL.JoinPath(homescreen)
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var con Content
   if err := json.NewDecoder(res.Body).Decode(&con.s); err != nil {
      return nil, err
   }
   return &con, nil
}
