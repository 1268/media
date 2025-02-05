package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/itv"
   "errors"
   "fmt"
   "net/http"
   "net/http/cookiejar"
   "path"
)

func (f *flags) download() error {
   var id itv.LegacyId
   err := id.Set(path.Base(f.address))
   if err != nil {
      return err
   }
   discovery, err := id.Discovery()
   if err != nil {
      return err
   }
   play, err := discovery.Playlist()
   if err != nil {
      return err
   }
   file, ok := play.Resolution1080()
   if !ok {
      return errors.New("resolution 1080")
   }
   http.DefaultClient.Jar, err = cookiejar.New(nil)
   if err != nil {
      return err
   }
   resp, err := http.Get(file.Href.S)
   if err != nil {
      return err
   }
   represents, err := internal.Representation(resp)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Wrapper = file
         return f.s.Download(&represent)
      }
   }
   return nil
}
