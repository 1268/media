package main

import (
   "41.neocities.org/media/criterion"
   "41.neocities.org/media/internal"
   "errors"
   "fmt"
   "os"
   "path"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/criterion.txt")
   if err != nil {
      return err
   }
   var token criterion.Token
   err = token.Unmarshal(data)
   if err != nil {
      return err
   }
   item, err := token.Video(path.Base(f.address))
   if err != nil {
      return err
   }
   files, err := token.Files(item)
   if err != nil {
      return err
   }
   file, ok := files.Dash()
   if !ok {
      return errors.New("VideoFiles.Dash")
   }
   represents, err := internal.Mpd(file)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Widevine = file
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   data, err := criterion.Token{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/criterion.txt", data, os.ModePerm)
}
