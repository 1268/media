package main

import (
   "154.pages.dev/encoding/hls"
   "154.pages.dev/media/nbc"
   "golang.org/x/exp/slices"
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
   slices.SortFunc(master.Stream, func(a, b hls.Stream) bool {
      return b.Bandwidth < a.Bandwidth
   })
   index := slices.IndexFunc(master.Stream, func(a hls.Stream) bool {
      if strings.HasSuffix(a.Resolution, f.resolution) {
         return a.Bandwidth <= f.bandwidth
      }
      return false
   })
   return f.HLS_Streams(master.Stream, index)
}
