package mech

import (
   "fmt"
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
   b := new(strings.Builder)
   title := Clean(n.Title())
   if season, err := n.Season(); err != nil {
      date, err := n.Date()
      if err != nil {
         return "", err
      }
      fmt.Fprint(b, title)
      fmt.Fprint(b, " - ")
      fmt.Fprint(b, date.Year())
   } else {
      fmt.Fprint(b, n.Series())
      fmt.Fprint(b, " - S")
      fmt.Fprint(b, season)
      fmt.Fprint(b, " E")
      episode, err := n.Episode()
      if err != nil {
         return "", err
      }
      fmt.Fprint(b, episode)
      fmt.Fprint(b, " - ")
      fmt.Fprint(b, title)
   }
   return b.String(), nil
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
