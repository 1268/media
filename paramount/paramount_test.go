package paramount

import (
   "fmt"
   "testing"
   "time"
)

var tests = []struct{
   content_id string
   key_id string
   location []string
   url string
}{
   {
      content_id: "wjQ4RChi6BHHu4MVTncppVuCwu44uq2Q",
      key_id: "twwHMCIoRtaIS+/clhhs9A==",
      location: []string{"US"},
      url: "paramountplus.com/movies/video/wjQ4RChi6BHHu4MVTncppVuCwu44uq2Q",
   },
   {
      content_id: "esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
      key_id: "H94BVNcqT0WRKzTwzgd36w==",
      location: []string{"US"},
      url: "paramountplus.com/shows/video/esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
   },
   {
      content_id: "rZ59lcp4i2fU4dAaZJ_iEgKqVg_ogrIf",
      key_id: "Sryog4HeT2CLHx38NftIMA==",
      location: []string{"US"},
      url: "cbs.com/shows/video/rZ59lcp4i2fU4dAaZJ_iEgKqVg_ogrIf",
   },
   {
      content_id: "N5ySQTDzhLW2YyWGWuZvCb_wGsCQ_jCJ",
      key_id: "w4pjkntES4yAxVEfcL0azQ==",
      location: []string{"CA", "GB"},
      url: "paramountplus.com/shows/video/N5ySQTDzhLW2YyWGWuZvCb_wGsCQ_jCJ",
   },
   {
      content_id: "WNujiS5PHkY5wN9doNY6MSo_7G8uBUcX",
      key_id: "bsT01+Q1Ta+39TayayKhBg==",
      location: []string{"AU"},
      url: "paramountplus.com/shows/video/WNujiS5PHkY5wN9doNY6MSo_7G8uBUcX",
   },
   {
      content_id: "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
      key_id: "BsO37qHORXefruKryNAaVQ==",
      location: []string{"AU", "GB"},
      url: "paramountplus.com/movies/video/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
   },
}

func Test(t *testing.T) {
   for _, test1 := range tests {
      session, err := ComCbsApp.Session(test1.content_id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(session)
      time.Sleep(time.Second)
   }
}
