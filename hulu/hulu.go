package hulu

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "path"
   "strings"
   "time"
)

func (a *Authenticate) Playlist(link *DeepLink) (*Playlist, error) {
   var p playlist_request
   p.ContentEabId = link.EabId
   p.DeejayDeviceId = 166
   p.Unencrypted = true
   p.Playback.Audio.Codecs.SelectionMode = "ALL"
   p.Playback.Audio.Codecs.Values = []codec_value{
      {Type: "AAC"},
      {Type: "EC3"},
   }
   p.Playback.Video.Codecs.SelectionMode = "ALL"
   p.Playback.Drm.SelectionMode = "ALL"
   p.Playback.Drm.Values = []drm_value{
      {
         SecurityLevel: "L3",
         Type:          "WIDEVINE",
         Version:       "MODULAR",
      },
   }
   p.Playback.Manifest.Type = "DASH"
   p.Playback.Segments.SelectionMode = "ALL"
   p.Playback.Segments.Values = func() []segment_value {
      var s segment_value
      s.Encryption.Mode = "CENC"
      s.Encryption.Type = "CENC"
      s.Type = "FMP4"
      return []segment_value{s}
   }()
   p.Playback.Version = 2 // needs to be exactly 2 for 1080p
   p.Version = 9999999
   p.Playback.Video.Codecs.Values = []codec_value{
      {
         Height:  9999,
         Level:   "9",
         Profile: "HIGH",
         Type:    "H264",
         Width:   9999,
      },
      {
         Height:  9999,
         Level:   "9",
         Profile: "MAIN_10",
         Tier:    "MAIN",
         Type:    "H265",
         Width:   9999,
      },
   }
   data, err := json.Marshal(p)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://play.hulu.com/v6/playlist", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "authorization": {"Bearer " + a.Data.UserToken},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   play := &Playlist{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

func (a *Authenticate) Details(link *DeepLink) (*Details, error) {
   data, err := json.Marshal(map[string][]string{
      "eabs": {link.EabId},
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://guide.hulu.com/guide/details", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer "+a.Data.UserToken)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var value struct {
      Items []Details
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Items[0], nil
}

func (Authenticate) Marshal(email, password string) ([]byte, error) {
   resp, err := http.PostForm(
      "https://auth.hulu.com/v2/livingroom/password/authenticate", url.Values{
         "friendly_name": {"!"},
         "password":      {password},
         "serial_number": {"!"},
         "user_email":    {email},
      },
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   return io.ReadAll(resp.Body)
}

func (a *Authenticate) DeepLink(id *EntityId) (*DeepLink, error) {
   req, err := http.NewRequest("", "https://discover.hulu.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/content/v5/deeplink/playback"
   req.URL.RawQuery = url.Values{
      "id":        {id.Data},
      "namespace": {"entity"},
   }.Encode()
   req.Header.Set("authorization", "Bearer "+a.Data.UserToken)
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
   link := &DeepLink{}
   err = json.NewDecoder(resp.Body).Decode(link)
   if err != nil {
      return nil, err
   }
   return link, nil
}

type Authenticate struct {
   Data struct {
      UserToken string `json:"user_token"`
   }
}

func (a *Authenticate) Unmarshal(data []byte) error {
   return json.Unmarshal(data, a)
}
func (p *Playlist) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      p.WvServer, "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type Playlist struct {
   StreamUrl string `json:"stream_url"`
   WvServer  string `json:"wv_server"`
}

type DeepLink struct {
   EabId string `json:"eab_id"`
}

type Details struct {
   EpisodeName   string `json:"episode_name"`
   EpisodeNumber int    `json:"episode_number"`
   Headline      string
   PremiereDate  time.Time `json:"premiere_date"`
   SeasonNumber  int       `json:"season_number"`
   SeriesName    string    `json:"series_name"`
}

func (d *Details) Show() string {
   return d.SeriesName
}

func (d *Details) Season() int {
   return d.SeasonNumber
}

func (d *Details) Episode() int {
   return d.EpisodeNumber
}

func (d *Details) Year() int {
   return d.PremiereDate.Year()
}

func (d *Details) Title() string {
   if d.EpisodeName != "" {
      return d.EpisodeName
   }
   return d.Headline
}

type EntityId struct {
   Data string
}

func (e *EntityId) String() string {
   return e.Data
}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
func (e *EntityId) Set(s string) error {
   e.Data = path.Base(s)
   return nil
}

type codec_value struct {
   Height  int    `json:"height,omitempty"`
   Level   string `json:"level,omitempty"`
   Profile string `json:"profile,omitempty"`
   Tier    string `json:"tier,omitempty"`
   Type    string `json:"type"`
   Width   int    `json:"width,omitempty"`
}

type drm_value struct {
   SecurityLevel string `json:"security_level"`
   Type          string `json:"type"`
   Version       string `json:"version"`
}

type playlist_request struct {
   ContentEabId   string `json:"content_eab_id"`
   DeejayDeviceId int    `json:"deejay_device_id"`
   Unencrypted    bool   `json:"unencrypted"`
   Version        int    `json:"version"`
   Playback       struct {
      Audio struct {
         Codecs struct {
            SelectionMode string        `json:"selection_mode"`
            Values        []codec_value `json:"values"`
         } `json:"codecs"`
      } `json:"audio"`
      Video struct {
         Codecs struct {
            SelectionMode string        `json:"selection_mode"`
            Values        []codec_value `json:"values"`
         } `json:"codecs"`
      } `json:"video"`
      Drm struct {
         SelectionMode string      `json:"selection_mode"`
         Values        []drm_value `json:"values"`
      } `json:"drm"`
      Manifest struct {
         Type string `json:"type"`
      } `json:"manifest"`
      Segments struct {
         SelectionMode string          `json:"selection_mode"`
         Values        []segment_value `json:"values"`
      } `json:"segments"`
      Version int `json:"version"`
   } `json:"playback"`
}

type segment_value struct {
   Encryption struct {
      Mode string `json:"mode"`
      Type string `json:"type"`
   } `json:"encryption"`
   Type string `json:"type"`
}
