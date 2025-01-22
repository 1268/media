package max

import "net/http"

const (
   device_info  = "!/!(!/!;!/!;!/!)"
   disco_client = "!:!:beam:!"
   prd_api      = "https://default.prd.api.discomax.com"
)

type BoltToken struct {
   St string
}

func (b *BoltToken) New() error {
   req, err := http.NewRequest("", prd_api+"/token?realm=bolt", nil)
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "x-device-info":  {device_info},
      "x-disco-client": {disco_client},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   for _, cookie := range resp.Cookies() {
      if cookie.Name == "st" {
         b.St = cookie.Value
         return nil
      }
   }
   return http.ErrNoCookie
}
