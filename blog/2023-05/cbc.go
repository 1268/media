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
   val := make(url.Values)
   req.URL.Scheme = "https"
   req.Header["X-Claims-Token"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJIYXNIRCI6IkZhbHNlIiwiVGllciI6Ik1lbWJlciIsIkhhc0FkcyI6IlRydWUiLCJSY0lkIjoiZmE4N2M1NDAtMDRlNy00NDYxLTk3N2EtYzhiZDQ3MGQxNDBhIiwiTWF4aW11bU51bWJlck9mU3RyZWFtcyI6IjUiLCJSY1RlbGNvIjoiYXVjdW4iLCJQcGlkIjoiNGVkMmUyMWRhMjRmOTg1Yzg2ODJiMjJlYTQwMjI0NTAyODQ3ODE0MjAzNWZhZjVjNjk2MzRlNWFiZGU1ZDIzYSIsImV4cCI6MTY4NDEwODE5NX0.u4uuqv9PBlB8DWUzhuXAIJ3pnpGmoyxLtyIDnoBVANQ"}
   req.Header["X-Forwarded-For"] = []string{"99.224.0.0"}
   val["idMedia"] = []string{"929078"}
   val["appCode"] = []string{"gem"}
   val["manifestType"] = []string{"desktop"}
   val["output"] = []string{"json"}
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
