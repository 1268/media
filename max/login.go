package max

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

func (p *Playback) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      p.Drm.Schemes.Widevine.LicenseUrl, "application/x-protobuf",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   return io.ReadAll(resp.Body)
}

func (v *LinkLogin) Playback(watch *WatchUrl) (*Playback, error) {
   var body playback_request
   body.ConsumptionType = "streaming"
   body.EditId = watch.EditId
   data, err := json.MarshalIndent(body, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", prd_api, bytes.NewReader(data))
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b bytes.Buffer
      b.WriteString("/playback-orchestrator/any/playback-orchestrator/v1")
      b.WriteString("/playbackInfo")
      return b.String()
   }()
   req.Header = http.Header{
      "authorization": {"Bearer " + v.Data.Attributes.Token},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var resp_body Playback
   err = json.NewDecoder(resp.Body).Decode(&resp_body)
   if err != nil {
      return nil, err
   }
   if err := resp_body.Errors; len(err) >= 1 {
      return nil, errors.New(err[0].Message)
   }
   return &resp_body, nil
}

type Playback struct {
   Drm struct {
      Schemes struct {
         Widevine struct {
            LicenseUrl string
         }
      }
   }
   Errors []struct {
      Message string
   }
   Fallback struct {
      Manifest struct {
         Url Url
      }
   }
}

type playback_request struct {
   AppBundle            string `json:"appBundle"`            // required
   ApplicationSessionId string `json:"applicationSessionId"` // required
   Capabilities         struct {
      Manifests struct {
         Formats struct {
            Dash struct{} `json:"dash"` // required
         } `json:"formats"` // required
      } `json:"manifests"` // required
   } `json:"capabilities"` // required
   ConsumptionType string `json:"consumptionType"`
   DeviceInfo      struct {
      Player struct {
         MediaEngine struct {
            Name    string `json:"name"`    // required
            Version string `json:"version"` // required
         } `json:"mediaEngine"` // required
         PlayerView struct {
            Height int `json:"height"` // required
            Width  int `json:"width"`  // required
         } `json:"playerView"` // required
         Sdk struct {
            Name    string `json:"name"`    // required
            Version string `json:"version"` // required
         } `json:"sdk"` // required
      } `json:"player"` // required
   } `json:"deviceInfo"` // required
   EditId            string   `json:"editId"`
   FirstPlay         bool     `json:"firstPlay"`         // required
   Gdpr              bool     `json:"gdpr"`              // required
   PlaybackSessionId string   `json:"playbackSessionId"` // required
   UserPreferences   struct{} `json:"userPreferences"`   // required
}

func (w *WatchUrl) MarshalText() ([]byte, error) {
   var b bytes.Buffer
   if w.VideoId != "" {
      b.WriteString("/video/watch/")
      b.WriteString(w.VideoId)
   }
   if w.EditId != "" {
      b.WriteByte('/')
      b.WriteString(w.EditId)
   }
   return b.Bytes(), nil
}

func (w *WatchUrl) UnmarshalText(data []byte) error {
   s := string(data)
   if !strings.Contains(s, "/video/watch/") {
      return errors.New("/video/watch/ not found")
   }
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "play.max.com")
   s = strings.TrimPrefix(s, "/video/watch/")
   var found bool
   w.VideoId, w.EditId, found = strings.Cut(s, "/")
   if !found {
      return errors.New("/ not found")
   }
   return nil
}

type LinkLogin struct {
   Data struct {
      Attributes struct {
         Token string
      }
   }
}

func (v *LinkLogin) Unmarshal(data []byte) error {
   return json.Unmarshal(data, v)
}

type WatchUrl struct {
   EditId  string
   VideoId string
}

type Url struct {
   String string
}

func (f *Url) UnmarshalText(data []byte) error {
   f.String = strings.Replace(string(data), "_fallback", "", 1)
   return nil
}

// you must
// /authentication/linkDevice/initiate
// first or this will always fail
func (LinkLogin) Marshal(token *BoltToken) ([]byte, error) {
   req, err := http.NewRequest(
      "POST", prd_api+"/authentication/linkDevice/login", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("cookie", "st="+token.St)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

