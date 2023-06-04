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
   content_ID string
   dash_cenc bool
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
   // b
   flag.StringVar(&f.content_ID, "b", "", "content ID")
   // client
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "client", f.Client_ID, "client ID")
   // d
   flag.BoolVar(&f.dash_cenc, "d", false, "DASH_CENC")
   // h
   flag.Int64Var(&f.height, "h", 720, "maximum height")
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
