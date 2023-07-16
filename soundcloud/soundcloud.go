package soundcloud

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

// Also available is "hls", but all transcodings are quality "sq".
// Same for "api-mobile.soundcloud.com".
func (t Track) Progressive() (*Media, error) {
   ref := func() string {
      for _, code := range t.Media.Transcodings {
         if code.Format.Protocol == "progressive" {
            return code.URL
         }
      }
      return ""
   }
   req, err := http.NewRequest("GET", ref(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "client_id=" + client_ID
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   med := new(Media)
   if err := json.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}

func (t Track) String() string {
   var b []byte
   b = append(b, "ID: "...)
   b = strconv.AppendInt(b, t.ID, 10)
   b = append(b, "\ndisplay date: "...)
   b = append(b, t.Display_Date...)
   b = append(b, "\nusername: "...)
   b = append(b, t.User.Username...)
   b = append(b, "\navatar: "...)
   b = append(b, t.User.Avatar_URL...)
   b = append(b, "\ntitle: "...)
   b = append(b, t.Title...)
   if t.Artwork_URL != "" {
      b = append(b, "\nartwork: "...)
      b = append(b, t.Artwork_URL...)
   }
   for _, coding := range t.Media.Transcodings {
      b = append(b, "\n\nformat: "...)
      b = append(b, coding.Format.Protocol...)
      b = append(b, "\nURL: "...)
      b = append(b, coding.URL...)
   }
   return string(b)
}

type Media struct {
   URL string // cf-media.sndcdn.com/QaV7QR1lxpc6.128.mp3
}

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
const client_ID = "iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX"

type Track struct {
   ID int64
   Display_Date string // 2021-04-12T07:00:01Z
   User struct {
      Username string
      Avatar_URL string
   }
   Title string
   Artwork_URL string
   Media struct {
      Transcodings []struct {
         Format struct {
            Protocol string
         }
         URL string
      }
   }
}
func Resolve(ref string) (*Track, error) {
   req, err := http.NewRequest(
      "GET", "https://api-v2.soundcloud.com/resolve", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "client_id": {client_ID},
      "url": {ref},
   }.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var solve struct {
      Track Track
   }
   if err := json.NewDecoder(res.Body).Decode(&solve); err != nil {
      return nil, err
   }
   return &solve.Track, nil
}

func New_Track(id int) (*Track, error) {
   req, err := http.NewRequest("GET", "https://api-v2.soundcloud.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/tracks/" + strconv.Itoa(id)
   req.URL.RawQuery = "client_id=" + client_ID
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tra := new(Track)
   if err := json.NewDecoder(res.Body).Decode(tra); err != nil {
      return nil, err
   }
   return tra, nil
}

// i1.sndcdn.com/artworks-000308141235-7ep8lo-large.jpg
func (t Track) Artwork() string {
   if t.Artwork_URL == "" {
      t.Artwork_URL = t.User.Avatar_URL
   }
   return strings.Replace(t.Artwork_URL, "large", "t500x500", 1)
}

func (t Track) Time() (time.Time, error) {
   return time.Parse(time.RFC3339, t.Display_Date)
}

