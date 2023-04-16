package youtube

import "2a.pages.dev/rosso/protobuf"

var Upload_Date = map[string]protobuf.Varint{
   "Last hour": 1,
   "Today": 2,
   "This week": 3,
   "This month": 4,
   "This year": 5,
}

var Type = map[string]protobuf.Varint{
   "Video": 1,
   "Channel": 2,
   "Playlist": 3,
   "Movie": 4,
}

var Duration = map[string]protobuf.Varint{
   "Under 4 minutes": 1,
   "4 - 20 minutes": 3,
   "Over 20 minutes": 2,
}

var Sort_By = map[string]protobuf.Varint{
   "Relevance": 0,
   "Upload date": 2,
   "View count": 3,
   "Rating": 1,
}

var Features = map[string]protobuf.Number{
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
   protobuf.Message
}

func (p Params) Sort_By(value protobuf.Varint) {
   p.Message[1] = value
}

func (p Params) Upload_Date(value protobuf.Varint) {
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

func New_Params() Params {
   var p Params
   p.Message = make(protobuf.Message)
   p.Message[2] = make(protobuf.Message)
   return p
}
