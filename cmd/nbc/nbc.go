package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/nbc"
   "2a.pages.dev/rosso/hls"
)

type flags struct {
   bandwidth int64
   guid int64
   mech.Stream
}

func (f flags) download() error {
   meta, err := nbc.New_Metadata(f.guid)
   if err != nil {
      return err
   }
   f.Namer = meta
   video, err := meta.Video()
   if err != nil {
      return err
   }
   master, err := f.HLS(video.Manifest_Path)
   if err != nil {
      return err
   }
   index := master.Streams.Index(func(s hls.Stream) bool{
      return s.Bandwidth >= f.bandwidth
   })
   return f.HLS_Streams(master.Streams, index)
}
