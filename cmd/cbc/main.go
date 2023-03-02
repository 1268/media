package main

import (
   "2a.pages.dev/mech"
   "flag"
)

type flags struct {
   bandwidth int64
   email string
   id string
   mech.Stream
   name string
   password string
}

func main() {
   var f flags
   // b
   flag.StringVar(&f.id, "b", "", "ID")
   // e
   flag.StringVar(&f.email, "e", "", "email")
   // f
   flag.Int64Var(&f.bandwidth, "f", 2767479, "video bandwidth")
   // g
   flag.StringVar(&f.name, "g", "English", "audio name")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // p
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
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
