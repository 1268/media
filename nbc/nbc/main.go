package main

import (
   "flag"
   "mechanize.pages.dev"
)

type flags struct {
   guid int64
   mechanize.Stream
   resolution string
   bandwidth int64
}

func main() {
   var f flags
   flag.Int64Var(&f.guid, "b", 0, "GUID")
   flag.Int64Var(&f.bandwidth, "bandwidth", 8_000_000, "maximum bandwidth")
   flag.BoolVar(&f.Info, "i", false, "information")
   flag.IntVar(
      &http.Default_Client.Log_Level, "log",
      http.Default_Client.Log_Level, "log level",
   )
   flag.StringVar(&f.resolution, "r", "720", "resolution")
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
