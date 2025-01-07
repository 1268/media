package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/amc"
   "encoding/xml"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/amc.txt")
   if err != nil {
      return err
   }
   var auth amc.Authorization
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   data, err = auth.Refresh()
   if err != nil {
      return err
   }
   os.WriteFile(f.home + "/amc.txt", data, os.ModePerm)
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   play, err := auth.Playback(f.address.Nid)
   if err != nil {
      return err
   }
   wrap, ok := play.Dash()
   if !ok {
      return errors.New("Playback.Dash")
   }
   resp, err := http.Get(wrap.Source.Src)
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
         content, err := auth.Content(f.address.Path)
         if err != nil {
            return err
         }
         f.s.Namer, ok = content.Video()
         if !ok {
            return errors.New("ContentCompiler.Video")
         }
         f.s.Wrapper = wrap
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) login() error {
   var auth amc.Authorization
   err := auth.Unauth()
   if err != nil {
      return err
   }
   data, err := auth.Login(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/amc.txt", data, os.ModePerm)
}
