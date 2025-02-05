package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/kanopy"
   "errors"
   "fmt"
   "os"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/kanopy.txt")
   if err != nil {
      return err
   }
   var token kanopy.WebToken
   err = token.Unmarshal(data)
   if err != nil {
      return err
   }
   member, err := token.Membership()
   if err != nil {
      return err
   }
   plays, err := token.Plays(member, f.video_id)
   if err != nil {
      return err
   }
   manifest, ok := plays.Dash()
   if !ok {
      return errors.New("VideoPlays.Dash")
   }
   represents, err := internal.Mpd(manifest.Url)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Wrapper = kanopy.Wrapper{manifest, &token}
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   data, err := kanopy.WebToken{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/kanopy.txt", data, os.ModePerm)
}
