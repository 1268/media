package rakuten

import (
   "errors"
   "strings"
)

var classification_id = map[string]int{
   "cz": 272,
   "dk": 283,
   "fi": 284,
   "fr": 23,
   "ie": 41,
   "it": 36,
   "nl": 323,
   "no": 286,
   "pt": 64,
   "se": 282,
   "ua": 276,
   "uk": 18,
}

type address struct {
   market_code string
   movie string
   season string
   episode string
}

func (a *address) Set(data string) error {
   data = strings.TrimPrefix(data, "https://")
   data = strings.TrimPrefix(data, "www.")
   data = strings.TrimPrefix(data, "rakuten.tv")
   data = strings.TrimPrefix(data, "/")
   var found bool
   a.market_code, data, found = strings.Cut(data, "/")
   if !found {
      return errors.New("market code not found")
   }
   data, a.movie, found = strings.Cut(data, "movies/")
   if !found {
      data = strings.TrimPrefix(data, "player/episodes/stream/")
      a.season, a.episode, found = strings.Cut(data, "/")
      if !found {
         return errors.New("episode not found")
      }
   }
   return nil
}
