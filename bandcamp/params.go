package bandcamp

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/xml"
   "encoding/json"
   "io"
)

var Client = http.Default_Client

type Params struct {
   A_ID int
   I_ID int
   I_Type string
}

func New_Params(ref string) (*Params, error) {
   res, err := Client.Get(ref)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var scan xml.Scanner
   scan.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   scan.Sep = []byte(`<p id="report-account-vm"`)
   scan.Scan()
   var p struct {
      Report_Params []byte `xml:"data-tou-report-params,attr"`
   }
   if err := scan.Decode(&p); err != nil {
      return nil, err
   }
   param := new(Params)
   if err := json.Unmarshal(p.Report_Params, param); err != nil {
      return nil, err
   }
   return param, nil
}

func (p Params) Band() (*Band, error) {
   return new_band(p.A_ID)
}

func (p Params) Tralbum() (*Tralbum, error) {
   switch p.I_Type {
   case "a":
      return new_tralbum('a', p.I_ID)
   case "t":
      return new_tralbum('t', p.I_ID)
   }
   return nil, invalid_type{p.I_Type}
}
