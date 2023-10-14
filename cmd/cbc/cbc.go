package main

import (
   "154.pages.dev/media/cbc"
   "154.pages.dev/stream"
   "154.pages.dev/stream/hls"
   "os"
   "slices"
   "strings"
)

func (f *flags) master() (*hls.Master, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   profile, err := gem.Read_Profile(home + "/cbc/profile.json")
   if err != nil {
      return nil, err
   }
   gem, err := gem.New_Catalog_Gem(f.address)
   if err != nil {
      return nil, err
   }
   media, err := profile.Media(gem.Item())
   if err != nil {
      return nil, err
   }
   f.s.Name, err = stream.Format_Film(gem.Structured_Metadata)
   if err != nil {
      return nil, err
   }
   return f.s.HLS(media.URL)
}

func (f flags) download() error {
   master, err := f.master()
   if err != nil {
      return err
   }
   // video
   master.Stream = slices.DeleteFunc(master.Stream, func(a hls.Stream) bool {
      return a.Resolution == ""
   })
   slices.SortFunc(master.Stream, func(a, b hls.Stream) int {
      return int(b.Bandwidth - a.Bandwidth)
   })
   index := slices.IndexFunc(master.Stream, func(a hls.Stream) bool {
      if strings.HasSuffix(a.Resolution, f.resolution) {
         return a.Bandwidth <= f.bandwidth
      }
      return false
   })
   if err := f.s.HLS_Streams(master.Stream, index); err != nil {
      return err
   }
   // audio
   master.Media = slices.DeleteFunc(master.Media, func(a hls.Media) bool {
      return a.Type != "AUDIO"
   })
   index = slices.IndexFunc(master.Media, func(a hls.Media) bool {
      return a.Name == f.name
   })
   return f.s.HLS_Media(master.Media, index)
}

func (f flags) profile() error {
   login, err := gem.New_Token(f.email, f.password)
   if err != nil {
      return err
   }
   profile, err := login.Profile()
   if err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return profile.Write_File(home + "/cbc/profile.json")
}
