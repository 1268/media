package twitter

import (
   "encoding/json"
   "os"
   "testing"
)

func Test_Search(t *testing.T) {
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
   s, err := flow.open_account().search("filter:spaces")
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(s)
}
