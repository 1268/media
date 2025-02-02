package tv

import (
   "encoding/xml"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

const (
   out_of_country = "/out-of-country"
   verification_token = "__RequestVerificationToken"
)

type get_login struct {
   cookies []*http.Cookie
   section struct {
      Div     struct {
         Form struct {
            Input []struct {
               Name  string `xml:"name,attr"`
               Value string `xml:"value,attr"`
            } `xml:"input"`
         } `xml:"form"`
      } `xml:"div"`
   }
}

func (s *get_login) cookie_token() (*http.Cookie, error) {
   for _, cookie := range s.cookies {
      if cookie.Name == verification_token {
         return cookie, nil
      }
   }
   return nil, http.ErrNoCookie
}

func (s *get_login) input_token() (string, error) {
   for _, input := range s.section.Div.Form.Input {
      if input.Name == verification_token {
         return input.Value, nil
      }
   }
   return "", errors.New(verification_token)
}

// hard geo block
func (s *get_login) New() error {
   req, _ := http.NewRequest("", "https://show.sky.ch/de/login", nil)
   req.Header.Set("tv", "Emulator")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   if strings.HasSuffix(resp.Request.URL.Path, out_of_country) {
      return errors.New(out_of_country)
   }
   err = xml.NewDecoder(resp.Body).Decode(&s.section)
   if err != nil {
      return err
   }
   s.cookies = resp.Cookies()
   return nil
}

type post_login []*http.Cookie

///

// hard geo block
func (s *get_login) login(username, password string) (post_login, error) {
   input_token, err := s.input_token()
   if err != nil {
      return nil, err
   }
   data := url.Values{
      "password": {password},
      "username": {username},
      verification_token: {input_token},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://show.sky.ch/de/Authentication/Login",
      strings.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("content-type", "application/x-www-form-urlencoded")
   req.Header.Set("tv", "Emulator")
   cookie_token, err := s.cookie_token()
   if err != nil {
      return nil, err
   }
   req.AddCookie(cookie_token)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   _, err = io.Copy(io.Discard, resp.Body)
   if err != nil {
      return nil, err
   }
   return resp.Cookies(), nil
}

//sky-auth-token

