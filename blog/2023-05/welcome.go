package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
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
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}

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
