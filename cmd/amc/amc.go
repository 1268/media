package main

import (
   "2a.pages.dev/mech/amc"
   "2a.pages.dev/rosso/dash"
   "os"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   auth, err := amc.Read_Auth(home + "/mech/amc.json")
   if err != nil {
      return err
   }
   if err := auth.Refresh(); err != nil {
      return err
   }
   if err := auth.Write_File(home + "/mech/amc.json"); err != nil {
      return err
   }
   if !f.Info {
      content, err := auth.Content(f.address)
      if err != nil {
         return err
      }
      f.Namer, err = content.Video()
      if err != nil {
         return err
      }
   }
   play, err := auth.Playback(f.address)
   if err != nil {
      return err
   }
   f.Poster = play
   reps, err := f.DASH(play.HTTP_DASH().Src)
   if err != nil {
      return err
   }
   {
      reps := dash.Filter(reps, dash.Video)
      index := dash.Index_Func(reps, func(r dash.Representation) bool {
         return r.Height >= f.height
      })
      err := f.DASH_Get(reps, index)
      if err != nil {
         return err
      }
   }
   return f.DASH_Get(dash.Filter(reps, dash.Audio), 0)
}

func (f flags) login() error {
   auth, err := amc.Unauth()
   if err != nil {
      return err
   }
   if err := auth.Login(f.email, f.password); err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return auth.Write_File(home + "/mech/amc.json")
}
