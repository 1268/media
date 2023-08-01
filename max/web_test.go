package max

import (
   "154.pages.dev/http/option"
   "strings"
   "testing"
)

const lotus = "https://www.max.com/a/video/the-white-lotus-s1-e1"

func Test_Web(t *testing.T) {
   option.No_Location()
   option.Verbose()
   next, err := new_next_data(lotus)
   if err != nil {
      t.Fatal(err)
   }
   page, err := next.page_data()
   if err != nil {
      t.Fatal(err)
   }
   m, err := page.media()
   if err != nil {
      t.Fatal(err)
   }
   if !strings.HasSuffix(m.Media.Desktop.Unprotected.Unencrypted.URL, "master_cl_de.m3u8") {
      t.Fatal(m)
   }
}
