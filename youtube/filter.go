package youtube

import "2a.pages.dev/rosso/protobuf"

type Filters struct {
   protobuf.Message
}

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

func (f Filters) Sort_By(value protobuf.Varint) {
   f.Message[1] = value
}

func (f Filters) Upload_Date(value protobuf.Varint) {
   f.Get(2)[1] = value
}

func (f Filters) Type(value protobuf.Varint) {
   f.Get(2)[2] = value
}

func (f Filters) Duration(value protobuf.Varint) {
   f.Get(2)[3] = value
}

func (f Filters) Features(num protobuf.Number) {
   f.Get(2)[num] = protobuf.Varint(1)
}

func New_Filters() Filters {
   var f Filters
   f.Message = make(protobuf.Message)
   f.Message[2] = make(protobuf.Message)
   return f
}
