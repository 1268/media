package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/rakuten"
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
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.Parse()
   text.Transport{}.Set()
   if f.address.String() != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

type flags struct {
   address        rakuten.Address
   representation string
   s              internal.Stream
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
