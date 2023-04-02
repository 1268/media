package main

import (
   "2a.pages.dev/mech/paramount"
   "2a.pages.dev/rosso/dash"
   "fmt"
   "strings"
)

func (f flags) downloadable(preview *paramount.Preview) error {
   fmt.Println(paramount.Downloadable(f.guid))
   fmt.Println(preview.Name())
   return nil
}

func (f flags) dash(preview *paramount.Preview) error {
   var err error
   f.Poster, err = paramount.New_Session(f.guid)
   if err != nil {
      return err
   }
   f.Name = preview.Name()
   reps, err := f.Stream.DASH(paramount.DASH_CENC(f.guid))
   if err != nil {
      return err
   }
   audio := reps.Filter(func(r dash.Representation) bool {
      if r.MIME_Type != "audio/mp4" {
         return false
      }
      if r.Role() == "description" {
         return false
      }
      return true
   })
   index := audio.Index(func(a, b dash.Representation) bool {
      if !strings.HasPrefix(b.Adaptation.Lang, f.lang) {
         return false
      }
      if !strings.HasPrefix(b.Codecs, f.codecs) {
         return false
      }
      return true
   })
   if err := f.DASH_Get(audio, index); err != nil {
      return err
   }
   video := reps.Video()
   index = video.Index(func(a, b dash.Representation) bool {
      return b.Height == f.height
   })
   return f.DASH_Get(video, index)
}
