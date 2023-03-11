package main

import (
   "2a.pages.dev/mech/bandcamp"
   "flag"
   "time"
)

type flags struct {
   address string
   info bool
   sleep time.Duration
   verbose bool
}

func main() {
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.BoolVar(&f.info, "i", false, "information")
   flag.DurationVar(&f.sleep, "s", time.Second, "sleep")
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      bandcamp.Client.Log_Level = 2
   }
   if f.address != "" {
      param, err := bandcamp.New_Params(f.address)
      if err != nil {
         panic(err)
      }
      if param.I_Type != "" {
         tralb, err := param.Tralbum()
         if err != nil {
            panic(err)
         }
         if err := f.tralbum(tralb); err != nil {
            panic(err)
         }
      } else {
         err := f.band(param)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
