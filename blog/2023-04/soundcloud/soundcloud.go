package soundcloud

import (
   "2a.pages.dev/mech/soundcloud"
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
)

var client_IDs = []string{
   "SSdQ80vM8nLPhbDBylHl2JFK6ElhBr9B",
   "dbdsA8b6V6Lw7wzu1x0T4CLxt58yd4Bf",
}

func Resolve(ref, client_ID string) (*soundcloud.Track, error) {
   req := http.Get()
   req.URL.Host = "api-v2.soundcloud.com"
   req.URL.Path = "/resolve"
   req.URL.RawQuery = url.Values{
      "client_id": {client_ID},
      "url": {ref},
   }.Encode()
   req.URL.Scheme = "https"
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var solve struct {
      soundcloud.Track
   }
   if err := json.NewDecoder(res.Body).Decode(&solve); err != nil {
      return nil, err
   }
   return &solve.Track, nil
}
