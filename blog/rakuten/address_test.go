package rakuten

import "testing"

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
