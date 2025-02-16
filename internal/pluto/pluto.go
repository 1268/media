package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/pluto"
   "errors"
   "fmt"
)

func (f *flags) download() error {
   video, err := f.address.Vod(f.set_forward)
   if err != nil {
      return err
   }
   clips, err := video.Clips()
   if err != nil {
      return err
   }
   file, ok := clips.Dash()
   if !ok {
      return errors.New("Clips.Dash")
   }
   represents, err := internal.Mpd(file)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Client = pluto.Client{}
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
