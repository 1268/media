package youtube

import (
   "2a.pages.dev/rosso/http"
   "testing"
   "time"
)

func Test_Player(t *testing.T) {
   http.Default_Client.Log_Level = 2
   var req Request
   req.Content_Check_OK = true
   req.Context.Client.Name = "ANDROID"
   video_ID := androids[0]
   for range [16]struct{}{} {
      req.Context.Client.Version = max_android
      {
         p, err := req.Player(video_ID, nil)
         if err != nil {
            t.Fatal(err)
         }
         if len(p.Streaming_Data.Adaptive_Formats) == 0 {
            t.Fatal(p)
         }
         if p.Video_Details.View_Count == 0 {
            t.Fatal(p)
         }
      }
      time.Sleep(time.Second)
      req.Context.Client.Version = min_android
      {
         p, err := req.Player(video_ID, nil)
         if err != nil {
            t.Fatal(err)
         }
         if len(p.Streaming_Data.Adaptive_Formats) == 0 {
            t.Fatal(p)
         }
         if p.Video_Details.View_Count == 0 {
            t.Fatal(p)
         }
      }
      time.Sleep(time.Second)
   }
}
