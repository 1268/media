package pluto

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

func TestLicense(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   var pssh widevine.Pssh
   pssh.KeyId, err = hex.DecodeString(video_test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Module
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(Client{}, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

// the slug is useful as it sometimes contains the year, but its not worth
// parsing since its sometimes missing
var video_test = struct{
   id     string
   key_id string
   url    string
}{
   id:     "675a0fa22678a50014690c3f",
   key_id: "AAAAAGdaD6FuwTSRB/+yHg==",
   url:    "pluto.tv/on-demand/movies/675a0fa22678a50014690c3f",
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

func TestAddress(t *testing.T) {
   var web Address
   err := web.Set(video_test.url)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(web)
}

func TestVideo(t *testing.T) {
   var web Address
   web.Set(video_test.url)
   video, err := web.Video("")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
   name := text.Name(Namer{video})
   fmt.Printf("%q\n", name)
}
