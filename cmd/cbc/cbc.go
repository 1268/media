package main

import (
   "2a.pages.dev/mech/cbc"
   "2a.pages.dev/rosso/hls"
   "os"
)

func (f *flags) master() (*hls.Master, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   profile, err := cbc.Read_Profile(home + "/mech/cbc.json")
   if err != nil {
      return nil, err
   }
   asset, err := cbc.New_Asset(f.id)
   if err != nil {
      return nil, err
   }
   media, err := profile.Media(asset)
   if err != nil {
      return nil, err
   }
   f.Namer = asset
   return f.HLS(*media.URL)
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

func (f flags) download() error {
   master, err := f.master()
   if err != nil {
      return err
   }
   // video
   streams := master.Streams.Filter(func(s hls.Stream) bool {
      return s.Resolution != ""
   })
   index := streams.Bandwidth(f.bandwidth)
   if err := f.HLS_Streams(streams, index); err != nil {
      return err
   }
   // audio
   media := master.Media.Filter(func(m hls.Medium) bool {
      return m.Type == "AUDIO"
   })
   index = media.Index(func(a, b hls.Medium) bool {
      return b.Name == f.name
   })
   return f.HLS_Media(media, index)
}
