package main

import (
   "154.pages.dev/media/paramount"
   "154.pages.dev/stream"
   "154.pages.dev/stream/dash"
   "fmt"
   "io"
   "net/http"
   "os"
   "slices"
   "strings"
   option "154.pages.dev/http"
)

func (f flags) dash(token *paramount.App_Token) error {
   ref, err := paramount.DASH_CENC(f.content_ID)
   if err != nil {
      return err
   }
   res, err := http.Get(ref)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   f.s.Base = res.Request.URL
   reps, err := dash.Representations(res.Body)
   if err != nil {
      return err
   }
   if !f.s.Info {
      item, err := token.Item(f.content_ID)
      if err != nil {
         return err
      }
      f.s.Name, err = stream.Format_Film(item)
      if err != nil {
         return err
      }
      f.s.Poster, err = token.Session(f.content_ID)
      if err != nil {
         return err
      }
   }
   // video
   {
      reps := slices.DeleteFunc(slices.Clone(reps), dash.Not(dash.Video))
      slices.SortFunc(reps, func(a, b dash.Representation) int {
         return b.Bandwidth - a.Bandwidth
      })
      index := slices.IndexFunc(reps, func(a dash.Representation) bool {
         if a.Height <= f.height {
            return a.Bandwidth <= f.bandwidth
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
   index := slices.IndexFunc(reps, func(a dash.Representation) bool {
      if strings.HasPrefix(a.Adaptation.Lang, f.lang) {
         return strings.HasPrefix(a.Codecs, f.codec)
      }
      return false
   })
   return f.s.DASH_Get(reps, index)
}

func (f flags) downloadable(token *paramount.App_Token) error {
   item, err := token.Item(f.content_ID)
   if err != nil {
      return err
   }
   ref, err := paramount.Downloadable(f.content_ID)
   if err != nil {
      return err
   }
   if f.s.Info {
      fmt.Println(ref)
      return nil
   }
   name, err := stream.Format_Film(item)
   if err != nil {
      return err
   }
   res, err := http.Get(ref)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(name + ".mp4")
   if err != nil {
      return err
   }
   defer file.Close()
   pro := option.Progress_Length(res.ContentLength)
   if _, err := io.Copy(file, pro.Reader(res)); err != nil {
      return err
   }
   return nil
}

