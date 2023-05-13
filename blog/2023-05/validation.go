package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "services.radio-canada.ca"
   req.URL.Path = "/media/validation/v2"
   req.URL.Scheme = "https"
   req.Header["X-Forwarded-For"] = []string{"99.224.0.0"}
   val := make(url.Values)
   val["appCode"] = []string{"gem"}
   val["idMedia"] = []string{"958273"}
   val["output"] = []string{"json"}
   // "name": "adsEnabled", "value": false
   val["manifestType"] = []string{"desktop"}
   req.URL.RawQuery = val.Encode()
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
