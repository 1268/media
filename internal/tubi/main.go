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
   var port http.Transport
   port.ProxyFromEnvironment()
   port.DefaultClient()
   log.SetFlags(log.Ltime)
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.IntVar(&f.tubi, "b", 0, "Tubi ID")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.BoolVar(&f.content, "w", false, "write content")
   flag.Parse()
   switch {
   case f.content:
      err := f.write_content()
      if err != nil {
         panic(err)
      }
   case f.tubi >= 1:
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

type flags struct {
   content        bool
   representation string
   s              internal.Stream
   tubi           int
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
