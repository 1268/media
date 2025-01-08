package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/internal"
   "41.neocities.org/media/plex"
   "errors"
   "fmt"
   "io"
   "net/http"
)

func (f *flags) download() error {
   var user plex.Anonymous
   err := user.New()
   if err != nil {
      return err
   }
   match, err := user.Match(&f.address)
   if err != nil {
      return err
   }
   video, err := user.Video(match, f.set_forward)
   if err != nil {
      return err
   }
   part, ok := video.Dash()
   if !ok {
      return errors.New("OnDemand.Dash")
   }
   var req http.Request
   req.URL = part.Key.Url
   if f.set_forward != "" {
      req.Header = http.Header{
         "x-forwarded-for": {f.set_forward},
      }
   }
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   var mpd dash.Mpd
   mpd.Unmarshal(data)
   for represent := range mpd.Representation() {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Namer = &plex.Namer{match}
         f.s.Wrapper = part
         return f.s.Download(&represent)
      }
   }
   return nil
}

func get_forward() {
   for _, forward := range internal.Forward {
      fmt.Println(forward.Country, forward.Ip)
   }
}
