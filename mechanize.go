package mechanize

import (
   "encoding/json"
   "fmt"
   "os"
   "strings"
   "time"
)

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

type Namer interface {
   Series() string
   Season() (int64, error)
   Episode() (int64, error)
   Title() string
   Date() (time.Time, error)
}
