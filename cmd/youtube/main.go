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
   info bool
   refresh bool
   request int
   video_ID string
}

func main() {
   var f flags
   // a
   flag.Func("a", "address", func(s string) error {
      return youtube.Video_ID(s, &f.video_ID)
   })
   // audio
   flag.StringVar(&f.audio, "audio", "AUDIO_QUALITY_MEDIUM", "audio quality")
   // b
   flag.StringVar(&f.video_ID, "b", "", "video ID")
   // h
   flag.IntVar(&f.height, "h", 1080, "maximum height")
   // i
   flag.BoolVar(&f.info, "i", false, "information")
   // log
   flag.IntVar(
      &http.Default_Client.Log_Level, "log",
      http.Default_Client.Log_Level, "log level",
   )
   // r
   {
      var b strings.Builder
      b.WriteString("0: Android\n")
      b.WriteString("1: Android embed\n")
      b.WriteString("2: Android check")
      flag.IntVar(&f.request, "r", 0, b.String())
   }
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
