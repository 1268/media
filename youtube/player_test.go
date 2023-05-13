package youtube

import (
   "2a.pages.dev/rosso/http"
   "testing"
   "time"
)

func Test_Player(t *testing.T) {
   http.Default_Client.Log_Level = 2
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
