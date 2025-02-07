package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/tubi"
   "errors"
   "fmt"
   "os"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.name())
   if err != nil {
      return err
   }
   content := &tubi.Content{}
   err = content.Unmarshal(data)
   if err != nil {
      return err
   }
   if content.Series() {
      var ok bool
      content, ok = content.Get(f.tubi)
      if !ok {
         return errors.New("Content.Get")
      }
   }
   resource, ok := content.Resource()
   if !ok {
      return errors.New("Content.Resource")
   }
   represents, err := internal.Mpd(resource)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Widevine = resource
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) name() string {
   return fmt.Sprint(f.tubi) + ".txt"
}

func (f *flags) write_content() error {
   var content tubi.Content
   data, err := content.Marshal(f.tubi)
   if err != nil {
      return err
   }
   err = content.Unmarshal(data)
   if err != nil {
      return err
   }
   if content.Episode() {
      data, err = content.Marshal(content.SeriesId)
      if err != nil {
         return err
      }
   }
   return os.WriteFile(f.name(), data, os.ModePerm)
}
