package mech

import (
   "2a.pages.dev/mech/widevine"
   "2a.pages.dev/rosso/dash"
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/mp4"
   "encoding/base64"
   "fmt"
   "net/url"
   "os"
)

type Stream struct {
   Client_ID string
   Info bool
   Namer
   Poster widevine.Poster
   Private_Key string
   base *url.URL
}

func (s *Stream) DASH(ref string) ([]dash.Representer, error) {
   client := http.Default_Client
   client.CheckRedirect = nil
   res, err := client.Get(ref)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   s.base = res.Request.URL
   return dash.Representers(res.Body)
}

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
   client := http.Default_Client
   client.CheckRedirect = nil
   req, err := http.Get_Parse(item.Segment_Template.Get_Initialization())
   if err != nil {
      return err
   }
   req.URL = s.base.ResolveReference(req.URL)
   res, err := client.Do(req)
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
   pssh, err := base64.StdEncoding.DecodeString(item.Widevine())
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
   client.Log_Level = 0
   for _, ref := range media {
      req.URL, err = s.base.Parse(ref)
      if err != nil {
         return err
      }
      res, err := client.Do(req)
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
