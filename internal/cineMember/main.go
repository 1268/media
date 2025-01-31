package main

import (
   "41.neocities.org/media/cineMember"
   "41.neocities.org/media/internal"
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
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.BoolVar(&f.play, "w", false, "operation play")
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   switch {
   case f.password != "":
      err := f.write_user()
      if err != nil {
         panic(err)
      }
   case f.play:
      err := f.write_play()
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
   address        cineMember.Address
   email          string
   home           string
   password       string
   play           bool
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
