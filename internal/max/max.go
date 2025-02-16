package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/max"
   "fmt"
   "os"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/max.txt")
   if err != nil {
      return err
   }
   var login max.Login
   err = login.Unmarshal(data)
   if err != nil {
      return err
   }
   play, err := login.Playback(&f.url)
   if err != nil {
      return err
   }
   represents, err := internal.Mpd(play)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Client = play
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) do_initiate() error {
   var st max.St
   err := st.New()
   if err != nil {
      return err
   }
   os.WriteFile("st.txt", st, os.ModePerm)
   initiate, err := st.Initiate()
   if err != nil {
      return err
   }
   fmt.Printf("%+v\n", initiate)
   return nil
}

func (f *flags) do_login() error {
   data, err := os.ReadFile("st.txt")
   if err != nil {
      return err
   }
   var st max.St
   data, err = max.Login{}.Marshal(&st)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/max.txt", data, os.ModePerm)
}
