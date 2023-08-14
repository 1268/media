package youtube

import (
   "encoding/base64"
   "testing"
)

func Test_Filter_Sort(t *testing.T) {
   var p parameters
   p.Sort_By = values["SORT BY"]["Rating"]
   if s := p.encode(); s != "CAE=" {
      t.Fatal(s)
   }
}

func Test_Filter_Feature(t *testing.T) {
   var p parameters
   p.Filter = new(filter)
   p.Filter.Features = []uint64{values["FEATURES"]["Subtitles/CC"]}
   if s := p.encode(); s != "EgIoAQ==" {
      t.Fatal(s)
   }
}

func (p parameters) encode() string {
   return base64.StdEncoding.EncodeToString(p.Marshal())
}
