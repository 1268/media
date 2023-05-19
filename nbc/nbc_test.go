package nbc

import (
   "fmt"
   "testing"
   "time"
)

func Test_Video(t *testing.T) {
   for _, guid := range guids {
      meta, err := New_Metadata(guid)
      if err != nil {
         t.Fatal(err)
      }
      vid, err := meta.Video()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", vid)
      time.Sleep(time.Second)
   }
}
