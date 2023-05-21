package twitter

import (
   "net/http/httputil"
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
   res, err := flow.open_account().search()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   data, err := httputil.DumpResponse(res, true)
   if err != nil {
      t.Fatal(err)
   }
   os.Stdout.Write(data)
}
