package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/media/amc"
   "golang.org/x/exp/slices"
   "net/http"
   "os"
)

func (f flags) download() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   var auth amc.Auth_ID
   {
      b, err := os.ReadFile(home + "/amc/auth.json")
      if err != nil {
         return err
      }
      auth.Unmarshal(b)
   }
   if err := auth.Refresh(); err != nil {
      return err
   }
   {
      b, err := auth.Marshal()
      if err != nil {
         return err
      }
      os.WriteFile(home + "/amc/auth.json", b, 0666)
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
      reps := slices.DeleteFunc(slices.Clone(reps), dash.Audio)
      slices.SortFunc(reps, func(a, b dash.Representer) bool {
         return b.Height < a.Height
      })
      index := slices.IndexFunc(reps, func(a dash.Representer) bool {
         return a.Height <= f.height
      })
      err := f.DASH_Get(reps, index)
      if err != nil {
         return err
      }
   }
   // audio
   return f.DASH_Get(slices.DeleteFunc(reps, dash.Video), 0)
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
   {
      b, err := auth.Marshal()
      if err != nil {
         return err
      }
      os.WriteFile(home + "/amc/auth.json", b, 0666)
   }
   return nil
}

