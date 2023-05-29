package main

import (
   "2a.pages.dev/mech/cbc"
   "2a.pages.dev/rosso/hls"
   "os"
)

func (f flags) download() error {
   master, err := f.master()
   if err != nil {
      return err
   }
   // video
   index := master.Streams.Index(func(s hls.Stream) bool {
      return s.Bandwidth >= f.bandwidth
   })
   if err := f.HLS_Streams(master.Streams, index); err != nil {
      return err
   }
   // audio
   index = master.Media.Index(func(m hls.Medium) bool {
      return m.Name == f.name
   })
   return f.HLS_Media(master.Media, index)
}

func (f flags) profile() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   login, err := cbc.New_Token(f.email, f.password)
   if err != nil {
      return err
   }
   profile, err := login.Profile()
   if err != nil {
      return err
   }
   return profile.Write_File(home + "/mech/cbc.json")
}
func (f *flags) master() (*hls.Master, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   profile, err := cbc.Read_Profile(home + "/mech/cbc.json")
   if err != nil {
      return nil, err
   }
   gem, err := cbc.New_Catalog_Gem(f.address)
   if err != nil {
      return nil, err
   }
   media, err := profile.Media(gem.Item())
   if err != nil {
      return nil, err
   }
   f.Namer = gem.Structured_Metadata
   return f.HLS(media.URL)
}
