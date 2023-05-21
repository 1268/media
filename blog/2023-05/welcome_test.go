package twitter

import (
   "encoding/json"
   "os"
   "testing"
)

func Test_Twitter(t *testing.T) {
   flow, err := flow_welcome()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(flow)
}
