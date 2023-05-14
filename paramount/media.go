package paramount

import (
   "2a.pages.dev/rosso/http"
   "net/url"
   "strconv"
)

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
   req := http.Get(&url.URL{
      Scheme: "http",
      Host: "link.theplatform.com",
      Path: url_path(nil),
      RawQuery: query.Encode(),
   })
   client := http.Default_Client
   client.Status = http.StatusFound
   res, err := client.Do(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   return res.Header.Get("Location"), nil
}

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

