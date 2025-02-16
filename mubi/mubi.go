package mubi

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strconv"
   "strings"
)

type Film struct {
   Id    int64
   Title string
   Year  int
}

func (SecureUrl) Marshal(auth *Authenticate, film1 *Film) ([]byte, error) {
   req, _ := http.NewRequest("", "https://api.mubi.com", nil)
   req.URL.Path = func() string {
      b := []byte("/v3/films/")
      b = strconv.AppendInt(b, film1.Id, 10)
      b = append(b, "/viewing/secure_url"...)
      return string(b)
   }()
   req.Header = http.Header{
      "authorization":  {"Bearer " + auth.Token},
      "client":         {client},
      "client-country": {ClientCountry},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var data strings.Builder
      resp.Write(&data)
      return nil, errors.New(data.String())
   }
   return io.ReadAll(resp.Body)
}

var ClientCountry = "US"

// "android" requires headers:
// client-device-identifier
// client-version
const client = "web"

func (a Address) String() string {
   return a[0]
}

func (a *Address) Set(data string) error {
   var found bool
   _, (*a)[0], found = strings.Cut(data, "/films/")
   if !found {
      return errors.New("/films/ not found")
   }
   return nil
}

type Address [1]string

func (a *Authenticate) Unmarshal(data []byte) error {
   return json.Unmarshal(data, a)
}

func (a Address) Film() (*Film, error) {
   req, _ := http.NewRequest("", "https://api.mubi.com", nil)
   req.URL.Path = "/v3/films/" + a[0]
   req.Header = http.Header{
      "client":         {client},
      "client-country": {ClientCountry},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   film1 := &Film{}
   err = json.NewDecoder(resp.Body).Decode(film1)
   if err != nil {
      return nil, err
   }
   return film1, nil
}

// Mubi do this sneaky thing. you cannot download a video unless you have told
// the API that you are watching it. so you have to call
// `/v3/films/%v/viewing`, otherwise it wont let you get the MPD. if you have
// already viewed the video on the website that counts, but if you only use the
// tool it will error
func (a *Authenticate) Viewing(film1 *Film) error {
   req, _ := http.NewRequest("POST", "https://api.mubi.com", nil)
   req.URL.Path = func() string {
      b := []byte("/v3/films/")
      b = strconv.AppendInt(b, film1.Id, 10)
      b = append(b, "/viewing"...)
      return string(b)
   }()
   req.Header = http.Header{
      "authorization":  {"Bearer " + a.Token},
      "client":         {client},
      "client-country": {ClientCountry},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   var value struct {
      Message string
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return err
   }
   if value.Message != "" {
      return errors.New(value.Message)
   }
   return nil
}

func (LinkCode) Marshal() ([]byte, error) {
   req, _ := http.NewRequest("", "https://api.mubi.com/v3/link_code", nil)
   req.Header = http.Header{
      "client":         {client},
      "client-country": {ClientCountry},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var data strings.Builder
      resp.Write(&data)
      return nil, errors.New(data.String())
   }
   return io.ReadAll(resp.Body)
}

type Authenticate struct {
   Token string
   User  struct {
      Id int
   }
}

func (Authenticate) Marshal(code *LinkCode) ([]byte, error) {
   data, err := json.Marshal(map[string]string{"auth_token": code.AuthToken})
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.mubi.com/v3/authenticate", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "client":         {client},
      "client-country": {ClientCountry},
      "content-type":   {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var data bytes.Buffer
      resp.Write(&data)
      return nil, errors.New(data.String())
   }
   return io.ReadAll(resp.Body)
}

func (c *LinkCode) String() string {
   var b strings.Builder
   b.WriteString("TO LOG IN AND START WATCHING\n")
   b.WriteString("Go to\n")
   b.WriteString("mubi.com/android\n")
   b.WriteString("and enter the code below\n")
   b.WriteString(c.LinkCode)
   return b.String()
}

func (c *LinkCode) Unmarshal(data []byte) error {
   return json.Unmarshal(data, c)
}

func (s *SecureUrl) Unmarshal(data []byte) error {
   return json.Unmarshal(data, s)
}

func (s status) Error() string {
   return strings.ToLower(s[0])
}

type status [1]string

var forbidden = status{"HTTP Status 403 – Forbidden"}

func (t *TextTrack) String() string {
   return "id = " + t.Id
}

func (a *Authenticate) License(data []byte) ([]byte, error) {
   // final slash is needed
   req, err := http.NewRequest(
      "POST", "https://lic.drmtoday.com/license-proxy-widevine/cenc/",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   data, err = json.Marshal(map[string]any{
      "merchant":  "mubi",
      "sessionId": a.Token,
      "userId":    a.User.Id,
   })
   if err != nil {
      return nil, err
   }
   req.Header.Set("dt-custom-data", base64.StdEncoding.EncodeToString(data))
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   if strings.Contains(string(data), forbidden[0]) {
      return nil, forbidden
   }
   var value struct {
      License []byte
   }
   err = json.Unmarshal(data, &value)
   if err != nil {
      return nil, err
   }
   return value.License, nil
}

type LinkCode struct {
   AuthToken string `json:"auth_token"`
   LinkCode  string `json:"link_code"`
}

type TextTrack struct {
   Id  string
   Url string
}

type SecureUrl struct {
   TextTrackUrls []TextTrack `json:"text_track_urls"`
   Url           string
}

func (s *SecureUrl) Mpd() (*http.Response, error) {
   return http.Get(s.Url)
}
