package gem

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func New_Token(username, password string) (*Token, error) {
   body := url.Values{
      "client_id": {"7f44c935-6542-4ce7-ae05-eb887809741c"},
      "grant_type": {"password"},
      "password": {password},
      "scope": {strings.Join(scope, " ")},
      "username": {username},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://rcmnb2cprod.b2clogin.com", strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/rcmnb2cprod.onmicrosoft.com/B2C_1A_ExternalClient_ROPC_Auth/oauth2/v2.0/token"
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tok := new(Token)
   if err := json.NewDecoder(res.Body).Decode(tok); err != nil {
      return nil, err
   }
   return tok, nil
}

const forwarded_for = "99.224.0.0"

func (p Profile) Write_File(name string) error {
   text, err := json.MarshalIndent(p, "", " ")
   if err != nil {
      return err
   }
   return os.WriteFile(name, text, 0666)
}

func Read_Profile(name string) (*Profile, error) {
   text, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   pro := new(Profile)
   if err := json.Unmarshal(text, pro); err != nil {
      return nil, err
   }
   return pro, nil
}

type Profile struct {
   Claims_Token string `json:"claimsToken"`
}
func (t Token) Profile() (*Profile, error) {
   req, err := http.NewRequest("GET", "https://services.radio-canada.ca", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/ott/subscription/v2/gem/Subscriber/profile"
   req.URL.RawQuery = "device=phone_android"
   req.Header.Set("Authorization", "Bearer " + t.Access_Token)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   pro := new(Profile)
   if err := json.NewDecoder(res.Body).Decode(pro); err != nil {
      return nil, err
   }
   return pro, nil
}

type Token struct {
   Access_Token string
}

const manifest_type = "desktop"

func (p Profile) Media(item *Lineup_Item) (*Media, error) {
   req, err := http.NewRequest(
      "GET", "https://services.radio-canada.ca/media/validation/v2", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "appCode": {"gem"},
      "idMedia": {item.Formatted_ID_Media},
      "manifestType": {manifest_type},
      "output": {"json"},
      // you need this one the first request for a video, but can omit after
      // that. we will just send it every time.
      "tech": {"hls"},
   }.Encode()
   req.Header = http.Header{
      "X-Claims-Token": {p.Claims_Token},
      "X-Forwarded-For": {forwarded_for},
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   m := new(Media)
   if err := json.NewDecoder(res.Body).Decode(m); err != nil {
      return nil, err
   }
   if m.Message != "" {
      return nil, errors.New(m.Message)
   }
   m.URL = strings.Replace(m.URL, "[manifestType]", manifest_type, 1)
   return m, nil
}

type Media struct {
   Message string
   URL string
}

var scope = []string{
   "https://rcmnb2cprod.onmicrosoft.com/84593b65-0ef6-4a72-891c-d351ddd50aab/subscriptions.write",
   "https://rcmnb2cprod.onmicrosoft.com/84593b65-0ef6-4a72-891c-d351ddd50aab/toutv-profiling",
   "openid",
}

