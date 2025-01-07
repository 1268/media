package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/nbc"
   "encoding/xml"
   "fmt"
   "io"
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
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   var mpd dash.Mpd
   xml.Unmarshal(data, &mpd)
   for represent := range mpd.Representation() {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Namer = &meta
         var proxy nbc.DrmProxy
         proxy.New()
         f.s.Wrapper = &proxy
         return f.s.Download(&represent)
      }
   }
   return nil
}
