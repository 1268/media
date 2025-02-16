package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/nbc"
   "fmt"
)

func (f *flags) download() error {
   var meta nbc.Metadata
   err := meta.New(f.nbc)
   if err != nil {
      return err
   }
   vod, err := meta.Vod()
   if err != nil {
      return err
   }
   represents, err := internal.Mpd(vod)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         var client nbc.Client
         client.New()
         f.s.Client = &client
         return f.s.Download(&represent)
      }
   }
   return nil
}
