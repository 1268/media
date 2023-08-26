// Support for this software's development was paid for by
// Fredrick R. Brennan's Modular Font Editor K Foundation, Inc.
package twitter

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
)

func (t Tweet_Result) Video_MP4() []Variant {
   var variants []Variant
   for _, media := range t.Media_Details {
      for _, variant := range media.Video_Info.Variants {
         if variant.Content_Type == "video/mp4" {
            variants = append(variants, variant)
         }
      }
   }
   return variants
}

type Tweet_Result struct {
   Media_Details []struct {
      Video_Info struct {
         Variants []Variant
      }
   } `json:"mediaDetails"`
}

type Variant struct {
   Bitrate int
   Content_Type string
   URL string
}

func New_Tweet_Result(id int64) (*Tweet_Result, error) {
   req, err := http.NewRequest("GET", "https://cdn.syndication.twimg.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/tweet-result"
   req.URL.RawQuery = url.Values{
      "id": {strconv.FormatInt(id, 10)},
      "token": {"-"},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tweet := new(Tweet_Result)
   if err := json.NewDecoder(res.Body).Decode(tweet); err != nil {
      return nil, err
   }
   return tweet, nil
}
