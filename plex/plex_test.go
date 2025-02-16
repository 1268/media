package plex

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

var watch_tests = []struct{
   key_id string
   path string
   url string
}{
   {
      key_id: "47yjtPAH46ndYmeLgBQbfw==",
      path: "/watch/movie/southpaw-2015",
      url: "watch.plex.tv/watch/movie/southpaw-2015",
   },
   {
      url: "watch.plex.tv/movie/limitless",
      path: "/movie/limitless",
      key_id: "", // no DRM
   },
   {
      key_id: "", // no DRM
      path: "/show/broadchurch/season/3/episode/5",
      url: "watch.plex.tv/show/broadchurch/season/3/episode/5",
   },
}

func Test(t *testing.T) {
   var user Anonymous
   err := user.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range watch_tests {
      match, err := user.Match(&Address{test.path})
      if err != nil {
         t.Fatal(err)
      }
      video, err := user.Video(match, "")
      if err != nil {
         t.Fatal(err)
      }
      for _, media := range video.Media {
         for _, part := range media.Part {
            fmt.Println(part.Key)
            fmt.Println(part.License)
         }
      }
      time.Sleep(time.Second)
   }
}
