package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/kanopy"
   "41.neocities.org/media/internal"
   "41.neocities.org/text"
   "flag"
   "fmt"
   "io"
   "net/http"
   "os"
   "path/filepath"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/kanopy.txt")
   if err != nil {
      return err
   }
   var auth kanopy.Authenticate
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   deep, err := auth.DeepLink(&f.entity)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
   resp, err := http.Get(play.StreamUrl)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   var mpd dash.Mpd
   mpd.Unmarshal(data)
   for represent := range mpd.Representation() {
      if *represent.Width < f.min_width {
         if *represent.MimeType == "video/mp4" {
            continue
         }
      }
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Wrapper = play
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   data, err := kanopy.Authenticate{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/kanopy.txt", data, os.ModePerm)
}
func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.entity, "a", "address")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.Int64Var(&f.min_width, "m", 1280, "min width")
   flag.Parse()
   text.Transport{}.Set()
   switch {
   case f.password != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.entity.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

type flags struct {
   email          string
   entity         kanopy.EntityId
   home           string
   min_width      int64
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
