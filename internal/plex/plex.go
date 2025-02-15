package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/plex"
   "errors"
   "fmt"
)

func (f *flags) download() error {
   var user plex.User
   err := user.New()
   if err != nil {
      return err
   }
   match, err := user.Match(f.address)
   if err != nil {
      return err
   }
   metadata, err := user.Metadata(match)
   if err != nil {
      return err
   }
   client, ok := metadata.Dash(user)
   if !ok {
      return errors.New("Metadata.Dash")
   }
   represents, err := internal.Mpd(client)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Widevine = client
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
