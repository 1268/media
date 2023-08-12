package youtube

import "154.pages.dev/encoding/protobuf"

var values = map[string]map[string]uint64{
   "UPLOAD DATE": map[string]uint64{
      "Last hour": 1,
      "Today": 2,
      "This week": 3,
      "This month": 4,
      "This year": 5,
   },
   "TYPE": map[string]uint64{
      "Video": 1,
      "Channel": 2,
      "Playlist": 3,
      "Movie": 4,
   },
   "DURATION": map[string]uint64{
      "Under 4 minutes": 1,
      "4 - 20 minutes": 3,
      "Over 20 minutes": 2,
   },
   "FEATURES": map[string]uint64{
      "Live": 8,
      "4K": 14,
      "HD": 4,
      "Subtitles/CC": 5,
      "Creative Commons": 6,
      "360°": 15,
      "VR180": 26,
      "3D": 7,
      "HDR": 25,
      "Location": 23,
      "Purchased": 9,
   },
   "SORT BY": map[string]uint64{
      "Relevance": 0,
      "Upload date": 2,
      "View count": 3,
      "Rating": 1,
   },
}

type filter struct {
   Upload_Date uint64 // 2 1
   Type uint64 // 2 2
   Duration uint64 // 2 3
   Features []uint64 // 2
}

type parameters struct {
   Sort_By *uint64 // 1
   Filter *filter
}

func (p parameters) MarshalBinary() ([]byte, error) {
   var m protobuf.Message
   if p.Sort_By != nil {
      m = append(m, protobuf.Number(1).Varint(*p.Sort_By))
   }
   return m.Append(nil), nil
}
