package roku

import (
   "154.pages.dev/media"
   "154.pages.dev/widevine"
   "encoding/base64"
   "encoding/json"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Playback(t *testing.T) {
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
