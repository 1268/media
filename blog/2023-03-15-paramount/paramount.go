package main

import (
   "fmt"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func main() {
   var req http.Request
   req.URL = new(url.URL)
   req.URL.Host = "link.theplatform.com"
   req.URL.Path = "/s/dJ5BDC/media/guid/2198311517/tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD"
   req.URL.Scheme = "http"
   req.Header = make(http.Header)
   power := power_set(
      "assetTypes=DASH_CENC",
      "formats=MPEG-DASH",
   )
   for _, set := range power {
      req.URL.RawQuery = strings.Join(set, "&")
      res, err := new(http.Transport).RoundTrip(&req)
      if err != nil {
         panic(err)
      }
      if err := res.Body.Close(); err != nil {
         panic(err)
      }
      fmt.Println(res.Status, set)
      time.Sleep(time.Second)
   }
}

func power_set(a ...string) [][]string {
   b := [][]string{{}}
   for _, c := range a {
      for _, d := range b {
         b= append(b, append(d, c))
      }
   }
   return b
}
