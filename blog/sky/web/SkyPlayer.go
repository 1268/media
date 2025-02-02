package main

import (
   //"bytes"
   "net/http"
   "net/url"
   "os"
)

/*
x-forwarded-for fail
mullvad.net fail
proxy-seller.com pass
*/
func main() {
   var req http.Request
   req.Header = http.Header{}
   req.URL = &url.URL{}
   req.URL.Host = "show.sky.ch"
   req.URL.Path = "/de/SkyPlayerAjax/SkyPlayer"
   req.URL.Scheme = "https"
   values := url.Values{}
   values["id"] = []string{"2035"}
   values["contentType"] = []string{"2"}
   req.URL.RawQuery = values.Encode()
   req.Header["X-Requested-With"] = []string{"XMLHttpRequest"}
   req.Header.Set(
      "cookie",
      "_ASP.NET_SessionId_=oiqw4d0k0x1ymoap0rlmbntj",
   )
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
