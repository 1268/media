package main

import (
   "154.pages.dev/http/option"
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
   flag.StringVar(&f.audio_t, "at", "opus", "audio type")
   flag.BoolVar(&f.info, "i", false, "information")
   {
      var b strings.Builder
      b.WriteString("0: Android\n")
      b.WriteString("1: Android embed\n")
      b.WriteString("2: Android check")
      flag.IntVar(&f.request, "r", 0, b.String())
   }
   flag.BoolVar(&f.refresh, "refresh", false, "create OAuth refresh token")
   flag.StringVar(&f.video_q, "vq", "1080p", "video quality")
   flag.StringVar(&f.video_t, "vt", "vp9", "video type")
   flag.BoolVar(&f.trace, "t", false, "trace")
   flag.Parse()
   option.No_Location()
   if f.trace {
      option.Trace()
   } else {
      option.Verbose()
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
