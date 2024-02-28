package peacock

import (
   "bytes"
   "crypto/hmac"
   "crypto/md5"
   "crypto/sha1"
   "encoding/base64"
   "encoding/json"
   "errors"
   "fmt"
   "log/slog"
   "net/http"
   "net/url"
   "slices"
   "strings"
   "time"
)

func sign(method, path string, head http.Header, body []byte) string {
   timestamp := time.Now().Unix()
   text_headers := func() string {
      var s []string
      for k := range head {
         k = strings.ToLower(k)
         if strings.HasPrefix(k, "x-skyott-") {
            s = append(s, k + ": " + head.Get(k) + "\n")
         }
      }
      slices.Sort(s)
      return strings.Join(s, "")
   }()
   headers_md5 := md5.Sum([]byte(text_headers))
   payload_md5 := md5.Sum(body)
   signature := func() string {
      h := hmac.New(sha1.New, []byte(sky_key))
      fmt.Fprintln(h, method)
      fmt.Fprintln(h, path)
      fmt.Fprintln(h)
      fmt.Fprintln(h, sky_client)
      fmt.Fprintln(h, sky_version)
      fmt.Fprintf(h, "%x\n", headers_md5)
      fmt.Fprintln(h, timestamp)
      fmt.Fprintf(h, "%x\n", payload_md5)
      hashed := h.Sum(nil)
      return base64.StdEncoding.EncodeToString(hashed[:])
   }
   sky_ott := func() string {
      b := []byte("SkyOTT")
      // must be quoted
      b = fmt.Appendf(b, " client=%q", sky_client)
      // must be quoted
      b = fmt.Appendf(b, ",signature=%q", signature())
      // must be quoted
      b = fmt.Appendf(b, `,timestamp="%v"`, timestamp)
      // must be quoted
      b = fmt.Appendf(b, ",version=%q", sky_version)
      return string(b)
   }
   return sky_ott()
}

const (
   sky_client = "NBCU-ANDROID-v3"
   sky_key = "JuLQgyFz9n89D9pxcN6ZWZXKWfgj2PNBUb32zybj"
   sky_version = "1.0"
)

type SignIn struct {
   cookie *http.Cookie
}

func (s SignIn) Marshal() ([]byte, error) {
   return json.Marshal(s.cookie)
}

var Territory = "US"

func (s *SignIn) New(user, password string) error {
   slog.Debug("user", "identifier", user, "password", password)
   body := url.Values{
      "userIdentifier": {user},
      "password": {password},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://rango.id.peacocktv.com/signin/service/international",
      strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Content-Type": {"application/x-www-form-urlencoded"},
      "X-Skyott-Proposition": {"NBCUOTT"},
      "X-Skyott-Provider": {"NBCU"},
      "X-Skyott-Territory": {Territory},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusCreated {
      var b strings.Builder
      res.Write(&b)
      return errors.New(b.String())
   }
   for _, cookie := range res.Cookies() {
      if cookie.Name == "idsession" {
         s.cookie = cookie
         return nil
      }
   }
   return http.ErrNoCookie
}

// userToken is good for one day
type AuthToken struct {
   UserToken string
}

func (s SignIn) Auth() (*AuthToken, error) {
   var v struct {
      Auth struct {
         AuthScheme string `json:"authScheme"`
         Proposition string `json:"proposition"`
         Provider string `json:"provider"`
         ProviderTerritory string `json:"providerTerritory"`
      } `json:"auth"`
      Device struct {
         DrmDeviceId string `json:"drmDeviceId"`
         ID string `json:"id"`
         Platform string `json:"platform"`
         Type string `json:"type"`
      } `json:"device"`
   }
   v.Auth.AuthScheme = "MESSO"
   v.Auth.Proposition = "NBCUOTT"
   v.Auth.Provider = "NBCU"
   v.Auth.ProviderTerritory = Territory
   // if empty /drm/widevine/acquirelicense will fail with
   // {
   //    "errorCode": "OVP_00306",
   //    "description": "Security failure"
   // }
   v.Device.DrmDeviceId = "UNKNOWN"
   // if incorrect /video/playouts/vod will fail with
   // {
   //    "errorCode": "OVP_00311",
   //    "description": "Unknown deviceId"
   // }
   // changing this too often will result in a four hour block
   // {
   //    "errorCode": "OVP_00014",
   //    "description": "Maximum number of streaming devices exceeded"
   // }
   v.Device.ID = "PC"
   v.Device.Platform = "ANDROIDTV"
   v.Device.Type = "TV"
   body, err := json.Marshal(v)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://ovp.peacocktv.com/auth/tokens", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(s.cookie)
   req.Header.Set("content-type", "application/vnd.tokens.v1+json")
   req.Header.Set("x-sky-signature", sign(req.Method, req.URL.Path, nil, body))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   auth := new(AuthToken)
   if err := json.NewDecoder(res.Body).Decode(auth); err != nil {
      return nil, err
   }
   return auth, nil
}

func (s *SignIn) Unmarshal(b []byte) error {
   return json.Unmarshal(b, &s.cookie)
}
