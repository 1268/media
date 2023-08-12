package youtube

import (
   "154.pages.dev/encoding/protobuf"
   "154.pages.dev/http/option"
   "154.pages.dev/strconv"
   "fmt"
   "io"
   "mime"
   "net/http"
)

var Upload_Date = map[string]int{
   "Last hour": 1,
   "Today": 2,
   "This week": 3,
   "This month": 4,
   "This year": 5,
}

var Type = map[string]int{
   "Video": 1,
   "Channel": 2,
   "Playlist": 3,
   "Movie": 4,
}

var Duration = map[string]int{
   "Under 4 minutes": 1,
   "4 - 20 minutes": 3,
   "Over 20 minutes": 2,
}

var Features = map[string]int{
   "360°": 15,
   "3D": 7,
   "4K": 14,
   "Creative Commons": 6,
   "HD": 4,
   "HDR": 25,
   "Live": 8,
   "Location": 23,
   "Purchased": 9,
   "Subtitles/CC": 5,
   "VR180": 26,
}

type Params struct {
   m protobuf.Message
}

func (p *Params) Sort_By(value uint64) {
   p.m = append(p.m,
      protobuf.Number(1).Varint(value),
   )
}

var Sort_By = map[string]uint64{
   "Relevance": 0,
   "Upload date": 2,
   "View count": 3,
   "Rating": 1,
}

func (p Params) Upload_Date(value int) {
   p.Get(2)[1] = value
}

func (p Params) Type(value protobuf.Varint) {
   p.Get(2)[2] = value
}

func (p Params) Duration(value protobuf.Varint) {
   p.Get(2)[3] = value
}

func (p Params) Features(num protobuf.Number) {
   p.Get(2)[num] = protobuf.Varint(1)
}
