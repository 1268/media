package youtube

import (
   "encoding/base64"
   "testing"
)

func (p Params) to_string() string {
   raw := p.Marshal()
   return base64.StdEncoding.EncodeToString(raw)
}

func Test_Filter_Feature(t *testing.T) {
   param := New_Params()
   param.Features(Features["Subtitles/CC"])
   if s := param.to_string(); s != "EgIoAQ==" {
      t.Fatal(s)
   }
}

func Test_Filter_Sort(t *testing.T) {
   param := New_Params()
   param.Sort_By(Sort_By["Rating"])
   if s := param.to_string(); s != "CAE=" {
      t.Fatal(s)
   }
}
