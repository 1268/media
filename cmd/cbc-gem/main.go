package main

import (
   "2a.pages.dev/rosso/http"
   "flag"
   "mechanize.pages.dev"
)

type flags struct {
   address string
   bandwidth int64
   email string
   mechanize.Stream
   name string
   password string
   resolution string
}

func main() {
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.Int64Var(&f.bandwidth, "b", 3_000_000, "maximum bandwidth")
   flag.StringVar(&f.email, "e", "", "email")
   flag.BoolVar(&f.Info, "i", false, "information")
   flag.IntVar(
      &http.Default_Client.Log_Level, "log",
      http.Default_Client.Log_Level, "log level",
   )
   flag.StringVar(&f.name, "n", "English", "audio name")
   flag.StringVar(&f.password, "p", "", "password")
   flag.StringVar(&f.resolution, "r", "720", "resolution")
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
