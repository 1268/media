package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/kanopy"
   "errors"
   "fmt"
   "os"
   "slices"
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
   data, err = manifest.Url.Get()
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
