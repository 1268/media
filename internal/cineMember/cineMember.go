package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/cineMember"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
   "path"
)

func (f *flags) write_user() error {
   data, err := cineMember.Authenticate{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/cineMember.txt", data, os.ModePerm)
}

func (f *flags) base() string {
   return path.Base(f.address.String())
}

func (f *flags) write_play() error {
   data, err := os.ReadFile(f.home + "/cineMember.txt")
   if err != nil {
      return err
   }
   var user cineMember.Authenticate
   err = user.Unmarshal(data)
   if err != nil {
      return err
   }
   article, err := f.address.Article()
   if err != nil {
      return err
   }
   asset, ok := article.Film()
   if !ok {
      return errors.New("OperationArticle.Film")
   }
   data, err = cineMember.AssetPlay{}.Marshal(user, asset)
   if err != nil {
      return err
   }
   os.Mkdir(f.base(), os.ModePerm)
   return os.WriteFile(f.base()+"/play.txt", data, os.ModePerm)
}

func (f *flags) download() error {
   data, err := os.ReadFile(f.base() + "/play.txt")
   if err != nil {
      return err
   }
   var play cineMember.AssetPlay
   err = play.Unmarshal(data)
   if err != nil {
      return err
   }
   title, ok := play.Dash()
   if !ok {
      return errors.New("OperationPlay.Dash")
   }
   resp, err := http.Get(title.Manifest)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   var mpd dash.Mpd
   err = mpd.Unmarshal(data)
   if err != nil {
      return err
   }
   mpd.Set(resp.Request.URL)
   for represent := range mpd.Representation() {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Wrapper = title
         return f.s.Download(&represent)
      }
   }
   return nil
}
