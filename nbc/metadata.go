package nbc

import (
   "2a.pages.dev/rosso/http"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/json"
   "fmt"
   "net/url"
   "time"
)

func (m Metadata) Name() string {
   var b []byte
   b = append(b, m.Series_Short_Title...)
   b = append(b, '-')
   b = append(b, m.Secondary_Title...)
   return string(b)
}

func New_Metadata(guid int64) (*Metadata, error) {
   body := func(r *http.Request) error {
      var p page_request
      p.Query = graphQL_compact(query)
      p.Variables.App = "nbc"
      p.Variables.Name = fmt.Sprint(guid)
      p.Variables.One_App = true
      p.Variables.Platform = "android"
      p.Variables.Type = "VIDEO"
      b, err := json.MarshalIndent(p, "", " ")
      if err != nil {
         return err
      }
      r.Body_Bytes(b)
      return nil
   }
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "friendship.nbc.co",
      Path: "/v2/graphql",
   })
   req.Header.Set("Content-Type", "application/json")
   err := body(req)
   if err != nil {
      return nil, err
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var page struct {
      Data struct {
         Bonanza_Page struct {
            Metadata Metadata
         } `json:"bonanzaPage"`
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&page); err != nil {
      return nil, err
   }
   return &page.Data.Bonanza_Page.Metadata, nil
}

var secret_key = []byte("2b84a073ede61c766e4c0b3f1e656f7f")

type Metadata struct {
   MPX_Account_ID string `json:"mpxAccountId"`
   MPX_GUID string `json:"mpxGuid"`
   Series_Short_Title string `json:"seriesShortTitle"`
   Secondary_Title string `json:"secondaryTitle"`
}
func authorization(b []byte) string {
   now := time.Now().UnixMilli()
   hash := hmac.New(sha256.New, secret_key)
   fmt.Fprint(hash, now)
   b = append(b, "NBC-Security key=android_nbcuniversal,version=2.4"...)
   b = append(b, ",time="...)
   b = fmt.Append(b, now)
   b = append(b, ",hash="...)
   b = fmt.Appendf(b, "%x", hash.Sum(nil))
   return string(b)
}

func (m Metadata) Video() (*Video, error) {
   body := func(r *http.Request) error {
      var v video_request
      v.Device = "android"
      v.Device_ID = "android"
      v.External_Advertiser_ID = "NBC"
      v.MPX.Account_ID = m.MPX_Account_ID
      b, err := json.MarshalIndent(v, "", " ")
      if err != nil {
         return err
      }
      r.Body_Bytes(b)
      return nil
   }
   req := http.Post(&url.URL{
      Scheme: "http",
      Host: "access-cloudpath.media.nbcuni.com",
      Path: "/access/vod/nbcuniversal/" + m.MPX_GUID,
   })
   req.Header = http.Header{
      "Authorization": {authorization(nil)},
      "Content-Type": {"application/json"},
   }
   err := body(req)
   if err != nil {
      return nil, err
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   vid := new(Video)
   if err := json.NewDecoder(res.Body).Decode(vid); err != nil {
      return nil, err
   }
   return vid, nil
}

