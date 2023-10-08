package roku

import (
   "154.pages.dev/http/option"
   "154.pages.dev/media"
   "154.pages.dev/widevine"
   "encoding/base64"
   "encoding/hex"
   "encoding/json"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Playback(t *testing.T) {
   option.No_Location()
   option.Trace()
   site, err := New_Cross_Site()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   for _, test := range tests {
      play, err := site.Playback(test.playback_ID)
      if err != nil {
         t.Fatal(err)
      }
      enc.Encode(play)
      time.Sleep(time.Second)
   }
}

func Test_Post(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      if test.pssh != "" {
         pssh, err := base64.StdEncoding.DecodeString(test.pssh)
         if err != nil {
            t.Fatal(err)
         }
         mod, err := widevine.New_Module(private_key, client_ID, pssh)
         if err != nil {
            t.Fatal(err)
         }
         site, err := New_Cross_Site()
         if err != nil {
            t.Fatal(err)
         }
         play, err := site.Playback(test.playback_ID)
         if err != nil {
            t.Fatal(err)
         }
         key, err := mod.Key(play)
         if err != nil {
            t.Fatal(err)
         }
         if hex.EncodeToString(key) != test.key {
            t.Fatal(key)
         }
      }
   }
}
