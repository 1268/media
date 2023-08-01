package max

import (
   "154.pages.dev/http/option"
   "encoding/json"
   "os"
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
   med, err := page.media()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(med)
}
