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
   f.Name = asset.Name()
   return f.HLS(*media.URL)
}

func (f flags) download() error {
   master, err := f.master()
   if err != nil {
      return err
   }
   media := master.Media.Filter(func(m hls.Medium) bool {
      return m.Type == "AUDIO"
   })
   index := media.Index(func(a, b hls.Medium) bool {
      return b.Name == f.name
   })
   if err := f.HLS_Media(media, index); err != nil {
      return err
   }
   streams := master.Streams.Filter(func(s hls.Stream) bool {
      return s.Resolution != ""
   })
   return f.HLS_Streams(streams, streams.Bandwidth(f.bandwidth))
}

func (f flags) profile() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   login, err := cbc.New_Login(f.email, f.password)
   if err != nil {
      return err
   }
   web, err := login.Web_Token()
   if err != nil {
      return err
   }
   top, err := web.Over_The_Top()
   if err != nil {
      return err
   }
   profile, err := top.Profile()
   if err != nil {
      return err
   }
   return profile.Write_File(home + "/mech/cbc.json")
}
