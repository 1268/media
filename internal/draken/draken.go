package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/draken"
   "encoding/xml"
   "fmt"
   "io"
   "net/http"
   "os"
   "path"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/draken.txt")
   if err != nil {
      return err
   }
   var login draken.AuthLogin
   err = login.Unmarshal(data)
   if err != nil {
      return err
   }
   var movie draken.FullMovie
   err = movie.New(path.Base(f.address))
   if err != nil {
      return err
   }
   title, err := login.Entitlement(&movie)
   if err != nil {
      return err
   }
   play, err := login.Playback(&movie, title)
   if err != nil {
      return err
   }
   resp, err := http.Get(play.Playlist)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   var mpd dash.Mpd
   xml.Unmarshal(data, &mpd)
   for represent := range mpd.Representation() {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Namer = &draken.Namer{movie}
         f.s.Wrapper = &draken.Wrapper{&login, play}
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   data, err := draken.AuthLogin{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/draken.txt", data, os.ModePerm)
}
