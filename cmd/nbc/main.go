package main

import (
   "2a.pages.dev/mech/nbc"
   "flag"
)

func main() {
   var f flags
   flag.Int64Var(&f.guid, "b", 0, "GUID")
   flag.Int64Var(&f.bandwidth, "f", 3_000_000, "target bandwidth")
   flag.BoolVar(&f.Info, "i", false, "information")
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      nbc.Client.Log_Level = 2
   }
   if f.guid >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
