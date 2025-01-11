package kanopy

import (
   "fmt"
   "net/http"
   "os"
   "testing"
   "time"
)

func TestItems(t *testing.T) {
   http.DefaultClient.Transport = transport{}
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token web_token
   err = token.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      video, err := token.video(test.video_id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      if len(video.AncestorVideoIds) <= 1 {
         continue
      }
      // episodes
      items, err := token.items(video.AncestorVideoIds[0])
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      fmt.Println(items.item(video.VideoId))
      // seasons
      items, err = token.items(video.AncestorVideoIds[1])
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      fmt.Println(items.item(video.AncestorVideoIds[0]))
   }
}

type transport struct{}

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   fmt.Println(req.URL)
   return http.DefaultTransport.RoundTrip(req)
}
