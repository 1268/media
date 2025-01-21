package main

import (
   "41.neocities.org/dash"
   "fmt"
   "io"
   "net/http"
   "slices"
)

func (f *flags) download() error {
   fhd, err := f.address.Fhd().Info()
   if err != nil {
      return err
   }
   resp, err := http.Get(fhd.Url)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   var mpd dash.Mpd
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
         hd, err := f.address.Hd().Info()
         if err != nil {
            return err
         }
         fhd.LicenseUrl = hd.LicenseUrl
         f.s.Wrapper = fhd
         return f.s.Download(&represent)
      }
   }
   return nil
}
