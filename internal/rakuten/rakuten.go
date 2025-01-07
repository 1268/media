package main

import (
   "41.neocities.org/dash"
   "encoding/xml"
   "fmt"
   "io"
   "net/http"
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
   xml.Unmarshal(data, &mpd)
   for represent := range mpd.Representation() {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Namer, err = f.address.Movie()
         if err != nil {
            return err
         }
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
