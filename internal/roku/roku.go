package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/roku"
   "fmt"
   "net/http"
   "os"
)

func (f *flags) download() error {
   var token *roku.AccountToken
   if f.token_read {
      data, err := os.ReadFile(f.home + "/roku.txt")
      if err != nil {
         return err
      }
      token = &roku.AccountToken{}
      err = token.Unmarshal(data)
      if err != nil {
         return err
      }
   }
   var auth roku.AccountAuth
   data, err := auth.Marshal(token)
   if err != nil {
      return err
   }
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   play, err := auth.Playback(f.roku)
   if err != nil {
      return err
   }
   resp, err := http.Get(play.Url)
   if err != nil {
      return err
   }
   represents, err := internal.Representation(resp)
   if err != nil {
      return err
   }
   for _, represent := range represents {
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

func (f *flags) write_token() error {
   // AccountAuth
   data, err := os.ReadFile("auth.txt")
   if err != nil {
      return err
   }
   var auth roku.AccountAuth
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   // AccountCode
   data, err = os.ReadFile("code.txt")
   if err != nil {
      return err
   }
   var code roku.AccountCode
   err = code.Unmarshal(data)
   if err != nil {
      return err
   }
   // AccountToken
   data, err = roku.AccountToken{}.Marshal(&auth, &code)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/roku.txt", data, os.ModePerm)
}

func write_code() error {
   // AccountAuth
   var auth roku.AccountAuth
   data, err := auth.Marshal(nil)
   if err != nil {
      return err
   }
   err = os.WriteFile("auth.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   // AccountCode
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   var code roku.AccountCode
   data, err = code.Marshal(&auth)
   if err != nil {
      return err
   }
   err = os.WriteFile("code.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   err = code.Unmarshal(data)
   if err != nil {
      return err
   }
   fmt.Println(code)
   return nil
}
