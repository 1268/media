package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/nbc"
   "fmt"
   "io"
   "net/http"
   "slices"
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
   mpd.BaseUrl = &dash.Url{resp.Request.URL}
   err = mpd.Unmarshal(data)
   if err != nil {
      return err
   }
   represents := slices.SortedFunc(mpd.Representation(),
      func(a, b dash.Representation) int {
         return a.Bandwidth - b.Bandwidth
      },
   )
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
