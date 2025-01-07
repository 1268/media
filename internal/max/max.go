package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/max"
   "encoding/xml"
   "fmt"
   "io"
   "net/http"
   "os"
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
   var mpd dash.Mpd
   xml.Unmarshal(data, &mpd)
   for represent := range mpd.Representation() {
      if *represent.MimeType == "video/mp4" {
         if *represent.Width < f.min_width {
            continue
         }
         if *represent.Width > f.max_width {
            continue
         }
      }
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Namer, err = login.Routes(&f.url)
         if err != nil {
            return err
         }
         f.s.Wrapper = play
         return f.s.Download(&represent)
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
