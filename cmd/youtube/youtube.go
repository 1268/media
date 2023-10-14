package main

import (
   "154.pages.dev/media/youtube"
   "154.pages.dev/stream"
   "fmt"
   "net/http"
   "os"
   "slices"
   "strings"
   option "154.pages.dev/http"
)

func (f flags) download() error {
   play, err := f.player()
   if err != nil {
      return err
   }
   forms := play.Streaming_Data.Adaptive_Formats
   slices.SortFunc(forms, func(a, b youtube.Format) int {
      return int(b.Bitrate - a.Bitrate)
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
      // need to do audio first, because URLs expire quickly
      index := slices.IndexFunc(forms, func(a youtube.Format) bool {
         if a.Audio_Quality == f.audio_q {
            return strings.Contains(a.MIME_Type, f.audio_t)
         }
         return false
      })
      err := encode(forms[index], stream.Format_Video(play))
      if err != nil {
         return err
      }
      // video
      index = slices.IndexFunc(forms, func(a youtube.Format) bool {
         // 1080p60
         if strings.HasPrefix(a.Quality_Label, f.video_q) {
            return strings.Contains(a.MIME_Type, f.video_t)
         }
         return false
      })
      return encode(forms[index], stream.Format_Video(play))
   }
   return nil
}

func encode(f youtube.Format, name string) error {
   file, err := func() (*os.File, error) {
      ext, err := f.Ext()
      if err != nil {
         return nil, err
      }
      return os.Create(name + ext)
   }()
   if err != nil {
      return err
   }
   defer file.Close()
   option.Silent()
   pro := option.Progress_Length(f.Content_Length)
   for _, byte_range := range f.Ranges() {
      err := func() error {
         res, err := http.Get(f.URL + byte_range)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         if _, err := file.ReadFrom(pro.Reader(res)); err != nil {
            return err
         }
         return nil
      }()
      if err != nil {
         return err
      }
   }
   return nil
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
   return token.Write_File(home + "/youtube.json")
}

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
      token, err = youtube.Read_Token(home + "/youtube.json")
      if err != nil {
         return nil, err
      }
      if err := token.Refresh(); err != nil {
         return nil, err
      }
   }
   return f.r.Player(token)
}

