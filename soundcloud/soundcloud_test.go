package soundcloud

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

const (
   id = 936653761
   address = "https://soundcloud.com/kino-scmusic/mqymd53jtwag"
)

func Test_Resolve(t *testing.T) {
   track, err := Resolve(address)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(track)
}

func Test_Track(t *testing.T) {
   track, err := New_Track(id)
   if err != nil {
      t.Fatal(err)
   }
   media, err := track.Progressive()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   enc.Encode(media)
}
