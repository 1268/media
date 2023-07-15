package youtube

import (
   "154.pages.dev/encoding/protobuf"
   "154.pages.dev/strconv"
   "fmt"
   "io"
   "mime"
   "net/http"
)

func (f Format) Encode(w io.Writer) error {
   req, err := http.Get_Parse(f.URL)
   if err != nil {
      return err
   }
   val := req.URL.Query()
   if err != nil {
      return err
   }
   pro := http.Progress_Bytes(w, f.Content_Length)
   client := http.Default_Client
   client.CheckRedirect = nil
   client.Log_Level = 0
   var pos int64
   for pos < f.Content_Length {
      val.Set("range", fmt.Sprint(pos, "-", pos+chunk-1))
      req.URL.RawQuery = val.Encode()
      res, err := client.Do(req)
      if err != nil {
         return err
      }
      if _, err := io.Copy(pro, res.Body); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
      pos += chunk
   }
   return nil
}

var Upload_Date = map[string]protobuf.Varint{
   "Last hour": 1,
   "Today": 2,
   "This week": 3,
   "This month": 4,
   "This year": 5,
}

var Type = map[string]protobuf.Varint{
   "Video": 1,
   "Channel": 2,
   "Playlist": 3,
   "Movie": 4,
}

var Duration = map[string]protobuf.Varint{
   "Under 4 minutes": 1,
   "4 - 20 minutes": 3,
   "Over 20 minutes": 2,
}

var Sort_By = map[string]protobuf.Varint{
   "Relevance": 0,
   "Upload date": 2,
   "View count": 3,
   "Rating": 1,
}

var Features = map[string]protobuf.Number{
   "360°": 15,
   "3D": 7,
   "4K": 14,
   "Creative Commons": 6,
   "HD": 4,
   "HDR": 25,
   "Live": 8,
   "Location": 23,
   "Purchased": 9,
   "Subtitles/CC": 5,
   "VR180": 26,
}

type Params struct {
   protobuf.Message
}

func (p Params) Sort_By(value protobuf.Varint) {
   p.Message[1] = value
}

func (p Params) Upload_Date(value protobuf.Varint) {
   p.Get(2)[1] = value
}

func (p Params) Type(value protobuf.Varint) {
   p.Get(2)[2] = value
}

func (p Params) Duration(value protobuf.Varint) {
   p.Get(2)[3] = value
}

func (p Params) Features(num protobuf.Number) {
   p.Get(2)[num] = protobuf.Varint(1)
}

func New_Params() Params {
   var p Params
   p.Message = make(protobuf.Message)
   p.Message[2] = make(protobuf.Message)
   return p
}
type Format struct {
   Quality_Label string `json:"qualityLabel"`
   Audio_Quality string `json:"audioQuality"`
   Bitrate strconv.Rate
   Content_Length int64 `json:"contentLength,string"`
   MIME_Type string `json:"mimeType"`
   URL string
}

func (f Format) String() string {
   var b []byte
   b = append(b, "quality: "...)
   if f.Quality_Label != "" {
      b = append(b, f.Quality_Label...)
   } else {
      b = append(b, f.Audio_Quality...)
   }
   b = append(b, "\nbitrate: "...)
   b = fmt.Append(b, f.Bitrate)
   b = append(b, "\ntype: "...)
   b = append(b, f.MIME_Type...)
   return string(b)
}

const chunk = 10_000_000

func (f Format) Ext() (string, error) {
   media, _, err := mime.ParseMediaType(f.MIME_Type)
   if err != nil {
      return "", err
   }
   switch media {
   case "audio/mp4":
      return ".m4a", nil
   case "audio/webm":
      return ".weba", nil
   case "video/mp4":
      return ".m4v", nil
   case "video/webm":
      return ".webm", nil
   }
   return "", fmt.Errorf(f.MIME_Type)
}
