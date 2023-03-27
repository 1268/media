package main

import (
   "2a.pages.dev/mech/youtube"
   "2a.pages.dev/rosso/http"
   "flag"
   "strings"
)

type flags struct {
   audio string
   height int
   http.Client
   info bool
   refresh bool
   request int
   video_ID string
}

func main() {
   f := flags{Client: http.Default_Client}
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
   // log
   flag.IntVar(&f.Log_Level, "log", f.Log_Level, "log level")
   // r
   var buf strings.Builder
   buf.WriteString("0: Android\n")
   buf.WriteString("1: Android embed\n")
   buf.WriteString("2: Android check")
   flag.IntVar(&f.request, "r", 0, buf.String())
   // refresh
   flag.BoolVar(&f.refresh, "refresh", false, "create OAuth refresh token")
   flag.Parse()
   if f.refresh {
      err := f.do_refresh()
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
