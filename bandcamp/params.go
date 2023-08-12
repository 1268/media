package bandcamp

import (
   "154.pages.dev/encoding/xml"
   "encoding/json"
   "io"
   "net/http"
)

func New_Params(ref string) (*Params, error) {
   res, err := http.Get(ref)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   _, text = xml.Cut(text, nil, []byte(`<p id="report-account-vm"`))
   var p struct {
      Report_Params []byte `xml:"data-tou-report-params,attr"`
   }
   if err := xml.Unmarshal(text, &p); err != nil {
      return nil, err
   }
   param := new(Params)
   if err := json.Unmarshal(p.Report_Params, param); err != nil {
      return nil, err
   }
   return param, nil
}

type Params struct {
   A_ID int
   I_ID int
   I_Type string
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
