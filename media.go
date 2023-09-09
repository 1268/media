package media

import (
   "encoding/json"
   "os"
   "strconv"
   "time"
)

type Namer interface {
   Series() string
   Season() (int64, error)
   Episode() (int64, error)
   Title() string
   Date() (time.Time, error)
}

func Name(n Namer) (string, error) {
   var b []byte
   title := Clean(n.Title())
   if series := n.Series(); series != "" {
      b = append(b, series...)
      b = append(b, " - S"...)
      {
         s, err := n.Season()
         if err != nil {
            return "", err
         }
         b = strconv.AppendInt(b, s, 10)
      }
      b = append(b, " E"...)
      {
         e, err := n.Episode()
         if err != nil {
            return "", err
         }
         b = strconv.AppendInt(b, e, 10)
      }
      b = append(b, " - "...)
      b = append(b, title...)
   } else {
      b = append(b, title...)
      b = append(b, " - "...)
      {
         d, err := n.Date()
         if err != nil {
            return "", err
         }
         b = d.AppendFormat(b, "2006")
      }
   }
   return string(b), nil
}

func Format(n Namer) (string, error) {
   var b []byte
   b = append(b, "series: "...)
   b = append(b, n.Series()...)
   b = append(b, "\nseason: "...)
   {
      s, err := n.Season()
      if err != nil {
         return "", err
      }
      b = strconv.AppendInt(b, s, 10)
   }
   b = append(b, "\nepisode: "...)
   {
      e, err := n.Episode()
      if err != nil {
         return "", err
      }
      b = strconv.AppendInt(b, e, 10)
   }
   b = append(b, "\ntitle: "...)
   b = append(b, n.Title()...)
   b = append(b, "\ndate: "...)
   {
      d, err := n.Date()
      if err != nil {
         return "", err
      }
      b = append(b, d.String()...)
   }
   return string(b), nil
}

func Clean(path string) string {
   m := map[byte]bool{
      '"': true,
      '*': true,
      '/': true,
      ':': true,
      '<': true,
      '>': true,
      '?': true,
      '\\': true,
      '|': true,
   }
   b := []byte(path)
   for k, v := range b {
      if m[v] {
         b[k] = '-'
      }
   }
   return string(b)
}

func User(name string) (map[string]string, error) {
   b, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   var m map[string]string
   if err := json.Unmarshal(b, &m); err != nil {
      return nil, err
   }
   return m, nil
}

