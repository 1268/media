package paramount

import (
   "2a.pages.dev/rosso/http"
   "net/url"
   "strconv"
)

const (
   aid = 2198311517
   cms_account_id = "dJ5BDC"
)

func DASH_CENC(content_ID string) (string, error) {
   query := url.Values{
      "assetTypes": {"DASH_CENC"},
      "formats": {"MPEG-DASH"},
   }
   return location(content_ID, query)
}

func Downloadable(content_ID string) (string, error) {
   query := url.Values{
      "assetTypes": {"Downloadable"},
      "formats": {"MPEG4"},
   }
   return location(content_ID, query)
}

func location(content_ID string, query url.Values) (string, error) {
   url_path := func(b []byte) string {
      b = append(b, "/s/"...)
      b = append(b, cms_account_id...)
      b = append(b, "/media/guid/"...)
      b = strconv.AppendInt(b, aid, 10)
      b = append(b, '/')
      b = append(b, content_ID...)
      return string(b)
   }
   client := http.Default_Client
   client.Status = http.StatusFound
   req := http.Get()
   req.URL.Scheme = "http"
   req.URL.Host = "link.theplatform.com"
   req.URL.Path = url_path(nil)
   req.URL.RawQuery = query.Encode()
   res, err := client.Do(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   return res.Header.Get("Location"), nil
}
