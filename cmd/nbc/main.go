package main

import (
   "2a.pages.dev/rosso/http"
   "flag"
)

func main() {
   var f flags
   flag.Int64Var(&f.guid, "b", 0, "GUID")
   flag.Int64Var(&f.bandwidth, "f", 5429000, "target bandwidth")
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
