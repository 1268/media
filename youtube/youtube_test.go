package youtube

import (
   "fmt"
   "testing"
)

var id_tests = []string{
   "https://youtube.com/shorts/9Vsdft81Q6w",
   "https://youtube.com/watch?v=XY-hOqcPGCY",
}

func Test_ID(t *testing.T) {
   for _, test := range id_tests {
      var id string
      err := Video_ID(test, &id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(id)
   }
}
