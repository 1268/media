package twitter

import (
   "encoding/json"
   "os"
   "testing"
)

func Test_Search(t *testing.T) {
   http.Default_Client.Log_Level = 2
   access, err := access_token()
   if err != nil {
      t.Fatal(err)
   }
   guest, err := guest_token(access)
   if err != nil {
      t.Fatal(err)
   }
   f, err := welcome(access, guest)
   if err != nil {
      t.Fatal(err)
   }
   if err := f.next_link(access, guest); err != nil {
      t.Fatal(err)
   }
   s, err := f.open_account().search("filter:spaces")
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(s)
}
