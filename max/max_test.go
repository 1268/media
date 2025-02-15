package max

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = []struct {
   url        string
   video_type string
   key_id     []string
}{
   {
      url:        "play.max.com/video/watch/5c762883-279e-40ed-ab84-43fdda9d88a0/560abdc4-ee5e-4f86-807e-38bb9feabe0e",
      video_type: "MOVIE",
      key_id: []string{
         "AQC1NR9S5CJX8MEgkYbXpg==",
         "AQFz3ZsVFjESfMh2rISgjw==",
         "AQLexzSxi5gJMbgkQogYJQ==",
         "AQW5UW421beBH+jIn3XASw==",
      },
   },
   {
      video_type: "EPISODE",
      url:        "play.max.com/video/watch/28ae9450-8192-4277-b661-e76eaad9b2e6/e19442fb-c7ac-4879-8d50-a301f613cb96",
      key_id:     nil,
   },
}

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
   data, err := os.ReadFile(home + "/max.txt")
   if err != nil {
      t.Fatal(err)
   }
   var login LinkLogin
   err = login.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var watch WatchUrl
      watch.UnmarshalText([]byte(test.url))
      play, err := login.Playback(&watch)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", play.Fallback)
      time.Sleep(time.Second)
   }
}
