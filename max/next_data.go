package max

import (
   "154.pages.dev/encoding/json"
   "io"
   "net/http"
)

func new_next_data(ref string) (*next_data, error) {
   res, err := http.Get(ref)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   sep := json.Split(`__NEXT_DATA__" type="application/json">`)
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   next := new(next_data)
   if _, err := sep.After(text, next); err != nil {
      return nil, err
   }
   return next, nil
}
