package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/pluto"
   "errors"
   "fmt"
   "net/http"
)

func (f *flags) download() error {
   video, err := f.address.Video(f.set_forward)
   if err != nil {
      return err
   }
   clip, err := video.Clip()
   if err != nil {
      return err
   }
   var (
      req http.Request
      ok  bool
   )
   req.URL, ok = clip.Dash()
   if !ok {
      return errors.New("EpisodeClip.Dash")
   }
   req.URL.Scheme = pluto.Base[0].Scheme
   req.URL.Host = pluto.Base[0].Host
   resp, err := http.DefaultClient.Do(&req)
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
         f.s.Wrapper = pluto.Wrapper{}
         return f.s.Download(&represent)
      }
   }
   return nil
}

func get_forward() {
   for _, forward := range internal.Forward {
      fmt.Println(forward.Country, forward.Ip)
   }
}
