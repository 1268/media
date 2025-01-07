package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/itv"
   "encoding/xml"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/http/cookiejar"
   "path"
)

func (f *flags) download() error {
   var id itv.LegacyId
   err := id.Set(path.Base(f.address))
   if err != nil {
      return err
   }
   discovery, err := id.Discovery()
   if err != nil {
      return err
   }
   play, err := discovery.Playlist()
   if err != nil {
      return err
   }
   file, ok := play.Resolution1080()
   if !ok {
      return errors.New("resolution 1080")
   }
   http.DefaultClient.Jar, err = cookiejar.New(nil)
   if err != nil {
      return err
   }
   resp, err := http.Get(file.Href.Data)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   var mpd dash.Mpd
   xml.Unmarshal(data, &mpd)
   for represent := range mpd.Representation() {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Namer = itv.Namer{discovery}
         f.s.Wrapper = file
         return f.s.Download(&represent)
      }
   }
   return nil
}
