package soundcloud

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
   "strconv"
   "strings"
   "time"
)

func Resolve(ref string) ([]Track, error) {
   req := http.Get()
   req.URL.Host = "api-v2.soundcloud.com"
   req.URL.Path = "/resolve"
   req.URL.RawQuery = url.Values{
      "client_id": {client_ID},
      "url": {ref},
   }.Encode()
   req.URL.Scheme = "https"
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var solve struct {
      Kind string
      Track
   }
   if err := json.NewDecoder(res.Body).Decode(&solve); err != nil {
      return nil, err
   }
   if solve.Kind == "track" {
      return []Track{solve.Track}, nil
   }
   return User_Tracks(solve.ID)
}

// We can also paginate, but for now this is good enough.
func User_Tracks(id int64) ([]Track, error) {
   var b []byte
   b = append(b, "https://api-v2.soundcloud.com/users/"...)
   b = strconv.AppendInt(b, id, 10)
   b = append(b, "/tracks"...)
   req, err := http.NewRequest("GET", string(b), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "client_id": {client_ID},
      "limit": {"999"},
   }.Encode()
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var user struct {
      Collection []Track
   }
   if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
      return nil, err
   }
   return user.Collection, nil
}

// i1.sndcdn.com/artworks-000308141235-7ep8lo-large.jpg
func (t Track) Artwork() string {
   if t.Artwork_URL == "" {
      t.Artwork_URL = t.User.Avatar_URL
   }
   return strings.Replace(t.Artwork_URL, "large", "t500x500", 1)
}

func (t Track) Name() string {
   return t.User.Username + "-" + t.Title
}

func (t Track) String() string {
   var b []byte
   b = append(b, "ID: "...)
   b = strconv.AppendInt(b, t.ID, 10)
   b = append(b, "\ndisplay date: "...)
   b = append(b, t.Display_Date...)
   b = append(b, "\nusername: "...)
   b = append(b, t.User.Username...)
   b = append(b, "\navatar URL: "...)
   b = append(b, t.User.Avatar_URL...)
   b = append(b, "\ntitle: "...)
   b = append(b, t.Title...)
   if t.Artwork_URL != "" {
      b = append(b, "\nartwork URL: "...)
      b = append(b, t.Artwork_URL...)
   }
   for _, coding := range t.Media.Transcodings {
      b = append(b, "\nformat: "...)
      b = append(b, coding.Format.Protocol...)
      b = append(b, "\nURL: "...)
      b = append(b, coding.URL...)
   }
   return string(b)
}

func (t Track) Time() (time.Time, error) {
   return time.Parse(time.RFC3339, t.Display_Date)
}
// Also available is "hls", but all transcodings are quality "sq".
// Same for "api-mobile.soundcloud.com".
func (t Track) Progressive() (*Media, error) {
   var ref string
   for _, code := range t.Media.Transcodings {
      if code.Format.Protocol == "progressive" {
         ref = code.URL
      }
   }
   req, err := http.NewRequest("GET", ref, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "client_id=" + client_ID
   res, err := Client.Do(req)
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

func New_Track(id int) (*Track, error) {
   req := http.Get()
   req.URL.Host = "api-v2.soundcloud.com"
   req.URL.Path = "/tracks/" + strconv.Itoa(id)
   req.URL.RawQuery = "client_id=" + client_ID
   req.URL.Scheme = "https"
   res, err := http.Default_Client.Do(req)
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

