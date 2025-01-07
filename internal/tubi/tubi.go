package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/tubi"
   "encoding/xml"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.name())
   if err != nil {
      return err
   }
   content := &tubi.VideoContent{}
   err = content.Unmarshal(data)
   if err != nil {
      return err
   }
   if content.Series() {
      var ok bool
      content, ok = content.Get(f.tubi)
      if !ok {
         return errors.New("VideoContent.Get")
      }
   }
   video, ok := content.Video()
   if !ok {
      return errors.New("VideoContent.Video")
   }
   resp, err := http.Get(video.Manifest.Url)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
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
         f.s.Namer = tubi.Namer{content}
         f.s.Wrapper = video
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) name() string {
   return fmt.Sprint(f.tubi) + ".txt"
}

func (f *flags) write_content() error {
   var content tubi.VideoContent
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
