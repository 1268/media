package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/paramount"
   "2a.pages.dev/rosso/http"
   "flag"
   "os"
)

type flags struct {
   mech.Stream
   bandwidth int
   codec string
   content_ID string
   dash_cenc bool
   height int
   lang string
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // b
   flag.StringVar(&f.content_ID, "b", "", "content ID")
   // bandwidth
   flag.IntVar(&f.bandwidth, "bandwidth", 5_000_000, "maximum bandwidth")
   // c
   flag.StringVar(&f.codec, "c", "mp4a", "audio codec")
   // client
   f.Client_ID = home + "/client_id.bin"
   flag.StringVar(&f.Client_ID, "client", f.Client_ID, "client ID")
   // d
   flag.BoolVar(&f.dash_cenc, "d", false, "DASH_CENC")
   // h
   flag.IntVar(&f.height, "h", 720, "maximum height")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // key
   f.Private_Key = home + "/private_key.pem"
   flag.StringVar(&f.Private_Key, "key", f.Private_Key, "private key")
   // language
   flag.StringVar(&f.lang, "language", "en", "audio language")
   // log
   flag.IntVar(
      &http.Default_Client.Log_Level, "log",
      http.Default_Client.Log_Level, "log level",
   )
   flag.Parse()
   if f.content_ID != "" {
      token, err := paramount.New_App_Token()
      if err != nil {
         panic(err)
      }
      if f.dash_cenc {
         err = f.dash(token)
      } else {
         err = f.downloadable(token)
      }
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
