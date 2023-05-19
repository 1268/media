package mech

import (
   "strconv"
   "strings"
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
   if series := n.Series(); series != "" {
      b = append(b, series...)
      b = append(b, " - S"...)
      season, err := n.Season()
      if err != nil {
         return "", err
      }
      b = strconv.AppendInt(b, season, 10)
      b = append(b, " E"...)
      episode, err := n.Episode()
      if err != nil {
         return "", err
      }
      b = strconv.AppendInt(b, episode, 10)
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
