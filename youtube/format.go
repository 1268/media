package youtube

import (
   "fmt"
   "io"
   "mime"
   "strconv.pages.dev"
)

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
