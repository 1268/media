package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/rosso/http"
   "flag"
)

type flags struct {
   guid int64
   mech.Stream
   resolution string
   bandwidth int64
}

func main() {
   var f flags
   flag.Int64Var(&f.guid, "b", 0, "GUID")
   flag.StringVar(&f.resolution, "f", "720", "resolution")
   flag.Int64Var(&f.bandwidth, "g", 9, "minimum bandwidth")
   flag.BoolVar(&f.Info, "i", false, "information")
   flag.IntVar(
      &http.Default_Client.Log_Level, "log",
      http.Default_Client.Log_Level, "log level",
   )
   flag.Parse()
   if f.guid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
