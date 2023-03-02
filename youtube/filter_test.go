package youtube

import (
   "encoding/base64"
   "testing"
)

func (f Filters) to_string() string {
   raw := f.Marshal()
   return base64.StdEncoding.EncodeToString(raw)
}

func Test_Filter_Feature(t *testing.T) {
   filter := New_Filters()
   filter.Features(Features["Subtitles/CC"])
   if str := filter.to_string(); str != "EgIoAQ==" {
      t.Fatal(str)
   }
}

func Test_Filter_Sort(t *testing.T) {
   filter := New_Filters()
   filter.Sort_By(Sort_By["Rating"])
   if str := filter.to_string(); str != "CAE=" {
      t.Fatal(str)
   }
}
