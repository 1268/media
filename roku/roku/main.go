package main

import (
   "flag"
   "mechanize.pages.dev"
   "mechanize.pages.dev/roku"
   "os"
)

type flags struct {
   bandwidth int
   codec string
   height int
   id string
   lang string
   mechanize.Stream
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // b
   flag.StringVar(&f.id, "b", "", "ID")
   // bandwidth
   flag.IntVar(&f.bandwidth, "bandwidth", 4_000_000, "maximum bandwidth")
   // c
   flag.StringVar(&f.codec, "c", "mp4a", "audio codec")
   // h
   flag.IntVar(&f.height, "h", 1080, "maximum height")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // language
   flag.StringVar(&f.lang, "language", "en", "audio language")
   // log
   flag.IntVar(
      &http.Default_Client.Log_Level, "log",
      http.Default_Client.Log_Level, "log level",
   )
   // client
   f.Client_ID = home + "/widevine/client_id.bin"
   flag.StringVar(&f.Client_ID, "client", f.Client_ID, "client ID")
   // key
   f.Private_Key = home + "/widevine/private_key.pem"
   flag.StringVar(&f.Private_Key, "key", f.Private_Key, "private key")
   flag.Parse()
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
