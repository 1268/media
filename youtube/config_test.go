package youtube

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
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

func Test_Format(t *testing.T) {
   var r Request
   r.Android()
   r.Video_ID = androids[0]
   play, err := r.Player(nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play)
}
