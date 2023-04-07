package soundcloud

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/json"
   "io"
)

type next_data struct {
   Runtime_Config struct {
      Client_ID string `json:"clientId"`
   } `json:"runtimeConfig"`
}

func new_next_data() (*next_data, error) {
   req := http.Get()
   req.URL.Scheme = "https"
   req.URL.Host = "m.soundcloud.com"
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   sep := []byte(` id="__NEXT_DATA__" type="application/json">`)
   next := new(next_data)
   if err := json.Cut(data, sep, next); err != nil {
      return nil, err
   }
   return next, nil
}
