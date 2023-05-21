package twitter

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func token() {
   req_body := strings.NewReader(`grant_type=client_credentials`)
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "api.twitter.com"
   req.URL.Path = "/oauth2/token"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.SetBasicAuth(
      "3nVuSoBZnx6U4vzUxf5w",
      "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys",
   )
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
