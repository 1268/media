package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/ctv"
   "fmt"
   "io"
   "net/http"
   "os"
   "path"
   "sort"
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
   reps, err := dash.Unmarshal(data, resp.Request.URL)
   if err != nil {
      return err
   }
   sort.Slice(reps, func(i, j int) bool {
      return reps[i].Bandwidth < reps[j].Bandwidth
   })
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(&rep, "\n\n")
      case rep.Id:
         data, err = os.ReadFile(f.base() + "/media.txt")
         if err != nil {
            return err
         }
         var media ctv.MediaContent
         err = media.Unmarshal(data)
         if err != nil {
            return err
         }
         f.s.Namer = &ctv.Namer{media}
         f.s.Wrapper = ctv.Wrapper{}
         return f.s.Download(rep)
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
   err = os.WriteFile(f.base() + "/media.txt", data, os.ModePerm)
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
   return os.WriteFile(f.base() + "/manifest.txt", []byte(manifest), os.ModePerm)
}
