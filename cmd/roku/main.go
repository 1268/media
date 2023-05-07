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
   codec string
   id string
   mech.Stream
   lang string
   height int64
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // b
   flag.StringVar(&f.id, "b", "", "ID")
   // c
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "c", f.Client_ID, "client ID")
   // f
   flag.Int64Var(&f.height, "f", 1080, "video height")
   // g
   flag.StringVar(&f.codec, "g", "mp4a", "audio codec")
   // h
   flag.StringVar(&f.lang, "h", "en", "audio lang")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // k
   f.Private_Key = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&f.Private_Key, "k", f.Private_Key, "private key")
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
