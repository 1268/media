package main

import (
   "154.pages.dev/media/nbc"
   "154.pages.dev/stream"
   "154.pages.dev/stream/hls"
   "slices"
   "strings"
)

func (f flags) download() error {
   meta, err := nbc.New_Metadata(f.guid)
   if err != nil {
      return err
   }
   on, err := meta.On_Demand()
   if err != nil {
      return err
   }
   f.s.Name, err = func() (string, error) {
      if meta.Episode_Number != nil {
         return stream.Format_Episode(meta)
      }
      return stream.Format_Film(meta)
   }()
   if err != nil {
      return err
   }
   master, err := f.s.HLS(on.Manifest_Path)
   if err != nil {
      return err
   }
   // video and audio
   slices.SortFunc(master.Stream, func(a, b hls.Stream) int {
      return int(b.Bandwidth - a.Bandwidth)
   })
   index := slices.IndexFunc(master.Stream, func(a hls.Stream) bool {
      if strings.HasSuffix(a.Resolution, f.resolution) {
         return a.Bandwidth <= f.bandwidth
      }
      return false
   })
   return f.s.HLS_Streams(master.Stream, index)
}
