package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/internal"
   "41.neocities.org/media/plex"
   "errors"
   "fmt"
   "io"
   "net/http"
   "sort"
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
   reps, err := dash.Unmarshal(data, resp.Request.URL)
   if err != nil {
      return err
   }
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
   for _, rep := range reps {
      switch f.representation {
      case "":
         if _, ok := rep.Ext(); ok {
            fmt.Print(&rep, "\n\n")
         }
      case rep.Id:
         f.s.Namer = &plex.Namer{match}
         f.s.Wrapper = part
         return f.s.Download(rep)
      }
   }
   return nil
}

func get_forward() {
   for _, forward := range internal.Forward {
      fmt.Println(forward.Country, forward.Ip)
   }
}
