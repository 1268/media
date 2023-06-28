package mech

import (
   "2a.pages.dev/mech/widevine"
   "2a.pages.dev/rosso/dash"
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/mp4"
   "fmt"
   "net/url"
   "os"
   net_http "net/http"
)

func (s Stream) DASH_Get(items []dash.Representer, index int) error {
   if s.Info {
      for i, item := range items {
         fmt.Println()
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   item := items[index]
   file_name, err := Name(s)
   if err != nil {
      return err
   }
   file, err := os.Create(file_name + item.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   req, err := net_http.NewRequest(
      "GET", item.Segment_Template.Get_Initialization(), nil,
   )
   if err != nil {
      return err
   }
   req.URL = s.Base.ResolveReference(req.URL)
   res, err := new(net_http.Client).Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   dec := make(mp4.Decrypt)
   if err := dec.Init(res.Body, file); err != nil {
      return err
   }
   media := item.Segment_Template.Get_Media()
   pro := http.Progress_Chunks(file, len(media))
   private_key, err := os.ReadFile(s.Private_Key)
   if err != nil {
      return err
   }
   client_ID, err := os.ReadFile(s.Client_ID)
   if err != nil {
      return err
   }
   pssh, err := item.Widevine()
   if err != nil {
      return err
   }
   mod, err := widevine.New_Module(private_key, client_ID, pssh)
   if err != nil {
      return err
   }
   keys, err := mod.Post(s.Poster)
   if err != nil {
      return err
   }
   for _, ref := range media {
      req.URL, err = s.Base.Parse(ref)
      if err != nil {
         return err
      }
      res, err := new(net_http.Client).Do(req)
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      err = dec.Segment(res.Body, pro, keys.Content().Key)
      if err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}

type Stream struct {
   Base *url.URL
   Client_ID string
   Info bool
   Namer
   Private_Key string
   widevine.Poster
}
