package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func main() {
   var req http.Request
   req.URL = new(url.URL)
   req.URL.Host = "link.theplatform.com"
   // <assetType>DASH_CENC</assetType>
   //req.URL.Path = "/s/dJ5BDC/media/guid/2198311517/tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD"
   
   // <assetType>DASH_CENC</assetType>
   // <assetType>Downloadable</assetType>
   //req.URL.Path = "/s/dJ5BDC/media/guid/2198311517/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_"
   
   // <assetType>DASH_CENC</assetType>
   // <assetType>StreamPack</assetType>
   req.URL.Path = "/s/dJ5BDC/media/guid/2198311517/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_"
   req.URL.Scheme = "http"
   req.Header = make(http.Header)
   power := power_set(
      "assetTypes=DASH_CENC",
      "assetTypes=Downloadable",
      "assetTypes=StreamPack",
      "format=smil",
      "formats=MPEG-DASH",
      "mbr=true",
   )
   for _, set := range power {
      req.URL.RawQuery = strings.Join(set, "&")
      if asset := req.URL.Query()["assetTypes"]; len(asset) >= 2 {
         continue
      }
      res, err := new(http.Client).Do(&req)
      if err != nil {
         panic(err)
      }
      if res.Header.Get("Content-Type") == "video/mp4" {
         fmt.Printf("pass %q\n", req.URL.RawQuery)
      } else {
         body, err := io.ReadAll(res.Body)
         if err != nil {
            panic(err)
         }
         if bytes.Contains(body, []byte("NoAssetTypeFormatMatches")) {
            fmt.Printf("fail %q\n", req.URL.RawQuery)
         } else {
            fmt.Printf("pass %q\n", req.URL.RawQuery)
         }
      }
      if err := res.Body.Close(); err != nil {
         panic(err)
      }
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
