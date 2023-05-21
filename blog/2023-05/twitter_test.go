package twitter

import (
   "encoding/json"
   "os"
   "testing"
)

func Test_Twitter(t *testing.T) {
   g, err := new_guest()
   if err != nil {
      t.Fatal(err)
   }
   flow, err := flow_welcome(g)
   if err != nil {
      t.Fatal(err)
   }
   if err := flow.next_link(g); err != nil {
      t.Fatal(err)
   }
   account := flow.open_account()
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(account)
}
