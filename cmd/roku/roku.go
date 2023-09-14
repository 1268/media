package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/media/roku"
   "errors"
   "net/http"
   "slices"
   "strings"
)

func (f flags) DASH(content *roku.Content) error {
   if !f.s.Info {
      site, err := roku.New_Cross_Site()
      if err != nil {
         return err
      }
      f.s.Poster, err = site.Playback(f.id)
      if err != nil {
         return err
      }
      f.s.Namer = content
   }
   res, err := http.Get(content.DASH().URL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   f.s.Base = res.Request.URL
   reps, err := dash.Representations(res.Body)
   if err != nil {
      return err
   }
   // video
   {
      reps := slices.DeleteFunc(slices.Clone(reps), dash.Not(dash.Video))
      index := slices.IndexFunc(reps, func(r dash.Representation) bool {
         if r.Bandwidth <= f.bandwidth {
            return r.Height <= f.height
         }
         return false
      })
      err := f.s.DASH_Get(reps, index)
      if err != nil {
         return err
      }
   }
   // audio
   reps = slices.DeleteFunc(reps, dash.Not(dash.Audio))
   index := slices.IndexFunc(reps, func(r dash.Representation) bool {
      return strings.Contains(r.Codecs, f.codec)
   })
   return f.s.DASH_Get(reps, index)
}
