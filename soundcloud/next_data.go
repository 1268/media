package soundcloud

import (
   "154.pages.dev/encoding/json"
   "io"
   "net/http"
)

func new_next_data() (*next_data, error) {
   res, err := http.Get("https://m.soundcloud.com")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   before := []byte(` id="__NEXT_DATA__" type="application/json">`)
   _, text = json.Cut(text, before, nil)
   next := new(next_data)
   if err := json.Unmarshal(text, next); err != nil {
      return nil, err
   }
   return next, nil
}

type next_data struct {
   Runtime_Config struct {
      Client_ID string `json:"clientId"`
   } `json:"runtimeConfig"`
}
