package bandcamp

import (
   "fmt"
   "testing"
)

const art_ID = 3809045440

func Test_Image(t *testing.T) {
   for _, img := range Images {
      ref := img.URL(art_ID)
      fmt.Println(ref)
   }
}
