package twitter

import (
   "os"
   "testing"
)

func Test_Twitter(t *testing.T) {
   res, err := task_one()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
