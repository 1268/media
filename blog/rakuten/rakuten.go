package rakuten

import (
   "errors"
   "strings"
)

// - https://www.rakuten.tv/uk/player/episodes/stream/hell-s-kitchen-usa-15/hell-s-kitchen-usa-15-1
//    - hell-s-kitchen-usa-15-1
//    - /v3/seasons/hell-s-kitchen-usa-15
// - https://www.rakuten.tv/fr/movies/infidele
//    - infidele
//    - /v3/movies/infidele

type address struct {
   episode string
   market_code string
   movie string
   season string
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
