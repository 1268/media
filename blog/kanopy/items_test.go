package kanopy

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestItems(t *testing.T) {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token web_token
   err = token.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      items, err := token.items(test.video_id)
      if err != nil {
         t.Fatal(err)
      }
      for _, item := range items.List {
         fmt.Printf("%+v\n", item)
      }
      time.Sleep(time.Second)
   }
}
