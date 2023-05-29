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
   {
      reps := dash.Filter(reps, dash.Video)
      var index int
      for index < len(reps) {
         if reps[index].Height == f.height {
            break
         }
         index++
      }
      err := f.DASH_Get(reps, index)
      if err != nil {
         return err
      }
   }
   reps = dash.Filter(reps, func(r dash.Representation) bool {
      if !dash.Audio(r) {
         return false
      }
      if !strings.HasPrefix(r.Adaptation.Lang, f.lang) {
         return false
      }
      return true
   })
   index := dash.Index_Func(reps, func(r dash.Representation) bool {
      return strings.Contains(r.Codecs, f.codec)
   })
   return f.DASH_Get(reps, index)
}
