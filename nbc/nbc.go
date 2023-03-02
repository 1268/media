package nbc

import (
   "2a.pages.dev/rosso/http"
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/hex"
   "encoding/json"
   "io"
   "strconv"
   "strings"
   "time"
)

const query = `
query bonanzaPage(
   $app: NBCUBrands!
   $name: String!
   $oneApp: Boolean
   $platform: SupportedPlatforms!
   $type: EntityPageType!
   $userId: String!
) {
   bonanzaPage(
      app: $app
      name: $name
      oneApp: $oneApp
      platform: $platform
      type: $type
      userId: $userId
   ) {
      metadata {
         ... on VideoPageData {
            mpxAccountId
            mpxGuid
            secondaryTitle
            seriesShortTitle
         }
      }
   }
}
`

var (
   Client = http.Default_Client
   secret_key = []byte("2b84a073ede61c766e4c0b3f1e656f7f")
)

func authorization() string {
   now := strconv.FormatInt(time.Now().UnixMilli(), 10)
   str := new(strings.Builder)
   str.WriteString("NBC-Security key=android_nbcuniversal,version=2.4")
   str.WriteString(",time=")
   str.WriteString(now)
   str.WriteString(",hash=")
   mac := hmac.New(sha256.New, secret_key)
   io.WriteString(mac, now)
   hex.NewEncoder(str).Write(mac.Sum(nil))
   return str.String()
}

func graphQL_compact(s string) string {
   old_new := []string{
      "\n", "",
      strings.Repeat(" ", 12), " ",
      strings.Repeat(" ", 9), " ",
      strings.Repeat(" ", 6), " ",
      strings.Repeat(" ", 3), " ",
   }
   return strings.NewReplacer(old_new...).Replace(s)
}

type Metadata struct {
   MPX_Account_ID string `json:"mpxAccountId"`
   MPX_GUID string `json:"mpxGuid"`
   Series_Short_Title string `json:"seriesShortTitle"`
   Secondary_Title string `json:"secondaryTitle"`
}

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
   req, err := http.NewRequest(
      "POST", "https://friendship.nbc.co/v2/graphql", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/json")
   res, err := Client.Do(req)
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
   var str strings.Builder
   str.WriteString("http://access-cloudpath.media.nbcuni.com")
   str.WriteString("/access/vod/nbcuniversal/")
   str.WriteString(m.MPX_GUID)
   req, err := http.NewRequest("POST", str.String(), bytes.NewReader(body))
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {authorization()},
      "Content-Type": {"application/json"},
   }
   res, err := Client.Do(req)
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

type Video struct {
   // this is only valid for one minute
   Manifest_Path string `json:"manifestPath"`
}

type page_request struct {
   Query string `json:"query"`
   Variables struct {
      App string `json:"app"` // String cannot represent a non string value
      Name string `json:"name"`
      One_App bool `json:"oneApp"`
      Platform string `json:"platform"`
      Type string `json:"type"` // can be empty
      User_ID string `json:"userId"`
   } `json:"variables"`
}

type video_request struct {
   Device string `json:"device"`
   Device_ID string `json:"deviceId"`
   External_Advertiser_ID string `json:"externalAdvertiserId"`
   MPX struct {
      Account_ID string `json:"accountId"`
   } `json:"mpx"`
}
