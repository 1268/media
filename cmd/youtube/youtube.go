package main

import (
   "2a.pages.dev/mech/youtube"
   "fmt"
   "os"
)

func (f flags) encode(form *youtube.Format, name string) error {
   ext, err := form.Ext()
   if err != nil {
      return err
   }
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
      for _, form := range forms {
         fmt.Println(form)
      }
   } else {
      fmt.Printf("%+v\n", play.Playability_Status)
      if f.audio != "" {
         form, ok := forms.Audio(f.audio)
         if ok {
            err := f.encode(form, play.Name())
            if err != nil {
               return err
            }
         }
      }
      if f.height >= 1 {
         form, ok := forms.Video(f.height)
         if ok {
            err := f.encode(form, play.Name())
            if err != nil {
               return err
            }
         }
      }
   }
   return nil
}

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
      token, err = youtube.Read_Token(home + "/mech/youtube.json")
      if err != nil {
         return nil, err
      }
      if err := token.Refresh(); err != nil {
         return nil, err
      }
   }
   return req.Player(f.video_ID, token)
}

func (f flags) do_refresh() error {
   code, err := youtube.New_Device_Code()
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
   return token.Write_File(home + "/mech/youtube.json")
}

