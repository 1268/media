package mech

import (
   "strconv"
   "strings"
   "time"
)

type Namer interface {
   Series() string
   Season() int64
   Episode() int64
   Title() string
   Date() (time.Time, error)
}

func Name(n Namer) (string, error) {
   var b []byte
   if series := n.Series(); series != "" {
      b = append(b, series...)
      b = append(b, " - S"...)
      b = strconv.AppendInt(b, n.Season(), 10)
      b = append(b, " E"...)
      b = strconv.AppendInt(b, n.Episode(), 10)
      b = append(b, " - "...)
      b = append(b, n.Title()...)
   } else {
      date, err := n.Date()
      if err != nil {
         return "", err
      }
      b = append(b, n.Title()...)
      b = append(b, " - "...)
      b = append(b, strconv.Itoa(date.Year())...)
   }
   return string(b), nil
}

func Clean(path string) string {
   mapping := func(r rune) rune {
      if strings.ContainsRune(`"*/:<>?\|`, r) {
         return '-'
      }
      return r
   }
   return strings.Map(mapping, path)
}
