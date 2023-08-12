package youtube

import (
   "154.pages.dev/encoding/json"
   "io"
   "net/http"
)

func new_config() (*config, error) {
   req, err := http.NewRequest("GET", "https://m.youtube.com", nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", "iPad")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   _, text = json.Cut(text, []byte("\nytcfg.set("), nil)
   con := new(config)
   if err := json.Unmarshal(text, con); err != nil {
      return nil, err
   }
   return con, nil
}

type config struct {
   Innertube_API_Key string
   Innertube_Client_Name string
   Innertube_Client_Version string
}
