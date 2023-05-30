package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/roku"
   "2a.pages.dev/rosso/http"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   bandwidth int64
   codec string
   height int64
   id string
   lang string
   mech.Stream
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
   flag.Int64Var(&f.bandwidth, "bandwidth", 4_000_000, "maximum bandwidth")
   // c
   flag.StringVar(&f.codec, "c", "mp4a", "audio codec")
   // client
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "client", f.Client_ID, "client ID")
   // h
   flag.Int64Var(&f.height, "h", 1080, "maximum height")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // key
   f.Private_Key = filepath.Join(home, "mech/private_key.pem")
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
