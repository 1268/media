package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

const image_test = "UpNXI3_ctAc"

func Test_Image(t *testing.T) {
   req, err := http.NewRequest("HEAD", "", nil)
   if err != nil {
      t.Fatal(err)
   }
   for _, img := range Images {
      req.URL = img.URL(image_test)
      fmt.Println("HEAD", req.URL)
      res, err := new(http.Transport).RoundTrip(req)
      if err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != http.StatusOK {
         t.Fatal(res.Status)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
