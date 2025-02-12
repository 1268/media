package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/rtbf"
   "errors"
   "fmt"
   "net/http"
   "os"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/rtbf.txt")
   if err != nil {
      return err
   }
   var login rtbf.AuvioLogin
   err = login.Unmarshal(data)
   if err != nil {
      return err
   }
   token, err := login.Token()
   if err != nil {
      return err
   }
   auth, err := token.Auth()
   if err != nil {
      return err
   }
   page, err := f.address.Page()
   if err != nil {
      return err
   }
   asset_id, ok := page.GetAssetId()
   if !ok {
      return errors.New("AuvioPage.GetAssetId")
   }
   title, err := auth.Entitlement(asset_id)
   if err != nil {
      return err
   }
   address, ok := title.Dash()
   if !ok {
      return errors.New("Entitlement.Dash")
   }
   resp, err := http.Get(address)
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
         f.s.Widevine = title
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   data, err := rtbf.AuvioLogin{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/rtbf.txt", data, os.ModePerm)
}
