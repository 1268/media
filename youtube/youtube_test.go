package youtube

import (
   "fmt"
   "testing"
   "time"
)

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
