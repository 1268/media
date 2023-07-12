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
var id_tests = []string{
   "https://youtube.com/shorts/9Vsdft81Q6w",
   "https://youtube.com/watch?v=XY-hOqcPGCY",
}

func Test_ID(t *testing.T) {
   var req Request
   for _, test := range id_tests {
      err := req.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(req.Video_ID)
   }
}

func Test_Player(t *testing.T) {
   var req Request
   req.Android()
   req.Video_ID = androids[0]
   for {
      for range [16]struct{}{} {
         p, err := req.Player(nil)
         if err != nil {
            t.Fatal(err)
         }
         if len(p.Streaming_Data.Adaptive_Formats) == 0 {
            t.Fatal(p)
         }
         if p.Video_Details.View_Count == 0 {
            t.Fatal(p)
         }
         time.Sleep(time.Second)
      }
      max_android.minor++
   }
}
