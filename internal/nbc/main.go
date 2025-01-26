package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/x/http"
   "flag"
   "log"
   "os"
   "path/filepath"
)

func main() {
   http.Transport{}.Set()
   log.SetFlags(log.Ltime)
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.IntVar(&f.nbc, "b", 0, "NBC ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.Parse()
   if f.nbc >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

type flags struct {
   nbc            int
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
