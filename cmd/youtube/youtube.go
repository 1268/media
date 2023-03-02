package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/youtube"
   "fmt"
   "os"
)

func (f flags) player() (*youtube.Player, error) {
   var (
      req youtube.Request
      token *youtube.Token
   )
   switch f.request {
   case 0:
      req = youtube.Android()
   case 1:
      req = youtube.Android_Embed()
   case 2:
      req = youtube.Android_Check()
      home, err := os.UserHomeDir()
      if err != nil {
         return nil, err
      }
      token, err = youtube.Open_Token(home + "/mech/youtube.json")
      if err != nil {
         return nil, err
      }
      if err := token.Refresh(); err != nil {
         return nil, err
      }
   }
   return req.Player(f.video_ID, token)
}

func download(form *youtube.Format, name string) error {
   ext, err := form.Ext()
   if err != nil {
      return err
   }
   name = mech.Clean(name)
   file, err := os.Create(name + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   return form.Encode(file)
}

func (f flags) download() error {
   play, err := f.player()
   if err != nil {
      return err
   }
   forms := play.Streaming_Data.Adaptive_Formats
   if f.info {
      text, err := play.MarshalText()
      if err != nil {
         return err
      }
      os.Stdout.Write(text)
   } else {
      fmt.Println(play.Playability_Status)
      if f.audio != "" {
         form, ok := forms.Audio(f.audio)
         if ok {
            err := download(form, play.Name())
            if err != nil {
               return err
            }
         }
      }
      if f.height >= 1 {
         form, ok := forms.Video(f.height)
         if ok {
            err := download(form, play.Name())
            if err != nil {
               return err
            }
         }
      }
   }
   return nil
}

func refresh() error {
   code, err := youtube.New_Code()
   if err != nil {
      return err
   }
   fmt.Println(code)
   fmt.Scanln()
   token, err := code.Token()
   if err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return token.Create(home + "/mech/youtube.json")
}
