package plex

import (
   "41.neocities.org/text"
   "fmt"
   "reflect"
   "testing"
   "time"
)

var watch_tests = []struct{
   key_id string
   path string
   url string
}{
   {
      key_id: "4310a7c8094acab73fceab9d5494f36f",
      path: "/movie/cruel-intentions",
      url: "watch.plex.tv/movie/cruel-intentions",
   },
   {
      key_id: "", // no DRM
      path: "/show/broadchurch/season/3/episode/5",
      url: "watch.plex.tv/show/broadchurch/season/3/episode/5",
   },
}

func TestMatch(t *testing.T) {
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
      name, err := text.Name(&Namer{match})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}

func TestVideo(t *testing.T) {
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

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Address{},
   Anonymous{},
   DiscoverMatch{},
   MediaPart{},
   Namer{},
   OnDemand{},
   Url{},
}
