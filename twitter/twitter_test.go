// Support for this software's development was paid for by
// Fredrick R. Brennan's Modular Font Editor K Foundation, Inc.
package twitter

import (
   "154.pages.dev/http/option"
   "encoding/json"
   "os"
   "testing"
)

// twitter.com/1500tasvir/status/1577533879283585025
const id = 1577533879283585025

func Test_Tweet(t *testing.T) {
   option.No_Location()
   option.Verbose()
   tweet, err := New_Tweet_Result(id)
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   enc.Encode(tweet.Video_MP4())
}
