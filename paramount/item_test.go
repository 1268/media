package paramount

import (
   "2a.pages.dev/mech/widevine"
   "encoding/base64"
   "encoding/json"
   "fmt"
   "net/http"
   "os"
   "testing"
   "time"
)

func Test_Item(t *testing.T) {
   token, err := new_app_token()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      item, err := token.item(test.content_ID)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(item.Name())
      time.Sleep(time.Second)
   }
}
