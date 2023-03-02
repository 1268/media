package main

import (
   "2a.pages.dev/mech/paramount"
   "2a.pages.dev/rosso/dash"
   "2a.pages.dev/rosso/hls"
   "strings"
   "strconv"
)

func (f flags) dash(preview *paramount.Preview) error {
   var err error
   f.Poster, err = paramount.New_Session(f.guid)
   if err != nil {
      return err
   }
   f.Name = preview.Name()
   reps, err := f.Stream.DASH(paramount.DASH(f.guid))
   if err != nil {
      return err
   }
   audio := reps.Filter(func(r dash.Representation) bool {
      if r.MIME_Type != "audio/mp4" {
         return false
      }
      if r.Role() == "description" {
         return false
      }
      return true
   })
   index := audio.Index(func(a, b dash.Representation) bool {
      if !strings.HasPrefix(b.Adaptation.Lang, f.lang) {
         return false
      }
      if !strings.HasPrefix(b.Codecs, f.codecs) {
         return false
      }
      return true
   })
   if err := f.DASH_Get(audio, index); err != nil {
      return err
   }
   video := reps.Video()
   index = video.Index(func(a, b dash.Representation) bool {
      return b.Height == f.height
   })
   return f.DASH_Get(video, index)
}

func (f flags) stream_pack(preview *paramount.Preview) error {
   f.Name = preview.Name()
   master, err := f.Stream.HLS(paramount.Stream_Pack(f.guid))
   if err != nil {
      return err
   }
   streams := master.Streams.Filter(func(s hls.Stream) bool {
      return s.Resolution != ""
   })
   index := streams.Index(func(a, b hls.Stream) bool {
      return strings.HasSuffix(b.Resolution, strconv.FormatInt(f.height, 10))
   })
   return f.HLS_Streams(streams, index)
}

func (f flags) hls_clear(preview *paramount.Preview) error {
   f.Name = preview.Name()
   master, err := f.Stream.HLS(paramount.HLS_Clear(f.guid))
   if err != nil {
      return err
   }
   streams := master.Streams.Filter(func(s hls.Stream) bool {
      return s.Resolution != ""
   })
   index := streams.Index(func(a, b hls.Stream) bool {
      return strings.HasSuffix(b.Resolution, strconv.FormatInt(f.height, 10))
   })
   return f.HLS_Streams(streams, index)
}
