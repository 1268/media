package twitter

import (
   "2a.pages.dev/rosso/http"
   "os"
   "testing"
)

func Test_Twitter(t *testing.T) {
   http.Default_Client.Log_Level = 2
   g, err := New_Guest()
   if err != nil {
      t.Fatal(err)
   }
   flow, err := g.flow_welcome()
   if err != nil {
      t.Fatal(err)
   }
   res, err := g.next_link(flow)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
