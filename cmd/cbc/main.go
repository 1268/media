package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/cbc"
   "flag"
)

type flags struct {
   bandwidth int64
   email string
   id string
   mech.Stream
   name string
   password string
   verbose bool
}

func main() {
   var f flags
   flag.StringVar(&f.id, "b", "", "ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.Int64Var(&f.bandwidth, "f", 3687532, "video bandwidth")
   flag.StringVar(&f.name, "g", "English", "audio name")
   flag.BoolVar(&f.Info, "i", false, "information")
   flag.StringVar(&f.password, "p", "", "password")
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      cbc.Client.Log_Level = 2
   }
   if f.email != "" {
      err := f.profile()
      if err != nil {
         panic(err)
      }
   } else if f.id != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
