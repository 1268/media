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
   var n next_data
   {
      s, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      sep := json.Split(`__NEXT_DATA__" type="application/json">`)
      if _, err := sep.After(s, &n); err != nil {
         return nil, err
      }
   }
   return &n, nil
}
