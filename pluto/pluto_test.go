package pluto

import (
   "41.neocities.org/widevine"
   "bytes"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

var tests = []struct{
   id     string
   key_id string
   url    string
}{
   {
      id:     "5c4bb2b308d10f9a25bbc6af",
      key_id: "AAAAAGbZBRrrxvnmpuNLhg==",
      url:    "pluto.tv/on-demand/movies/5c4bb2b308d10f9a25bbc6af",
   },
   {
      url:    "pluto.tv/on-demand/series/66d0bb64a1c89200137fb0e6/episode/66fb16fda2922a00135e87f7",
      id:     "66fb16fda2922a00135e87f7",
      key_id: "",
   },
}

func TestClip(t *testing.T) {
   clip, err := OnDemand{Id: video_test.id}.Clip()
   if err != nil {
      t.Fatal(err)
   }
   manifest, ok := clip.Dash()
   if !ok {
      t.Fatal("EpisodeClip.Dash")
   }
   manifest.Scheme = Base[0].Scheme
   manifest.Host = Base[0].Host
   fmt.Printf("%+v\n", manifest)
}

func TestVideo(t *testing.T) {
   var web Address
   web.Set(video_test.url)
   video, err := web.Video("")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}
