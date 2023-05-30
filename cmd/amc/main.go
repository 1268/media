package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/rosso/http"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   address string
   email string
   mech.Stream
   password string
   height int64
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // a
   flag.StringVar(&f.address, "a", "", "address")
   // client
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "client", f.Client_ID, "client ID")
   // e
   flag.StringVar(&f.email, "e", "", "email")
   // h
   flag.Int64Var(&f.height, "h", 1080, "maximum height")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // key
   f.Private_Key = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&f.Private_Key, "key", f.Private_Key, "private key")
   // log
   flag.IntVar(
      &http.Default_Client.Log_Level, "log",
      http.Default_Client.Log_Level, "log level",
   )
   // p
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   if f.email != "" {
      err := f.login()
      if err != nil {
         panic(err)
      }
   } else if f.address != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
