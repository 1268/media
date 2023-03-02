package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

const image_test = "UpNXI3_ctAc"

func Test_Image(t *testing.T) {
   for _, img := range Images {
      ref := img.Address(image_test)
      fmt.Println("HEAD", ref)
      res, err := http.Head(ref)
      if err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != http.StatusOK {
         t.Fatal(res.Status)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
