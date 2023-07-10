package main

import (
   "2a.pages.dev/stream/hls"
   "encoding.pages.dev/slices"
   "mechanize.pages.dev/nbc"
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
   // video and audio
   slices.Sort(master.Stream, func(a, b hls.Stream) bool {
      return b.Bandwidth < a.Bandwidth
   })
   index := slices.Index(master.Stream, func(a hls.Stream) bool {
      if strings.HasSuffix(a.Resolution, f.resolution) {
         return a.Bandwidth <= f.bandwidth
      }
      return false
   })
   return f.HLS_Streams(master.Stream, index)
}
