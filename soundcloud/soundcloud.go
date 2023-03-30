package soundcloud

import (
   "net/url"
   "path"
)

type Media struct {
   URL string // cf-media.sndcdn.com/QaV7QR1lxpc6.128.mp3
}

func (m Media) Ext() (string, error) {
   ref, err := url.Parse(m.URL)
   if err != nil {
      return "", err
   }
   return path.Ext(ref.Path), nil
}

const client_ID = "iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX"

type Image struct {
   Size string
   Crop bool
}

var Images = []Image{
   {Size: "t120x120"},
   {Size: "t1240x260", Crop: true},
   {Size: "t200x200"},
   {Size: "t20x20"},
   {Size: "t240x240"},
   {Size: "t2480x520", Crop: true},
   {Size: "t250x250"},
   {Size: "t300x300"},
   {Size: "t40x40"},
   {Size: "t47x47"},
   {Size: "t500x"},
   {Size: "t500x500"},
   {Size: "t50x50"},
   {Size: "t60x60"},
   {Size: "t67x67"},
   {Size: "t80x80"},
   {Size: "tx250"},
}
