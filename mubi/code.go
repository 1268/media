package mubi

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

func (*LinkCode) Marshal() ([]byte, error) {
   req, err := http.NewRequest("", "https://api.mubi.com/v3/link_code", nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "client": {client},
      "client-country": {ClientCountry},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   return io.ReadAll(resp.Body)
}

var ClientCountry = "US"

// "android" requires headers:
// client-device-identifier
// client-version
const client = "web"

type Address struct {
   Data string
}

func (a *Address) Set(text string) error {
   var ok bool
   _, a.Data, ok = strings.Cut(text, "/films/")
   if !ok {
      return errors.New("/films/")
   }
   return nil
}

type TextTrack struct {
   Id string
   Url string
}

type FilmResponse struct {
   Id int64
   Title string
   Year int
}

func (a *Address) String() string {
   return a.Data
}

type Namer struct {
   Film *FilmResponse
}

func (n Namer) Title() string {
   return n.Film.Title
}

func (n Namer) Year() int {
   return n.Film.Year
}

func (t *TextTrack) String() string {
   return "id = " + t.Id
}

func (a *Address) Film() (*FilmResponse, error) {
   req, err := http.NewRequest("", "https://api.mubi.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v3/films/" + a.Data
   req.Header = http.Header{
      "client": {client},
      "client-country": {ClientCountry},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   film := &FilmResponse{}
   err = json.NewDecoder(resp.Body).Decode(film)
   if err != nil {
      return nil, err
   }
   return film, nil
}

func (c *LinkCode) String() string {
   var b strings.Builder
   b.WriteString("TO LOG IN AND START WATCHING\n")
   b.WriteString("Go to\n")
   b.WriteString("mubi.com/android\n")
   b.WriteString("and enter the code below\n")
   b.WriteString(c.LinkCode)
   return b.String()
}

func (Namer) Episode() int {
   return 0
}

func (Namer) Season() int {
   return 0
}

func (Namer) Show() string {
   return ""
}

type LinkCode struct {
   AuthToken string `json:"auth_token"`
   LinkCode string `json:"link_code"`
}

func (c *LinkCode) Unmarshal(data []byte) error {
   return json.Unmarshal(data, c)
}
