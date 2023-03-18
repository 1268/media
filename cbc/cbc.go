package cbc

import (
   "2a.pages.dev/rosso/http"
   "bytes"
   "encoding/json"
   "errors"
   "net/url"
   "os"
   "strings"
)

const forwarded_for = "99.224.0.0"

var Client = http.Default_Client

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

func (p Profile) Media(a *Asset) (*Media, error) {
   req, err := http.NewRequest("GET", a.Play_Session.URL, nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "X-Claims-Token": {p.Claims_Token},
      "X-Forwarded-For": {forwarded_for},
   }
   res, err := Client.Do(req)
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

type Asset struct {
   Apple_Content_ID string `json:"appleContentId"`
   Play_Session struct {
      URL string
   } `json:"playSession"`
}

func New_Asset(id string) (*Asset, error) {
   var b strings.Builder
   b.WriteString("https://services.radio-canada.ca/ott/cbc-api/v2/assets/")
   b.WriteString(id)
   res, err := Client.Get(b.String())
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   a := new(Asset)
   if err := json.NewDecoder(res.Body).Decode(a); err != nil {
      return nil, err
   }
   return a, nil
}

func Open_Profile(name string) (*Profile, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   pro := new(Profile)
   if err := json.NewDecoder(file).Decode(pro); err != nil {
      return nil, err
   }
   return pro, nil
}

func (p Profile) Create(name string) error {
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   return json.NewEncoder(file).Encode(p)
}

const api_key = "3f4beddd-2061-49b0-ae80-6f1f2ed65b37"

type Login struct {
   Access_Token string
   Expires_In string
}

func (l Login) Web_Token() (*Web_Token, error) {
   req, err := http.NewRequest(
      "GET", "https://cloud-api.loginradius.com/sso/jwt/api/token", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "access_token": {l.Access_Token},
      "apikey": {api_key},
      "jwtapp": {"jwt"},
   }.Encode()
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   web := new(Web_Token)
   if err := json.NewDecoder(res.Body).Decode(web); err != nil {
      return nil, err
   }
   return web, nil
}

type Over_The_Top struct {
   Access_Token string `json:"accessToken"`
}

func (o Over_The_Top) Profile() (*Profile, error) {
   req, err := http.NewRequest(
      "GET", "https://services.radio-canada.ca/ott/cbc-api/v2/profile", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("OTT-Access-Token", o.Access_Token)
   res, err := Client.Do(req)
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

type Profile struct {
   Tier string
   Claims_Token string `json:"claimsToken"`
}

type Web_Token struct {
   Signature string
}

func New_Login(email, password string) (*Login, error) {
   auth := map[string]string{
      "email": email,
      "password": password,
   }
   raw_auth, err := json.MarshalIndent(auth, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.loginradius.com/identity/v2/auth/login",
      bytes.NewReader(raw_auth),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   req.URL.RawQuery = "apiKey=" + api_key
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   login := new(Login)
   if err := json.NewDecoder(res.Body).Decode(login); err != nil {
      return nil, err
   }
   return login, nil
}

func (w Web_Token) Over_The_Top() (*Over_The_Top, error) {
   token := map[string]string{"jwt": w.Signature}
   raw_token, err := json.MarshalIndent(token, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://services.radio-canada.ca/ott/cbc-api/v2/token",
      bytes.NewReader(raw_token),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   top := new(Over_The_Top)
   if err := json.NewDecoder(res.Body).Decode(top); err != nil {
      return nil, err
   }
   return top, nil
}
