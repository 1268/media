package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/nbc"
   "fmt"
   "net/http"
)

func (f *flags) download() error {
   var meta nbc.Metadata
   err := meta.New(f.nbc)
   if err != nil {
      return err
   }
   demand, err := meta.OnDemand()
   if err != nil {
      return err
   }
   resp, err := http.Get(demand.PlaybackUrl)
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
         var proxy nbc.DrmProxy
         proxy.New()
         f.s.Wrapper = &proxy
         return f.s.Download(&represent)
      }
   }
   return nil
}
