package youtube

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
   "os"
   "strings"
)

func (d Device_Code) String() string {
   var b strings.Builder
   b.WriteString("1. Go to\n")
   b.WriteString(d.Verification_URL)
   b.WriteString("\n\n2. Enter this code\n")
   b.WriteString(d.User_Code)
   b.WriteString("\n\n3. Press Enter to continue")
   return b.String()
}

type Device_Code struct {
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
   t := new(Token)
   if err := json.NewDecoder(file).Decode(t); err != nil {
      return nil, err
   }
   return t, nil
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

func New_Device_Code() (*Device_Code, error) {
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
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   code := new(Device_Code)
   if err := json.NewDecoder(res.Body).Decode(code); err != nil {
      return nil, err
   }
   return code, nil
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
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(t)
}

func (d Device_Code) Token() (*Token, error) {
   body := url.Values{
      "client_id": {client_ID},
      "client_secret": {client_secret},
      "device_code": {d.Device_Code},
      "grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://oauth2.googleapis.com/token", strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   t := new(Token)
   if err := json.NewDecoder(res.Body).Decode(t); err != nil {
      return nil, err
   }
   return t, nil
}
