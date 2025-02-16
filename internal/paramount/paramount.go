package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/paramount"
   "fmt"
)

func (f *flags) do_read() error {
   // item
   var token paramount.AppToken
   if f.intl {
      token = paramount.ComCbsCa
   } else {
      token = paramount.ComCbsApp
   }
   var item paramount.Item
   data, err := item.Marshal(&token, f.content_id)
   if err != nil {
      return err
   }
   err = item.Unmarshal(data)
   if err != nil {
      return err
   }
   // mpd
   represents, err := internal.Mpd(item)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         // INTL does NOT allow anonymous key request, so if you are INTL you
         // will need to use US VPN until someone codes the INTL login
         f.s.Client, err = paramount.ComCbsApp.Session(f.content_id)
         if err != nil {
            return err
         }
         return f.s.Download(&represent)
      }
   }
   return nil
}
