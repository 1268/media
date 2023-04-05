package youtube

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/json"
)

type config struct {
   Innertube_Client_Name string
   Innertube_Client_Version string
}

func new_config() (*config, error) {
   req := http.Get()
   req.URL.Scheme = "https"
   req.URL.Host = "m.youtube.com"
   req.Header.Set("User-Agent", "iPad")
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var scan json.Scanner
   scan.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Sep = []byte("\nytcfg.set(")
   con := new(config)
   if err := json.NewDecoder(res.Body).Decode(con); err != nil {
      return nil, err
   }
   return con, nil
}
