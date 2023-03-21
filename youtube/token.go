package youtube

import (
   "encoding/json"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func (c Code) String() string {
   var b strings.Builder
   b.WriteString("1. Go to\n")
   b.WriteString(c.Verification_URL)
   b.WriteString("\n\n2. Enter this code\n")
   b.WriteString(c.User_Code)
   b.WriteString("\n\n3. Press Enter to continue")
   return b.String()
}

type Code struct {
   Device_Code string
   User_Code string
   Verification_URL string
}

func Open_Token(name string) (*Token, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   tok := new(Token)
   if err := json.NewDecoder(file).Decode(tok); err != nil {
      return nil, err
   }
   return tok, nil
}

func (t Token) Create(name string) error {
   file, err := os.Create(name)
   if err != nil {
      return err
   }
   defer file.Close()
   return json.NewEncoder(file).Encode(t)
}

type Token struct {
   Access_Token string
   Error string
   Refresh_Token string
}

func New_Code() (*Code, error) {
   body := url.Values{
      "client_id": {client_ID},
      "scope": {"https://www.googleapis.com/auth/youtube"},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://oauth2.googleapis.com/device/code",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := HTTP_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   c := new(Code)
   if err := json.NewDecoder(res.Body).Decode(c); err != nil {
      return nil, err
   }
   return c, nil
}

func (t *Token) Refresh() error {
   body := url.Values{
      "client_id": {client_ID},
      "client_secret": {client_secret},
      "grant_type": {"refresh_token"},
      "refresh_token": {t.Refresh_Token},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://oauth2.googleapis.com/token", strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := HTTP_Client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(t)
}

func (c Code) Token() (*Token, error) {
   body := url.Values{
      "client_id": {client_ID},
      "client_secret": {client_secret},
      "device_code": {c.Device_Code},
      "grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://oauth2.googleapis.com/token", strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := HTTP_Client.Do(req)
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
