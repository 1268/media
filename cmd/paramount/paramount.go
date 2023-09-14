package main

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/http/option"
   "154.pages.dev/media"
   "154.pages.dev/media/paramount"
   "fmt"
   "io"
   "net/http"
   "os"
   "slices"
   "strings"
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
   f.Base = res.Request.URL
   reps, err := dash.Representations(res.Body)
   if err != nil {
      return err
   }
   if !f.Info {
      item, err := token.Item(f.content_ID)
      if err != nil {
         return err
      }
      f.Namer = item
      f.Poster, err = token.Session(f.content_ID)
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
      err := f.DASH_Get(reps, index)
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
   return f.DASH_Get(reps, index)
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
   if f.Info {
      fmt.Println(ref)
      return nil
   }
   name, err := media.Name(item)
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

