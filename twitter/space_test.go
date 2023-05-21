package twitter

import (
   "fmt"
   "testing"
)

func Test_Guest(t *testing.T) {
   g, err := New_Guest()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println("Authorization: Bearer", bearer)
   fmt.Println("X-Guest-Token:", g.Guest_Token)
}
