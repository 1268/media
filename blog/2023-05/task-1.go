package twitter

import (
   "io"
   "net/http"
   "net/url"
   "strings"
)

func task_one() (*http.Response, error) {
   var req_body = strings.NewReader(`
   {
      "flow_token": null,
      "input_flow_data": {
         "flow_context": {
            "start_location": {
               "location": "splash_screen"
            }
         }
      }
   }
   `)
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "api.twitter.com"
   req.URL.Path = "/1.1/onboarding/task.json"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["flow_name"] = []string{"welcome"}
   req.Header["Authorization"] = []string{"Bearer AAAAAAAAAAAAAAAAAAAAAFXzAwAAAAAAMHCxpeSDG1gLNLghVe8d74hl6k4%3DRUMF4xAQLsbeBhTSRrCiQpJtxoGWeyHrDb5te2jpGskWDFW82F"}
   req.Header["X-Guest-Token"] = []string{"1660058730438836226"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["User-Agent"] = []string{"TwitterAndroid/99"}
   req.URL.RawQuery = val.Encode()
   req.Body = io.NopCloser(req_body)
   return new(http.Transport).RoundTrip(&req)
}
