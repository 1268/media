package cbc

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
   "strings"
)

type profile struct {
   Claims_Token string `json:"claimsToken"`
}

func (t token) profile() (*profile, error) {
   req := http.Get()
   req.URL.Host = "services.radio-canada.ca"
   req.URL.Path = "/ott/subscription/v2/gem/Subscriber/profile"
   req.URL.Scheme = "https"
   req.URL.RawQuery = "device=phone_android"
   req.Header.Set("Authorization", "Bearer " + t.Access_Token)
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   pro := new(profile)
   if err := json.NewDecoder(res.Body).Decode(pro); err != nil {
      return nil, err
   }
   return pro, nil
}

type token struct {
   Access_Token string
}

func new_token(username, password string) (*token, error) {
   body := url.Values{
      "client_id": {"7f44c935-6542-4ce7-ae05-eb887809741c"},
      "grant_type": {"password"},
      "password": {password},
      "scope": {strings.Join([]string{
         "https://rcmnb2cprod.onmicrosoft.com/84593b65-0ef6-4a72-891c-d351ddd50aab/subscriptions.write",
         "https://rcmnb2cprod.onmicrosoft.com/84593b65-0ef6-4a72-891c-d351ddd50aab/toutv-profiling",
         "openid",
      }, " ")},
      "username": {username},
   }.Encode()
   req := http.Post()
   req.URL.Scheme = "https"
   req.URL.Host = "rcmnb2cprod.b2clogin.com"
   req.URL.Path = "/rcmnb2cprod.onmicrosoft.com/B2C_1A_ExternalClient_ROPC_Auth/oauth2/v2.0/token"
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.Body_String(body)
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tok := new(token)
   if err := json.NewDecoder(res.Body).Decode(tok); err != nil {
      return nil, err
   }
   return tok, nil
}
