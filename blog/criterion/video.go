package criterion

import "net/http"

func (a auth_token) video() (*http.Response, error) {
   req, err := http.NewRequest("", "https://api.vhx.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v2/sites/59054/videos/455774"
   req.Header.Set("authorization", "Bearer " + a.v.AccessToken)
   return http.DefaultClient.Do(req)
}
