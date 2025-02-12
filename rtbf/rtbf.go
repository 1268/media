package rtbf

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (a *AuvioLogin) Unmarshal(data []byte) error {
   err := json.Unmarshal(data, a)
   if err != nil {
      return err
   }
   if a.ErrorMessage != "" {
      return errors.New(a.ErrorMessage)
   }
   return nil
}

func (AuvioLogin) Marshal(id, password string) ([]byte, error) {
   resp, err := http.PostForm(
      "https://login.auvio.rtbf.be/accounts.login", url.Values{
         "APIKey":   {api_key},
         "loginID":  {id},
         "password": {password},
      },
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

func (a Address) Page() (*AuvioPage, error) {
   resp, err := http.Get(
      "https://bff-service.rtbf.be/auvio/v1.23/pages" + a.s,
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var value struct {
      Data struct {
         Content AuvioPage
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.Data.Content, nil
}

func (a *AuvioAuth) Entitlement(asset_id string) (*Entitlement, error) {
   req, _ := http.NewRequest("", "https://exposure.api.redbee.live", nil)
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v2/customer/RTBF/businessunit/Auvio/entitlement/")
      b.WriteString(asset_id)
      b.WriteString("/play")
      return b.String()
   }()
   req.Header = http.Header{
      "authorization":   {"Bearer " + a.SessionToken},
      "x-forwarded-for": {"91.90.123.17"},
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
   title := &Entitlement{}
   err = json.NewDecoder(resp.Body).Decode(title)
   if err != nil {
      return nil, err
   }
   return title, nil
}

func (e *Entitlement) Dash() (string, bool) {
   for _, format := range e.Formats {
      if format.Format == "DASH" {
         return format.MediaLocator, true
      }
   }
   return "", false
}

func (e *Entitlement) Wrap(data []byte) ([]byte, error) {
   req, err := http.NewRequest(
      "POST", "https://rbm-rtbf.live.ott.irdeto.com", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/licenseServer/widevine/v1/rbm-rtbf/license"
   req.URL.RawQuery = url.Values{
      "contentId":  {e.AssetId},
      "ls_session": {e.PlayToken},
   }.Encode()
   req.Header.Set("content-type", "application/x-protobuf")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

func (a *Address) Set(data string) error {
   a.s = strings.TrimPrefix(data, "https://")
   a.s = strings.TrimPrefix(a.s, "auvio.rtbf.be")
   return nil
}

func (a *Address) String() string {
   return a.s
}

type Address struct {
   s string
}

func (a *AuvioPage) GetAssetId() (string, bool) {
   if a.AssetId != "" {
      return a.AssetId, true
   }
   if a.Media != nil {
      return a.Media.AssetId, true
   }
   return "", false
}

// hard coded in JavaScript
const api_key = "4_Ml_fJ47GnBAW6FrPzMxh0w"

func (a *AuvioLogin) Token() (*WebToken, error) {
   resp, err := http.PostForm(
      "https://login.auvio.rtbf.be/accounts.getJWT", url.Values{
         "APIKey":      {api_key},
         "login_token": {a.SessionInfo.CookieValue},
      },
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var web WebToken
   err = json.NewDecoder(resp.Body).Decode(&web)
   if err != nil {
      return nil, err
   }
   if web.ErrorMessage != "" {
      return nil, errors.New(web.ErrorMessage)
   }
   return &web, nil
}

func (w *WebToken) Auth() (*AuvioAuth, error) {
   value := map[string]any{
      "device": map[string]string{
         "deviceId": "",
         "type":     "WEB",
      },
      "jwt": w.IdToken,
   }
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://exposure.api.redbee.live", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/v2/customer/RTBF/businessunit/Auvio/auth/gigyaLogin"
   req.Header.Set("content-type", "application/json")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   auth := &AuvioAuth{}
   err = json.NewDecoder(resp.Body).Decode(auth)
   if err != nil {
      return nil, err
   }
   return auth, nil
}

///

type AuvioAuth struct {
   SessionToken string
}

type AuvioLogin struct {
   ErrorMessage string
   SessionInfo  struct {
      CookieValue string
   }
}

type AuvioPage struct {
   AssetId string
   Media   *struct {
      AssetId string
   }
}

type Entitlement struct {
   AssetId   string
   PlayToken string
   Formats   []struct {
      Format       string
      MediaLocator string
   }
}

type WebToken struct {
   ErrorMessage string
   IdToken      string `json:"id_token"`
}
