package youtube

import (
   "154.pages.dev/encoding/json"
   "io"
   "net/http"
)

type config struct {
   Innertube_API_Key string
   Innertube_Client_Name string
   Innertube_Client_Version string
}

func new_config() (*config, error) {
   req, err := http.NewRequest("GET", "https://m.youtube.com", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", "iPad")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   con := new(config)
   {
      sep := json.Split("\nytcfg.set(")
      s, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      if _, err := sep.After(s, con); err != nil {
         return nil, err
      }
   }
   return con, nil
}
