package max

import (
   "154.pages.dev/http/option"
   "os"
   "testing"
)

const lotus = "https://www.max.com/a/video/the-white-lotus-s1-e1"

func Test_Web(t *testing.T) {
   option.No_Location()
   option.Verbose()
   res, err := video(lotus)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   file, err := os.Create("white-lotus-s1-e1.html")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   file.ReadFrom(res.Body)
}
