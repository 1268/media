package plex

import (
   "fmt"
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
   var user1 User
   err := user1.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range watch_tests {
      match, err := user1.Match(Address{test.path})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(match)
      time.Sleep(time.Second)
   }
}
