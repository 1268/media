package main

import (
   "154.pages.dev/http/option"
   "154.pages.dev/media"
   "154.pages.dev/media/roku"
   "flag"
   "os"
)

type flags struct {
   bandwidth int
   codec string
   height int
   id string
   lang string
   s media.Stream
   trace bool
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   flag.StringVar(&f.id, "b", "", "ID")
   flag.IntVar(&f.bandwidth, "bandwidth", 4_000_000, "maximum bandwidth")
   flag.StringVar(&f.codec, "c", "mp4a", "audio codec")
   flag.StringVar(
      &f.s.Client_ID, "client", home + "/widevine/client_id.bin", "client ID",
   )
   flag.IntVar(&f.height, "h", 1080, "maximum height")
   flag.BoolVar(&f.s.Info, "i", false, "information")
   flag.StringVar(
      &f.s.Private_Key, "key", home + "/widevine/private_key.pem", "private key",
   )
   flag.StringVar(&f.lang, "language", "en", "audio language")
   flag.BoolVar(&f.trace, "t", false, "trace")
   flag.Parse()
   if f.trace {
      option.Trace()
   } else {
      option.Verbose()
   }
   if f.id != "" {
      content, err := roku.New_Content(f.id)
      if err != nil {
         panic(err)
      }
      if err := f.DASH(content); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
