package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/max"
   "fmt"
   "io"
   "net/http"
   "os"
   "sort"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/max.txt")
   if err != nil {
      return err
   }
   var login max.LinkLogin
   err = login.Unmarshal(data)
   if err != nil {
      return err
   }
   play, err := login.Playback(&f.url)
   if err != nil {
      return err
   }
   resp, err := http.Get(play.Fallback.Manifest.Url.String)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
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
      if rep.GetAdaptationSet().GetPeriod().Id != "0" {
         continue
      }
      if rep.MimeType == "video/mp4" {
         if rep.Width < f.min_width {
            continue
         }
         if rep.Width > f.max_width {
            continue
         }
      }
      switch f.representation {
      case "":
         if _, ok := rep.Ext(); ok {
            fmt.Print(&rep, "\n\n")
         }
      case rep.Id:
         f.s.Namer, err = login.Routes(&f.url)
         if err != nil {
            return err
         }
         f.s.Wrapper = play
         return f.s.Download(rep)
      }
   }
   return nil
}

func (f *flags) do_initiate() error {
   var token max.BoltToken
   err := token.New()
   if err != nil {
      return err
   }
   os.WriteFile("token.txt", []byte(token.St), os.ModePerm)
   initiate, err := token.Initiate()
   if err != nil {
      return err
   }
   fmt.Printf("%+v\n", initiate)
   return nil
}
func (f *flags) do_login() error {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      return err
   }
   var token max.BoltToken
   token.St = string(data)
   data, err = max.LinkLogin{}.Marshal(&token)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/max.txt", data, os.ModePerm)
}
