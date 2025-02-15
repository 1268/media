package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/roku"
   "fmt"
   "os"
)

func (f *flags) download() error {
   var code *roku.Code
   if f.token_read {
      data, err := os.ReadFile(f.home + "/roku.txt")
      if err != nil {
         return err
      }
      code = &roku.Code{}
      err = code.Unmarshal(data)
      if err != nil {
         return err
      }
   }
   var token roku.Token
   data, err := token.Marshal(code)
   if err != nil {
      return err
   }
   err = token.Unmarshal(data)
   if err != nil {
      return err
   }
   play, err := token.Playback(f.roku)
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

func (f *flags) write_token() error {
   data, err := os.ReadFile("activation.txt")
   if err != nil {
      return err
   }
   var activation roku.Activation
   err = activation.Unmarshal(data)
   if err != nil {
      return err
   }
   data, err = os.ReadFile("token.txt")
   if err != nil {
      return err
   }
   var token roku.Token
   err = token.Unmarshal(data)
   if err != nil {
      return err
   }
   data, err = roku.Code{}.Marshal(&activation, &token)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/roku.txt", data, os.ModePerm)
}

func write_code() error {
   var token roku.Token
   data, err := token.Marshal(nil)
   if err != nil {
      return err
   }
   err = os.WriteFile("token.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   err = token.Unmarshal(data)
   if err != nil {
      return err
   }
   var activation roku.Activation
   data, err = activation.Marshal(&token)
   if err != nil {
      return err
   }
   err = os.WriteFile("activation.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   err = activation.Unmarshal(data)
   if err != nil {
      return err
   }
   fmt.Println(activation)
   return nil
}
