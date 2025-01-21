package main

import (
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
   flag.StringVar(&f.roku, "b", "", "Roku ID")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.BoolVar(&f.code_write, "code", false, "write code")
   flag.BoolVar(&f.token_write, "token", false, "write token")
   flag.BoolVar(&f.token_read, "t", false, "read token")
   flag.Parse()
   text.Transport{}.Set()
   switch {
   case f.code_write:
      err := write_code()
      if err != nil {
         panic(err)
      }
   case f.token_write:
      err := f.write_token()
      if err != nil {
         panic(err)
      }
   case f.roku != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
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

type flags struct {
   code_write     bool
   home           string
   representation string
   roku           string
   s              internal.Stream
   token_read     bool
   token_write    bool
}
