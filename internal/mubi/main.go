package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/mubi"
   "41.neocities.org/x/http"
   "flag"
   "os"
   "path/filepath"
)

func main() {
   http.Transport{}.DefaultClient()
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.address, "a", "address")
   flag.BoolVar(&f.auth, "auth", false, "authenticate")
   flag.BoolVar(&f.code, "code", false, "link code")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.BoolVar(&f.secure, "w", false, "secure URL")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.Parse()
   switch {
   case f.auth:
      err := f.write_auth()
      if err != nil {
         panic(err)
      }
   case f.code:
      err := write_code()
      if err != nil {
         panic(err)
      }
   case f.secure:
      err := f.write_secure()
      if err != nil {
         panic(err)
      }
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
   address        mubi.Address
   auth           bool
   code           bool
   home           string
   representation string
   s              internal.Stream
   secure         bool
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
