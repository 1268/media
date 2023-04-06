package nbc

import (
   "2a.pages.dev/rosso/http"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/hex"
   "encoding/json"
   "io"
   "strconv"
   "strings"
   "time"
)

func New_Metadata(guid int64) (*Metadata, error) {
   var p page_request
   p.Query = graphQL_compact(query)
   p.Variables.App = "nbc"
   p.Variables.Name = strconv.FormatInt(guid, 10)
   p.Variables.One_App = true
   p.Variables.Platform = "android"
   p.Variables.Type = "VIDEO"
   body, err := json.MarshalIndent(p, "", " ")
   if err != nil {
      return nil, err
   }
   req := http.Post()
   req.Body_Bytes(body)
   req.Header.Set("Content-Type", "application/json")
   req.URL.Host = "friendship.nbc.co"
   req.URL.Path = "/v2/graphql"
   req.URL.Scheme = "https"
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

func (m Metadata) Name() string {
   var b strings.Builder
   b.WriteString(m.Series_Short_Title)
   b.WriteByte('-')
   b.WriteString(m.Secondary_Title)
   return b.String()
}

func (m Metadata) Video() (*Video, error) {
   var v video_request
   v.Device = "android"
   v.Device_ID = "android"
   v.External_Advertiser_ID = "NBC"
   v.MPX.Account_ID = m.MPX_Account_ID
   body, err := json.MarshalIndent(v, "", " ")
   if err != nil {
      return nil, err
   }
   req := http.Post()
   req.Body_Bytes(body)
   req.Header = http.Header{
      "Authorization": {authorization()},
      "Content-Type": {"application/json"},
   }
   req.URL.Host = "access-cloudpath.media.nbcuni.com"
   req.URL.Path = "/access/vod/nbcuniversal/" + m.MPX_GUID
   req.URL.Scheme = "http"
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
var secret_key = []byte("2b84a073ede61c766e4c0b3f1e656f7f")

func authorization() string {
   now := strconv.FormatInt(time.Now().UnixMilli(), 10)
   b := new(strings.Builder)
   b.WriteString("NBC-Security key=android_nbcuniversal,version=2.4")
   b.WriteString(",time=")
   b.WriteString(now)
   b.WriteString(",hash=")
   mac := hmac.New(sha256.New, secret_key)
   io.WriteString(mac, now)
   hex.NewEncoder(b).Write(mac.Sum(nil))
   return b.String()
}

type Metadata struct {
   MPX_Account_ID string `json:"mpxAccountId"`
   MPX_GUID string `json:"mpxGuid"`
   Series_Short_Title string `json:"seriesShortTitle"`
   Secondary_Title string `json:"secondaryTitle"`
}
