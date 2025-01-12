package rakuten

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type OnDemand struct {
   AudioLanguage            string `json:"audio_language"`
   AudioQuality             string `json:"audio_quality"`
   ClassificationId         int    `json:"classification_id"`
   ContentId                string `json:"content_id"`
   ContentType              string `json:"content_type"`
   DeviceIdentifier         string `json:"device_identifier"`
   DeviceSerial             string `json:"device_serial"`
   DeviceStreamVideoQuality string `json:"device_stream_video_quality"`
   Player                   string `json:"player"`
   SubtitleLanguage         string `json:"subtitle_language"`
   VideoType                string `json:"video_type"`
}

func (a *Address) video(quality string) *OnDemand {
   var o OnDemand
   o.AudioLanguage = "ENG"
   o.AudioQuality = "2.0"
   o.ContentType = "movies"
   o.DeviceSerial = "!"
   o.Player = "atvui40:DASH-CENC:WVM"
   o.SubtitleLanguage = "MIS"
   o.VideoType = "stream"
   o.DeviceIdentifier = "atvui40"
   o.ClassificationId = a.ClassificationId
   o.ContentId = a.ContentId
   o.DeviceStreamVideoQuality = quality
   return &o
}

var classification_id = map[string]int{
   "cz": 272,
   "dk": 283,
   "fi": 284,
   "fr": 23,
   "ie": 41,
   "it": 36,
   "nl": 323,
   "no": 286,
   "pt": 64,
   "se": 282,
   "ua": 276,
   "uk": 18,
}

// geo block
func (o *OnDemand) Info() (*StreamInfo, error) {
   data, err := json.MarshalIndent(o, "", " ")
   if err != nil {
      return nil, err
   }
   resp, err := http.Post(
      "https://gizmo.rakuten.tv/v3/avod/streamings",
      "application/json", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   var value struct {
      Data struct {
         StreamInfos []StreamInfo `json:"stream_infos"`
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data.StreamInfos[0], nil
}

func (*GizmoMovie) Show() string {
   return ""
}

func (g *GizmoMovie) Title() string {
   return g.Data.Title
}

func (*GizmoMovie) Season() int {
   return 0
}

func (*GizmoMovie) Episode() int {
   return 0
}

func (g *GizmoMovie) Year() int {
   return g.Data.Year
}

func (a *Address) Hd() *OnDemand {
   return a.video("HD")
}

func (a *Address) Fhd() *OnDemand {
   return a.video("FHD")
}

type StreamInfo struct {
   LicenseUrl   string `json:"license_url"`
   Url          string
   VideoQuality string `json:"video_quality"`
}

func (s *StreamInfo) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      s.LicenseUrl, "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type Address struct {
   ClassificationId int
   ContentId        string
   MarketCode       string
}

type GizmoMovie struct {
   Data struct {
      Title string
      Year  int
   }
}

func (a *Address) Movie() (*GizmoMovie, error) {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/movies/" + a.ContentId
   req.URL.RawQuery = url.Values{
      "market_code":       {a.MarketCode},
      "classification_id": {strconv.Itoa(a.ClassificationId)},
      "device_identifier": {"atvui40"},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   movie := &GizmoMovie{}
   err = json.NewDecoder(resp.Body).Decode(movie)
   if err != nil {
      return nil, err
   }
   return movie, nil
}

func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   s = strings.TrimPrefix(s, "rakuten.tv")
   s = strings.TrimPrefix(s, "/")
   var found bool
   a.MarketCode, a.ContentId, found = strings.Cut(s, "/movies/")
   if !found {
      return errors.New("/movies/ not found")
   }
   a.ClassificationId, found = classification_id[a.MarketCode]
   if !found {
      return errors.New("MarketCode not found")
   }
   return nil
}

func (a *Address) String() string {
   var b strings.Builder
   if a.MarketCode != "" {
      b.WriteString(a.MarketCode)
   }
   if a.ContentId != "" {
      b.WriteString("/movies/")
      b.WriteString(a.ContentId)
   }
   return b.String()
}
