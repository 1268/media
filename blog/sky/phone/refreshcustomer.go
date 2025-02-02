package phone

import (
   "io"
   "net/http"
   "net/url"
   "strings"
)

/*
x-forwarded-for fail
mullvad.net fail
proxy-seller.com pass
*/
func refresh_customer() (*http.Response, error) {
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.URL = &url.URL{}
   req.URL.Host = "gateway.prd.sky.ch"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/json; charset=UTF-8"}
   req.Header["Macaddress"] = []string{"66a427e0-2350-465a-8ad8-4c93def1299b"}
   req.URL.Path = "/user/authentication/refreshcustomer/1938441"
   return http.DefaultClient.Do(&req)
}

var body = strings.NewReader(`"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxOTM4NDQxIiwibWFjQWRkcmVzcyI6IjY2YTQyN2UwLTIzNTAtNDY1YS04YWQ4LTRjOTNkZWYxMjk5YiIsInRpbWUiOiIyMDI1LDIsMSwxMiwwLDAiLCJpc3MiOiJodHRwczovL3d3dy5za3kuY2giLCJhdWQiOiJTa3kgVXNlcnMifQ.fJ8T4oATKAyFWQKte6uYs2WZl_QEP7uNVTKDAISIql4"`)
