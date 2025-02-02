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
   http.Transport{}.DefaultClient()
   log.SetFlags(log.Ltime)
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.content_id, "b", "", "content ID")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.BoolVar(&f.write, "w", false, "write")
   flag.BoolVar(&f.intl, "n", false, "intl")
   flag.Parse()
   switch {
   case f.write:
      err := f.do_write()
      if err != nil {
         panic(err)
      }
   case f.content_id != "":
      err := f.do_read()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
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

type flags struct {
   content_id     string
   intl           bool
   representation string
   s              internal.Stream
   write          bool
}
