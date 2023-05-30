package main

import (
   "2a.pages.dev/mech/roku"
   "2a.pages.dev/rosso/dash"
   "strings"
)

func (f flags) DASH(content *roku.Content) error {
   if !f.Info {
      site, err := roku.New_Cross_Site()
      if err != nil {
         return err
      }
      f.Poster, err = site.Playback(f.id)
      if err != nil {
         return err
      }
      f.Namer = content
   }
   reps, err := f.Stream.DASH(content.DASH().URL)
   if err != nil {
      return err
   }
   // video
   {
      reps := reps.Filter(dash.Video)
      index := reps.Index(func(r dash.Represent) bool {
         if r.Bandwidth <= f.bandwidth {
            return r.Height <= f.height
         }
         return false
      })
      err := f.DASH_Get(reps, index)
      if err != nil {
         return err
      }
   }
   // audio
   reps = reps.Filter(dash.Audio)
   index := reps.Index(func(r dash.Represent) bool {
      return strings.Contains(r.Codecs, f.codec)
   })
   return f.DASH_Get(reps, index)
}
