package youtube

import (
   "2a.pages.dev/rosso/http"
   "errors"
   "io"
   "mime"
   "strconv"
)

func (f Format) Encode(w io.Writer) error {
   req, err := http.NewRequest("GET", f.URL, nil)
   if err != nil {
      return err
   }
   val := req.URL.Query()
   if err != nil {
      return err
   }
   pro := http.Progress_Bytes(w, f.Content_Length)
   client := http.Default_Client
   client.Log_Level = 0
   var pos int64
   for pos < f.Content_Length {
      var b []byte
      b = strconv.AppendInt(b, pos, 10)
      b = append(b, '-')
      b = strconv.AppendInt(b, pos+chunk-1, 10)
      val.Set("range", string(b))
      req.URL.RawQuery = val.Encode()
      res, err := client.Level(0).Do(req)
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
   return "", errors.New(f.MIME_Type)
}

type Formats []Format

func (f Formats) Audio(quality string) (*Format, bool) {
   for _, form := range f {
      if form.Audio_Quality == quality {
         return &form, true
      }
   }
   return nil, false
}

func (f Formats) Video(height int) (*Format, bool) {
   distance := func(f *Format) int {
      if f.Height > height {
         return f.Height - height
      }
      return height - f.Height
   }
   var (
      ok bool
      output *Format
   )
   for i, input := range f {
      // since codecs are in this order avc1,vp9,av01,
      // do "<=" so we can get last one
      if output == nil || distance(&input) <= distance(output) {
         output = &f[i]
         ok = true
      }
   }
   return output, ok
}

type Format struct {
   Audio_Quality string `json:"audioQuality"`
   Bitrate int64
   Content_Length int64 `json:"contentLength,string"`
   Height int
   MIME_Type string `json:"mimeType"`
   Quality_Label string `json:"qualityLabel"`
   URL string
   Width int
}
func (f Format) String() string {
   var b []byte
   b = append(b, "quality: "...)
   if f.Quality_Label != "" {
      b = append(b, f.Quality_Label...)
   } else {
      b = append(b, f.Audio_Quality...)
   }
   b = append(b, "\n\tbitrate: "...)
   b = strconv.AppendInt(b, f.Bitrate, 10)
   if f.Content_Length >= 1 { // Tq92D6wQ1mg
      b = append(b, "\n\tContent-Length: "...)
      b = strconv.AppendInt(b, f.Content_Length, 10)
   }
   b = append(b, "\n\tMIME type: "...)
   b = append(b, f.MIME_Type...)
   return string(b)
}

