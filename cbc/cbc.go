package cbc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "net/url"
   "os"
   "strings"
)

func (p Profile) Media(item *lineup_item) (*Media, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "services.radio-canada.ca",
      Path: "/media/validation/v2",
      RawQuery: url.Values{
         "appCode": {"gem"},
         "idMedia": {item.Formatted_ID_Media},
         "manifestType": {"desktop"},
         "output": {"json"},
      }.Encode(),
   })
   req.Header = http.Header{
      "X-Claims-Token": {p.Claims_Token},
      "X-Forwarded-For": {forwarded_for},
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   med := new(Media)
   if err := json.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   if med.Message != nil {
      return nil, errors.New(*med.Message)
   }
   return med, nil
}

var scope = []string{
   "https://rcmnb2cprod.onmicrosoft.com/84593b65-0ef6-4a72-891c-d351ddd50aab/subscriptions.write",
   "https://rcmnb2cprod.onmicrosoft.com/84593b65-0ef6-4a72-891c-d351ddd50aab/toutv-profiling",
   "openid",
}

func New_Token(username, password string) (*Token, error) {
   body := url.Values{
      "client_id": {"7f44c935-6542-4ce7-ae05-eb887809741c"},
      "grant_type": {"password"},
      "password": {password},
      "scope": {strings.Join(scope, " ")},
      "username": {username},
   }.Encode()
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "rcmnb2cprod.b2clogin.com",
      Path: "/rcmnb2cprod.onmicrosoft.com/B2C_1A_ExternalClient_ROPC_Auth/oauth2/v2.0/token",
   })
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.Body_String(body)
   res, err := http.Default_Client.Do(req)
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

// gem.cbc.ca/media/downton-abbey/s01e05
func Get_ID(input string) string {
   _, after, found := strings.Cut(input, "/media/")
   if found {
      return after
   }
   return input
}

type Media struct {
   Message *string
   URL *string
}

func (p Profile) Write_File(name string) error {
   data, err := json.MarshalIndent(p, "", " ")
   if err != nil {
      return err
   }
   return os.WriteFile(name, data, 0666)
}

func Read_Profile(name string) (*Profile, error) {
   data, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   pro := new(Profile)
   if err := json.Unmarshal(data, pro); err != nil {
      return nil, err
   }
   return pro, nil
}

type Profile struct {
   Claims_Token string `json:"claimsToken"`
}

func (t Token) Profile() (*Profile, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "services.radio-canada.ca",
      Path: "/ott/subscription/v2/gem/Subscriber/profile",
      RawQuery: "device=phone_android",
   })
   req.Header.Set("Authorization", "Bearer " + t.Access_Token)
   res, err := http.Default_Client.Do(req)
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

