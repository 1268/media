package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/pluto"
   "41.neocities.org/log"
   "flag"
   "os"
   "path/filepath"
)

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.address, "a", "address")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.BoolVar(&f.get_forward, "g", false, "get forward")
   flag.StringVar(&f.set_forward, "s", "", "set forward")
   flag.Parse()
   log.Transport{}.Set()
   switch {
   case f.get_forward:
      get_forward()
   case f.address.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

type flags struct {
   s              internal.Stream
   representation string
   address        pluto.Address
   get_forward    bool
   set_forward    string
}

func (f *flags) New() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home = filepath.ToSlash(home)
   f.s.ClientId = home + "/widevine/client_id.bin"
   f.s.PrivateKey = home + "/widevine/private_key.pem"
   return nil
}
