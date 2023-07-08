package soundcloud

import (
   "2a.pages.dev/rosso/http"
   "encoding.pages.dev/json"
   "io"
   "net/url"
)

type next_data struct {
   Runtime_Config struct {
      Client_ID string `json:"clientId"`
   } `json:"runtimeConfig"`
}

func new_next_data() (*next_data, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "m.soundcloud.com",
   })
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   next := new(next_data)
   {
      s, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      sep := []byte(` id="__NEXT_DATA__" type="application/json">`)
      if err := json.Cut(s, sep, next); err != nil {
         return nil, err
      }
   }
   return next, nil
}
