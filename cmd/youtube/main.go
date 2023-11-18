package main

import (
   "154.pages.dev/http"
   "154.pages.dev/media/youtube"
   "flag"
   "strings"
)

type flags struct {
   audio_q string
   audio_t string
   info bool
   r youtube.Request
   refresh bool
   request int
   video_q string
   video_t string
   trace bool
}

func main() {
   var f flags
   flag.Var(&f.r, "a", "address")
   flag.StringVar(&f.r.Video_ID, "b", "", "video ID")
   flag.StringVar(&f.audio_q, "aq", "AUDIO_QUALITY_MEDIUM", "audio quality")
   flag.StringVar(&f.audio_t, "ac", "opus", "audio codec")
   flag.BoolVar(&f.info, "i", false, "information")
   flag.IntVar(&f.request, "r", 0, func() string {
      var b strings.Builder
      b.WriteString("0: Android\n")
      b.WriteString("1: Android embed\n")
      b.WriteString("2: Android check")
      return b.String()
   }())
   flag.BoolVar(&f.refresh, "refresh", false, "create OAuth refresh token")
   flag.StringVar(&f.video_q, "vq", "1080p", "video quality")
   flag.StringVar(&f.video_t, "vc", "vp9", "video codec")
   flag.BoolVar(&f.trace, "t", false, "trace")
   flag.Parse()
   http.No_Location()
   if f.trace {
      http.Trace()
   } else {
      http.Verbose()
   }
   if f.refresh {
      err := f.do_refresh()
      if err != nil {
         panic(err)
      }
   } else if f.r.Video_ID != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
