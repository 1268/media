package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/media/amc"
   "net/http"
   "os"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   auth, err := amc.Read_Auth(home + "/amcplus.com/auth.json")
   if err != nil {
      return err
   }
   if err := auth.Refresh(); err != nil {
      return err
   }
   if err := auth.Write_File(home + "/amc.json"); err != nil {
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
   res, err := http.Get(play.HTTP_DASH().Src)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   reps, err := dash.Representers(res.Body)
   if err != nil {
      return err
   }
   // video
   {
      reps := slices.Delete(slices.Clone(reps), dash.Audio)
      slices.Sort(reps, func(a, b dash.Representer) bool {
         return b.Height < a.Height
      })
      index := slices.Index(reps, func(a dash.Representer) bool {
         return a.Height <= f.height
      })
      err := f.DASH_Get(reps, index)
      if err != nil {
         return err
      }
   }
   // audio
   return f.DASH_Get(slices.Delete(reps, dash.Video), 0)
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
   return auth.Write_File(home + "/amcplus.com/auth.json")
}
