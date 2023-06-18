package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/roku"
   "2a.pages.dev/rosso/http"
   "flag"
   "path/filepath"
)

type flags struct {
   bandwidth int
   codec string
   height int
   id string
   lang string
   mech.Stream
}

func main() {
   home, err := mech.Home()
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
   // client
   f.Client_ID = filepath.Join(home, "client_id.bin")
   flag.StringVar(&f.Client_ID, "client", f.Client_ID, "client ID")
   // h
   flag.IntVar(&f.height, "h", 1080, "maximum height")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // key
   f.Private_Key = filepath.Join(home, "private_key.pem")
   flag.StringVar(&f.Private_Key, "key", f.Private_Key, "private key")
   // language
   flag.StringVar(&f.lang, "language", "en", "audio language")
   // log
   flag.IntVar(
      &http.Default_Client.Log_Level, "log",
      http.Default_Client.Log_Level, "log level",
   )
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
