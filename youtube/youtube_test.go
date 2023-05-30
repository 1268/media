package youtube

import (
   "fmt"
   "testing"
   "time"
)

func Test_Player(t *testing.T) {
   for {
      for range [16]struct{}{} {
         p, err := Android().Player(androids[0], nil)
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
var id_tests = []string{
   "https://youtube.com/shorts/9Vsdft81Q6w",
   "https://youtube.com/watch?v=XY-hOqcPGCY",
}

func Test_ID(t *testing.T) {
   for _, test := range id_tests {
      var id string
      err := Video_ID(test, &id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(id)
   }
}
