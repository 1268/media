package rakuten

import (
   "fmt"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   for _, test := range web_tests {
      var content *gizmo_content
      if test.out.season_id != "" {
         season, err := test.out.season()
         if err != nil {
            t.Fatal(err)
         }
         var ok bool
         content, ok = test.out.content(season)
         if !ok {
            t.Fatal(season)
         }
      } else {
         var err error
         content, err = test.out.movie()
         if err != nil {
            t.Fatal(err)
         }
      }
      fmt.Print(content, "\n\n")
      time.Sleep(time.Second)
   }
}

func TestAddress(t *testing.T) {
   for _, test := range web_tests {
      var out address
      err := out.Set(test.in)
      if err != nil {
         t.Fatal(err)
      }
      if out != test.out {
         t.Fatal(test)
      }
   }
}
