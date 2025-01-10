package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/criterion"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
   "path"
   "slices"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/criterion.txt")
   if err != nil {
      return err
   }
   var token criterion.AuthToken
   err = token.Unmarshal(data)
   if err != nil {
      return err
   }
   item, err := token.Video(path.Base(f.address))
   if err != nil {
      return err
   }
   files, err := token.Files(item)
   if err != nil {
      return err
   }
   file, ok := files.Dash()
   if !ok {
      return errors.New("VideoFiles.Dash")
   }
   resp, err := http.Get(file.Links.Source.Href)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   var mpd dash.Mpd
   mpd.BaseUrl = &dash.Url{resp.Request.URL}
   mpd.Unmarshal(data)
   represents := slices.SortedFunc(mpd.Representation(),
      func(a, b dash.Representation) int {
         return a.Bandwidth - b.Bandwidth
      },
   )
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Namer = item
         f.s.Wrapper = file
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   data, err := criterion.AuthToken{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/criterion.txt", data, os.ModePerm)
}
