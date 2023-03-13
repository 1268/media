package main

import (
   "2a.pages.dev/mech/bandcamp"
   "2a.pages.dev/rosso/http"
   "fmt"
   "io"
   "os"
   "time"
)

func (f flags) tralbum(tralb *bandcamp.Tralbum) error {
   for i, track := range tralb.Tracks {
      if f.info {
         if i >= 1 {
            fmt.Println()
         }
         fmt.Println(track)
      } else if track.Streaming_URL != nil {
         req, err := http.NewRequest("GET", track.Streaming_URL.MP3_128, nil)
         if err != nil {
            return err
         }
         res, err := bandcamp.Client.Redirect(nil).Do(req)
         if err != nil {
            return err
         }
         file, err := os.Create(track.Name())
         if err != nil {
            return err
         }
         pro := http.Progress_Bytes(file, res.ContentLength)
         if _, err := io.Copy(pro, res.Body); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
         if err := file.Close(); err != nil {
            return err
         }
         time.Sleep(f.sleep)
      }
   }
   return nil
}

func (f flags) band(param *bandcamp.Params) error {
   band, err := param.Band()
   if err != nil {
      return err
   }
   for _, item := range band.Discography {
      tralb, err := item.Tralbum()
      if err != nil {
         return err
      }
      if err := f.tralbum(tralb); err != nil {
         return err
      }
   }
   return nil
}
