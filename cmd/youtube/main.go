package main

import (
   "2a.pages.dev/mech/youtube"
   "flag"
   "strings"
)

type flags struct {
   audio string
   height int
   info bool
   refresh bool
   request int
   verbose bool
   video_ID string
}

func main() {
   var f flags
   // a
   flag.Func("a", "address", func(s string) error {
      return youtube.Video_ID(s, &f.video_ID)
   })
   // b
   flag.StringVar(&f.video_ID, "b", "", "video ID")
   // f
   flag.IntVar(&f.height, "f", 1080, "target video height")
   // g
   flag.StringVar(&f.audio, "g", "AUDIO_QUALITY_MEDIUM", "target audio")
   // i
   flag.BoolVar(&f.info, "i", false, "information")
   // r
   var buf strings.Builder
   buf.WriteString("0: Android\n")
   buf.WriteString("1: Android embed\n")
   buf.WriteString("2: Android check")
   flag.IntVar(&f.request, "r", 0, buf.String())
   // refresh
   flag.BoolVar(&f.refresh, "refresh", false, "create OAuth refresh token")
   // v
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      youtube.HTTP_Client.Log_Level = 2
   }
   if f.refresh {
      err := refresh()
      if err != nil {
         panic(err)
      }
   } else if f.video_ID != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
