package max

import (
   "encoding/json"
   "net/http"
)

func (b *BoltToken) Initiate() (*LinkInitiate, error) {
   req, err := http.NewRequest(
      "POST", prd_api+"/authentication/linkDevice/initiate", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "cookie":        {"st=" + b.St},
      "x-device-info": {device_info},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   link := &LinkInitiate{}
   err = json.NewDecoder(resp.Body).Decode(link)
   if err != nil {
      return nil, err
   }
   return link, nil
}

type LinkInitiate struct {
   Data struct {
      Attributes struct {
         LinkingCode string
         TargetUrl   string
      }
   }
}
