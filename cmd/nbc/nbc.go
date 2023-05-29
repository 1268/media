package main

import (
   "2a.pages.dev/mech/nbc"
   "2a.pages.dev/rosso/hls"
   "strings"
)

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
   master.Streams.Sort(func(a, b hls.Stream) bool {
      return a.Bandwidth < b.Bandwidth
   })
   index := master.Streams.Index(func(s hls.Stream) bool{
      if strings.HasSuffix(s.Resolution, f.resolution) {
         return s.Bandwidth >= f.bandwidth
      }
      return false
   })
   return f.HLS_Streams(master.Streams, index)
}
