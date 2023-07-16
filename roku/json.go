package roku

import (
   "154.pages.dev/encoding/json"
   "io"
   "net/http"
)

func New_Cross_Site() (*Cross_Site, error) {
   // this has smaller body than www.roku.com
   res, err := http.Get("https://therokuchannel.roku.com")
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var site Cross_Site
   for _, cook := range res.Cookies() {
      if cook.Name == "_csrf" {
         site.cookie = cook
      }
   }
   {
      s, err := io.ReadAll(res.Body)
      if err != nil {
         return nil, err
      }
      sep := []byte("\tcsrf:")
      if err := json.Cut(s, sep, &site.token); err != nil {
         return nil, err
      }
   }
   return &site, nil
}
