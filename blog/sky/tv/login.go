package tv

import (
   "encoding/xml"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (p *login_page) login_page(
   username, password string,
) (*http.Response, error) {
   page_token, err := p.page_token()
   if err != nil {
      return nil, err
   }
   data := url.Values{
      "__RequestVerificationToken": {page_token},
      "password": {password},
      "username": {username},
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
   cookie_token, err := p.cookie_token()
   if err != nil {
      return nil, err
   }
   req.AddCookie(cookie_token)
   return http.DefaultClient.Do(req)
}

func (p *login_page) cookie_token() (*http.Cookie, error) {
   for _, cookie := range p.cookies {
      if cookie.Name == "__RequestVerificationToken" {
         return cookie, nil
      }
   }
   return nil, http.ErrNoCookie
}

func (p *login_page) page_token() (string, error) {
   for _, input := range p.section.Div.Form.Input {
      if input.Name == "__RequestVerificationToken" {
         return input.Value, nil
      }
   }
   return "", http.ErrNoCookie
}

type login_page struct {
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

const out_of_country = "/out-of-country"

func (p *login_page) New() error {
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
   err = xml.NewDecoder(resp.Body).Decode(&p.section)
   if err != nil {
      return err
   }
   p.cookies = resp.Cookies()
   return nil
}
