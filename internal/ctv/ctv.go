package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/ctv"
   "fmt"
   "io"
   "net/http"
   "os"
   "path"
   "slices"
)

func (f *flags) download() error {
   manifest, err := os.ReadFile(f.base() + "/manifest.txt")
   if err != nil {
      return err
   }
   resp, err := http.Get(string(manifest))
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   var mpd dash.Mpd
   err = mpd.Unmarshal(data)
   if err != nil {
      return err
   }
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
         f.s.Wrapper = ctv.Wrapper{}
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) base() string {
   return path.Base(f.address.Path)
}

func (f *flags) get_manifest() error {
   resolve, err := f.address.Resolve()
   if err != nil {
      return err
   }
   axis, err := resolve.Axis()
   if err != nil {
      return err
   }
   os.Mkdir(f.base(), os.ModePerm)
   // media
   var media ctv.MediaContent
   data, err := media.Marshal(axis)
   if err != nil {
      return err
   }
   err = os.WriteFile(f.base()+"/media.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   // manifest
   err = media.Unmarshal(data)
   if err != nil {
      return err
   }
   manifest, err := axis.Manifest(&media)
   if err != nil {
      return err
   }
   return os.WriteFile(f.base()+"/manifest.txt", []byte(manifest), os.ModePerm)
}
