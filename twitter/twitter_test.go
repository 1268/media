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
   tweet, err := new_tweet_result(id)
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   enc.Encode(tweet)
}
