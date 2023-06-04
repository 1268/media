package main

import (
   "2a.pages.dev/mech/roku"
   "2a.pages.dev/rosso/dash"
   "2a.pages.dev/rosso/slices"
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
      reps := slices.Delete(slices.Clone(reps), dash.Not(dash.Video))
      index := slices.Index(reps, func(r dash.Represent) bool {
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
   reps = slices.Delete(reps, dash.Not(dash.Audio))
   index := slices.Index(reps, func(r dash.Represent) bool {
      return strings.Contains(r.Codecs, f.codec)
   })
   return f.DASH_Get(reps, index)
}
