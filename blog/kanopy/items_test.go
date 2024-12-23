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
      for _, ancestor := range video.AncestorVideoIds {
         items, err := token.items(ancestor)
         if err != nil {
            t.Fatal(err)
         }
         time.Sleep(time.Second)
         for i, item := range items {
            if i >= 1 {
               fmt.Println()
            }
            fmt.Println(&item)
         }
      }
   }
}

type transport struct{}

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   fmt.Println(req.URL)
   return http.DefaultTransport.RoundTrip(req)
}
