package twitter

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
)

func flow_welcome() (*task, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/onboarding/task.json",
      RawQuery: "flow_name=welcome",
   })
   req.Header = http.Header{
      "Authorization": {"Bearer AAAAAAAAAAAAAAAAAAAAAFXzAwAAAAAAMHCxpeSDG1gLNLghVe8d74hl6k4%3DRUMF4xAQLsbeBhTSRrCiQpJtxoGWeyHrDb5te2jpGskWDFW82F"},
      "X-Guest-Token": {"1660058730438836226"},
      "Content-Type": {"application/json"},
      "User-Agent": {"TwitterAndroid/99"},
   }
   var t task
   {
      t.Input_Flow_Data.Flow_Context.Start_Location.Location = "splash_screen"
      b, err := json.Marshal(t)
      if err != nil {
         return nil, err
      }
      req.Body_Bytes(b)
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
      return nil, err
   }
   return &t, nil
}
