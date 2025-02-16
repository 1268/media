package main

import (
   "41.neocities.org/media/ctv"
   "41.neocities.org/media/internal"
   "fmt"
   "os"
)

func (f *flags) get_manifest() error {
   resolve, err := f.address.Resolve()
   if err != nil {
      return err
   }
   axis, err := resolve.Axis()
   if err != nil {
      return err
   }
   content, err := axis.Content()
   if err != nil {
      return err
   }
   data, err := ctv.Manifest{}.Marshal(axis, content)
   if err != nil {
      return err
   }
   return os.WriteFile("manifest.txt", data, os.ModePerm)
}

func (f *flags) download() error {
   data, err := os.ReadFile("manifest.txt")
   if err != nil {
      return err
   }
   var manifest ctv.Manifest
   err = manifest.Unmarshal(data)
   if err != nil {
      return err
   }
   represents, err := internal.Mpd(manifest)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Client = ctv.Client{}
         return f.s.Download(&represent)
      }
   }
   return nil
}
