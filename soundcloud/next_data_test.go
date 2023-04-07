package soundcloud

import (
   "encoding/json"
   "os"
   "testing"
)

func Test_Next_Data(t *testing.T) {
   next, err := new_next_data()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(next)
}
