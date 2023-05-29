package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/rosso/http"
   "flag"
)

type flags struct {
   address string
   bandwidth int64
   email string
   mech.Stream
   name string
   password string
}

func main() {
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.StringVar(&f.email, "e", "", "email")
   flag.Int64Var(&f.bandwidth, "f", 3687532, "video bandwidth")
   flag.StringVar(&f.name, "g", "English", "audio name")
   flag.BoolVar(&f.Info, "i", false, "information")
   flag.IntVar(
      &http.Default_Client.Log_Level, "log",
      http.Default_Client.Log_Level, "log level",
   )
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   if f.email != "" {
      err := f.profile()
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
