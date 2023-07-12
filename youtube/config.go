package youtube

import (
   "154.pages.dev/encoding/json"
   "io"
   "net/http"
   "net/url"
   "path"
   "strconv"
)

type config struct {
   Innertube_API_Key string
   Innertube_Client_Name string
   Innertube_Client_Version string
}

func new_config() (*config, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "m.youtube.com",
   })
   req.Header.Set("User-Agent", "iPad")
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   con := new(config)
   {
      s, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      sep := []byte("\nytcfg.set(")
      if err := json.Cut(s, sep, con); err != nil {
         return nil, err
      }
   }
   return con, nil
}
