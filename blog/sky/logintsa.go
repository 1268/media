package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.URL = &url.URL{}
   req.URL.Host = "gateway.prd.sky.ch"
   req.URL.Path = "/user/authentication/logintsa"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.Header["Accept-Encoding"] = []string{"compress"}
   req.Header["Accept-Language"] = []string{"en"}
   req.Header["Api-Version"] = []string{"1.0"}
   req.Header["Appversion"] = []string{"1.18.1.125"}
   req.Header["Content-Length"] = []string{"63"}
   req.Header["Content-Type"] = []string{"application/json; charset=UTF-8"}
   req.Header["Devicecode"] = []string{"ANDROID_INAPP"}
   req.Header["Devicename"] = []string{"Google AOSP on IA Emulator"}
   req.Header["Deviceplatform"] = []string{"ANDROID_INAPP"}
   req.Header["Environment"] = []string{"Sky"}
   req.Header["Macaddress"] = []string{"31f08176-dd2f-43a3-ae94-9e24190600fe"}
   req.Header["Osversion"] = []string{"28"}
   req.Header["Serialnumber"] = []string{"31f08176-dd2f-43a3-ae94-9e24190600fe"}
   req.Header["User-Agent"] = []string{"okhttp/5.0.0-alpha.10"}
   req.Header["Cookie"] = []string{
      "aws-waf-token=2e86b681-4c6d-40cd-9856-9ec0780664e5:HAoAkAsSO8kGAAAA:wWotxIx/qIxwEPx20cZJqorgSm4bt5YuAhntIxvP7HAXyKYgrnJD39XjU8Vlcwcb88umfrKppm+luczkW5DnyMk7l+eU7KbxOIi76foo8gRgpdS9e18/BwJVciM=",
   }
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

var body = strings.NewReader(`{"login":"EMAIL","password":"PASSWORD"}`)
