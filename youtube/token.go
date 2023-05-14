package youtube

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
   "os"
   "strings"
)

func (t *Token) Refresh() error {
   body := url.Values{
      "client_id": {client_ID},
      "client_secret": {client_secret},
      "grant_type": {"refresh_token"},
      "refresh_token": {t.Refresh_Token},
   }.Encode()
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "oauth2.googleapis.com",
      Path: "/token",
   })
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.Body_String(body)
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
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "oauth2.googleapis.com",
      Path: "/token",
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

func Read_Token(name string) (*Token, error) {
   data, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   tok := new(Token)
   if err := json.Unmarshal(data, tok); err != nil {
      return nil, err
   }
   return tok, nil
}

func (t Token) Write_File(name string) error {
   data, err := json.Marshal(t)
   if err != nil {
      return err
   }
   return os.WriteFile(name, data, 0666)
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
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "oauth2.googleapis.com",
      Path: "/device/code",
   })
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.Body_String(body)
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

