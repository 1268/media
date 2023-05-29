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
      reps := reps.Filter(dash.Video)
      index := reps.Index(func(r dash.Represent) bool {
         return r.Height >= f.height
      })
      err := f.DASH_Get(reps, index)
      if err != nil {
         return err
      }
   }
   reps = reps.Filter(func(r dash.Represent) bool {
      if strings.HasPrefix(r.Adaptation.Lang, f.lang) {
         return dash.Audio(r)
      }
      return false
   })
   index := reps.Index(func(r dash.Represent) bool {
      return strings.Contains(r.Codecs, f.codec)
   })
   return f.DASH_Get(reps, index)
}
