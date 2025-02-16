package max

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

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

func (n *Login) Unmarshal(data []byte) error {
   return json.Unmarshal(data, n)
}

const (
   device_info  = "!/!(!/!;!/!;!/!)"
   disco_client = "!:!:beam:!"
   prd_api      = "https://default.prd.api.discomax.com"
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

type Url [1]string

func (u *Url) UnmarshalText(data []byte) error {
   (*u)[0] = strings.Replace(string(data), "_fallback", "", 1)
   return nil
}

func (s *St) New() error {
   req, _ := http.NewRequest("", prd_api+"/token?realm=bolt", nil)
   req.Header = http.Header{
      "x-device-info":  {device_info},
      "x-disco-client": {disco_client},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   for _, cookie := range resp.Cookies() {
      if cookie.Name == "st" {
         (*s)[0] = cookie
         return nil
      }
   }
   return http.ErrNoCookie
}

// you must
// /authentication/linkDevice/initiate
// first or this will always fail
func (Login) Marshal(token St) ([]byte, error) {
   req, _ := http.NewRequest("POST", prd_api, nil)
   req.URL.Path = "/authentication/linkDevice/login"
   req.AddCookie(token[0])
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type St [1]*http.Cookie

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

type WatchUrl struct {
   EditId  string
   VideoId string
}

func (s St) Initiate() (*Initiate, error) {
   req, _ := http.NewRequest("POST", prd_api, nil)
   req.URL.Path = "/authentication/linkDevice/initiate"
   req.Header.Set("x-device-info", device_info)
   req.AddCookie(s[0])
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   value := &Initiate{}
   err = json.NewDecoder(resp.Body).Decode(value)
   if err != nil {
      return nil, err
   }
   return value, nil
}

type Initiate struct {
   Data struct {
      Attributes struct {
         LinkingCode string
         TargetUrl   string
      }
   }
}

type Login struct {
   Data struct {
      Attributes struct {
         Token string
      }
   }
}

func (n *Login) Playback(watch *WatchUrl) (*Playback, error) {
   data, err := json.Marshal(map[string]any{
      "consumptionType": "streaming",
      "editId": watch.EditId,
      "appBundle": "", // required
      "applicationSessionId": "", // required
      "firstPlay": false, // required
      "gdpr": false, // required
      "playbackSessionId": "", // required
      "userPreferences": struct{}{}, // required
      "capabilities": map[string]any{
         "manifests": map[string]any{
            "formats": map[string]any{
               "dash": struct{}{}, // required
            }, // required
         }, // required
      }, // required
      "deviceInfo": map[string]any{
         "player": map[string]any{
            "mediaEngine": map[string]string{
               "name": "", // required
               "version": "", // required
            }, // required
            "playerView": map[string]int{
               "height": 0, // required
               "width": 0, // required
            }, // required
            "sdk": map[string]string{
               "name": "", // required
               "version": "", // required
            }, // required
         }, // required
      }, // required
   })
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
      "authorization": {"Bearer " + n.Data.Attributes.Token},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var play Playback
   err = json.NewDecoder(resp.Body).Decode(&play)
   if err != nil {
      return nil, err
   }
   if err := play.Errors; len(err) >= 1 {
      return nil, errors.New(err[0].Message)
   }
   return &play, nil
}
