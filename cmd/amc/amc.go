package main

import (
   "2a.pages.dev/mech/amc"
   "os"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   auth, err := amc.Open_Auth(home + "/mech/amc.json")
   if err != nil {
      return err
   }
   if err := auth.Refresh(); err != nil {
      return err
   }
   if err := auth.Create(home + "/mech/amc.json"); err != nil {
      return err
   }
   if !f.info {
      content, err := auth.Content(f.address)
      if err != nil {
         return err
      }
      video_play, err := content.Video_Player()
      if err != nil {
         return err
      }
      f.Name, err = video_play.Name()
      if err != nil {
         return err
      }
   }
   play, err := auth.Playback(f.address)
   if err != nil {
      return err
   }
   reps, err := f.DASH(play.Source().Src)
   if err != nil {
      return err
   }
   f.Poster = play
   if err := f.DASH_Get(reps.Audio(), 0); err != nil {
      return err
   }
   video := reps.Video()
   return f.DASH_Get(video, video.Bandwidth(f.bandwidth))
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
   return auth.Create(home + "/mech/amc.json")
}
