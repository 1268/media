package tv

import (
   "encoding/xml"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (p *login_page) login_page() (*http.Response, error) {
   data = {
      'username': __login, 'password': __password,
      'returnUrl': 'Home/HomeTv', 'subscriptionUrl': '/de/subscription',
      cookie_check["rvtp"]: page_token,
   }
   return requests.post(
      auth_url, timeout=5, headers=login_headers,
      cookies=cookies,
      data=data,
      allow_redirects=False, verify=False
   )
}

func (p *login_page) page_token() (string, bool) {
   for _, input := range p.section.Div.Form.Input {
      if input.Name == "__RequestVerificationToken" {
         return input.Value, true
      }
   }
   return "", false
}

var cookie_check = map[string]string{
  "asp": "_ASP.NET_SessionId_",
  "cc": "SkyCheeseCake",
  "cc2": "SkyCake",
  "rvt": "__RequestVerificationToken",
  "rvtp": "__RequestVerificationToken",
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

func (p *login_page) New() error {
   req, _ := http.NewRequest("", "https://show.sky.ch/de/login", nil)
   req.URL.RawQuery = "forceClassicalTvLogin=True"
   req.Header = login_headers
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
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

const out_of_country = "/out-of-country"

var login_headers = http.Header{
   "accept-language": {"de"},
   "referer": {"https://show.sky.ch/de/tv/"},
   "tv": {"Emulator"},
   "user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:127.0) Gecko/20100101 Firefox/127.0"},
}

const auth_url = "https://show.sky.ch/de/Authentication/Login"

func (p login_page) asp_cookie() (*http.Cookie, bool) {
   for _, cookie := range p.cookies {
      if cookie.Name == "_ASP.NET_SessionId_" {
         return cookie, true
      }
   }
   return nil, false
}

func (p login_page) cookie_token() (*http.Cookie, bool) {
   for _, cookie := range p.cookies {
      if cookie.Name == "__RequestVerificationToken" {
         return cookie, true
      }
   }
   return nil, false
}

func (*login_page) sky_tv_device() http.Cookie {
   var cookie http.Cookie
   cookie.Name = "SkyTvDevice"
   cookie.Value = url.QueryEscape(`
   {
      "isSky": true,
      "keys": {
         "back": 461,
         "down": 40,
         "enter": 13,
         "ff": 417,
         "ff10": -1,
         "key0": -1,
         "key1": -1,
         "key2": -1,
         "key3": -1,
         "key4": -1,
         "key5": -1,
         "key6": -1,
         "key7": -1,
         "key8": -1,
         "key9": -1,
         "left": 37,
         "pause": 19,
         "play": 415,
         "playPause": -1,
         "rew": 412,
         "rew10": -1,
         "right": 39,
         "search": -1,
         "stop": 413,
         "up": 38
      },
      "type": {
         "code": "Desktop"
      },
      "year": ""
   }
   `)
   return cookie
}
