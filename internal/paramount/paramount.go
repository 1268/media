package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/paramount"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
   "slices"
)

func (f *flags) do_write() error {
   os.Mkdir(f.content_id, os.ModePerm)
   // item
   var token paramount.AppToken
   if f.intl {
      token = paramount.ComCbsCa
   } else {
      token = paramount.ComCbsApp
   }
   var item paramount.VideoItem
   data, err := item.Marshal(&token, f.content_id)
   if err != nil {
      return err
   }
   err = os.WriteFile(f.content_id + "/item.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   // mpd
   err = item.Unmarshal(data)
   if err != nil {
      return err
   }
   resp, err := http.Get(item.Mpd())
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   err = os.WriteFile(f.content_id + "/body.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   // Request
   data, err = resp.Request.URL.MarshalBinary()
   if err != nil {
      return err
   }
   return os.WriteFile(f.content_id + "/request.txt", data, os.ModePerm)
}
func (f *flags) do_read() error {
   data, err := os.ReadFile(f.content_id + "/request.txt")
   if err != nil {
      return err
   }
   var mpd dash.Mpd
   mpd.BaseUrl = &dash.Url{}
   err = mpd.BaseUrl.UnmarshalText(data)
   if err != nil {
      return err
   }
   data, err = os.ReadFile(f.content_id + "/body.txt")
   if err != nil {
      return err
   }
   mpd.Unmarshal(data)
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
         // INTL does NOT allow anonymous key request, so if you are INTL you
         // will need to use US VPN until someone codes the INTL login
         f.s.Wrapper, err = paramount.ComCbsApp.Session(f.content_id)
         if err != nil {
            return err
         }
         data, err = os.ReadFile(f.content_id + "/item.txt")
         if err != nil {
            return err
         }
         var item paramount.VideoItem
         err = item.Unmarshal(data)
         if err != nil {
            return err
         }
         f.s.Namer = &item
         return f.s.Download(&represent)
      }
   }
   return nil
}
