package bandcamp

import "strconv"

func (t Track) String() string {
   var b []byte
   b = append(b, "track num: "...)
   b = strconv.AppendInt(b, t.Track_Num, 10)
   b = append(b, "\ntitle: "...)
   b = append(b, t.Title...)
   b = append(b, "\nband: "...)
   b = append(b, t.Band_Name...)
   if t.Streaming_URL != nil {
      b = append(b, "\nURL: "...)
      b = append(b, t.Streaming_URL.MP3_128...)
   }
   return string(b)
}

type Track struct {
   Track_Num int64
   Title string
   Band_Name string
   Streaming_URL *struct {
      MP3_128 string `json:"mp3-128"`
   }
}
