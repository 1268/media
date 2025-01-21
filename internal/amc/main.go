package main

import (
   "41.neocities.org/media/amc"
   "41.neocities.org/media/internal"
   "41.neocities.org/text"
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
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.password, "p", "", "password")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.Parse()
   text.Transport{}.Set()
   switch {
   case f.email != "":
      err := f.login()
      if err != nil {
         panic(err)
      }
   case f.address.Nid != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

type flags struct {
   address        amc.Address
   email          string
   home           string
   password       string
   representation string
   s              internal.Stream
}

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home)
   f.s.ClientId = f.home + "/widevine/client_id.bin"
   f.s.PrivateKey = f.home + "/widevine/private_key.pem"
   return nil
}
