package main

import (
   "2a.pages.dev/stream/dash"
   "encoding.pages.dev/slices"
   "mechanize.pages.dev/roku"
   "net/http"
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
   res, err := http.Get(content.DASH().URL)
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
      reps := slices.Delete(slices.Clone(reps), dash.Not(dash.Video))
      index := slices.Index(reps, func(r dash.Representer) bool {
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
   index := slices.Index(reps, func(r dash.Representer) bool {
      return strings.Contains(r.Codecs, f.codec)
   })
   return f.DASH_Get(reps, index)
}
