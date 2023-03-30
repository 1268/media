package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/nbc"
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
   video, err := meta.Video()
   if err != nil {
      return err
   }
   f.Name = meta.Name()
   master, err := f.HLS(video.Manifest_Path)
   if err != nil {
      return err
   }
   streams := master.Streams
   return f.HLS_Streams(streams, streams.Bandwidth(f.bandwidth))
}
