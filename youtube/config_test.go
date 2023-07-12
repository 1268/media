package youtube

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Config(t *testing.T) {
   con, err := new_config()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(con)
}
