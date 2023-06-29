package mech

import (
   "2a.pages.dev/rosso/dash"
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/mp4"
   "io"
   "net/url"
)

type Decrypt struct {
   Representer dash.Representer
   URL *url.URL
   Writer io.Writer
   decrypt mp4.Decrypt
}

func (d *Decrypt) Init() error {
   ref := d.Representer.Segment_Template.Get_Initialization()
   req, err := http.Get_Parse(ref)
   if err != nil {
      return err
   }
   req.URL = d.URL.ResolveReference(req.URL)
   res, err := new(http.Client).Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   d.decrypt = make(mp4.Decrypt)
   return d.decrypt.Init(res.Body, d.Writer)
}

func (d Decrypt) Media(key []byte) error {
   media := d.Representer.Segment_Template.Get_Media()
   pro := http.Progress_Chunks(d.Writer, len(media))
   for _, ref := range media {
      req, err := http.Get_Parse(ref)
      if err != nil {
         return err
      }
      req.URL = d.URL.ResolveReference(req.URL)
      res, err := new(http.Client).Do(req)
      if err != nil {
         return err
      }
      pro.Add_Chunk(res.ContentLength)
      if err := d.decrypt.Segment(res.Body, pro, key); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}
