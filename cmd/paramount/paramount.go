package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/paramount"
   "2a.pages.dev/rosso/dash"
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/slices"
   "fmt"
   "io"
   "os"
   "strings"
)

func (f flags) dash(token *paramount.App_Token) error {
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
   ref, err := paramount.DASH_CENC(f.content_ID)
   if err != nil {
      return err
   }
   reps, err := f.Stream.DASH(ref)
   if err != nil {
      return err
   }
   // video
   {
      reps := slices.Filter(reps, dash.Video)
      slices.Sort(reps, func(a, b dash.Represent) bool {
         return b.Bandwidth < a.Bandwidth
      })
      index := slices.Index(reps, func(a dash.Represent) bool {
         return a.Height <= f.height
      })
      err := f.DASH_Get(reps, index)
      if err != nil {
         return err
      }
   }
   // audio
   reps = slices.Filter(reps, dash.Audio)
   index := slices.Index(reps, func(a dash.Represent) bool {
      if strings.HasPrefix(a.Adaptation.Lang, f.lang) {
         return strings.HasPrefix(a.Codecs, f.codecs)
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
   name, err := mech.Name(item)
   if err != nil {
      return err
   }
   client := http.Default_Client
   client.CheckRedirect = nil
   res, err := client.Get(ref)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(name + ".mp4")
   if err != nil {
      return err
   }
   defer file.Close()
   pro := http.Progress_Bytes(file, res.ContentLength)
   if _, err := io.Copy(pro, res.Body); err != nil {
      return err
   }
   return nil
}
