package youtube

import "154.pages.dev/protobuf"

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
   Upload_Date uint64 // 1
   Type uint64 // 2
   Duration uint64 // 3
   Features []uint64
}

type parameters struct {
   Sort_By uint64 // 1
   Filter *filter // 2
}

func (p parameters) Marshal() []byte {
   var m protobuf.Message
   if p.Sort_By >= 1 {
      m.Add_Varint(1, p.Sort_By)
   }
   if p.Filter != nil {
      m.Add(2, func(m *protobuf.Message) {
         if p.Filter.Upload_Date >= 1 {
            m.Add_Varint(1, p.Filter.Upload_Date)
         }
         if p.Filter.Type >= 1 {
            m.Add_Varint(2, p.Filter.Type)
         }
         if p.Filter.Duration >= 1 {
            m.Add_Varint(3, p.Filter.Duration)
         }
         for _, feature := range p.Filter.Features {
            m.Add_Varint(protobuf.Number(feature), 1)
         }
      })
   }
   return m.Append(nil)
}
