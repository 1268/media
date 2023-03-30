package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/paramount"
   "2a.pages.dev/rosso/http"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   asset_type int
   codecs string
   guid string
   height int64
   lang string
   mech.Stream
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // a
   flag.IntVar(&f.asset_type, "a", 0, `0 StreamPack
1 HLS_CLEAR
2 DASH_CENC`)
   // b
   flag.StringVar(&f.guid, "b", "", "GUID")
   // c
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "c", f.Client_ID, "client ID")
   // f
   flag.Int64Var(&f.height, "f", 720, "video height")
   // g
   flag.StringVar(&f.codecs, "g", "mp4a", "audio codec")
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
   if f.guid != "" {
      preview, err := paramount.New_Preview(f.guid)
      if err != nil {
         panic(err)
      }
      switch f.asset_type {
      case 0:
         err := f.stream_pack(preview)
         if err != nil {
            panic(err)
         }
      case 1:
         err := f.hls_clear(preview)
         if err != nil {
            panic(err)
         }
      case 2:
         err := f.dash(preview)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
