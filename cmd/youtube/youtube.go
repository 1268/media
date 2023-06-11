package main

import (
   "2a.pages.dev/mech/youtube"
   "2a.pages.dev/rosso/slices"
   "fmt"
   "os"
   "strings"
)

func (f flags) player() (*youtube.Player, error) {
   var token *youtube.Token
   switch f.request {
   case 0:
      f.r.Android()
   case 1:
      f.r.Android_Embed()
   case 2:
      f.r.Android_Check()
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
   return f.r.Player(token)
}

func (f flags) download() error {
   play, err := f.player()
   if err != nil {
      return err
   }
   forms := play.Streaming_Data.Adaptive_Formats
   slices.Sort(forms, func(a, b youtube.Format) bool {
      return b.Bitrate < a.Bitrate
   })
   if f.info {
      for i, form := range forms {
         if i >= 1 {
            fmt.Println()
         }
         fmt.Println(form)
      }
   } else {
      fmt.Printf("%+v\n", play.Playability_Status)
      // video
      index := slices.Index(forms, func(a youtube.Format) bool {
         // 1080p60
         if strings.HasPrefix(a.Quality_Label, f.video_q) {
            return strings.Contains(a.MIME_Type, f.video_t)
         }
         return false
      })
      err := f.encode(forms[index], play.Name())
      if err != nil {
         return err
      }
      // audio
      index = slices.Index(forms, func(a youtube.Format) bool {
         if a.Audio_Quality == f.audio_q {
            return strings.Contains(a.MIME_Type, f.audio_t)
         }
         return false
      })
      err = f.encode(forms[index], play.Name())
      if err != nil {
         return err
      }
   }
   return nil
}

func (f flags) encode(form youtube.Format, name string) error {
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
