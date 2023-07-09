package nbc

import (
   "crypto/hmac"
   "crypto/sha256"
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func (m Metadata) Title() string {
   return m.Secondary_Title
}

func (m Metadata) Series() string {
   return m.Series_Short_Title
}

func (m Metadata) Season() (int64, error) {
   return m.Season_Number, nil
}

func (m Metadata) Episode() (int64, error) {
   return m.Episode_Number, nil
}

func (m Metadata) Date() (time.Time, error) {
   return time.Parse(time.RFC3339, m.Air_Date)
}

type Metadata struct {
   MPX_Account_ID string `json:"mpxAccountId"`
   MPX_GUID string `json:"mpxGuid"`
   Series_Short_Title string `json:"seriesShortTitle"`
   Season_Number int64 `json:"seasonNumber,string"`
   Episode_Number int64 `json:"episodeNumber,string"`
   Secondary_Title string `json:"secondaryTitle"`
   Air_Date string `json:"airDate"`
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
            Metadata *Metadata
         } `json:"bonanzaPage"`
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&page); err != nil {
      return nil, err
   }
   if page.Data.Bonanza_Page.Metadata == nil {
      return nil, fmt.Errorf(".data.bonanzaPage.metadata is null")
   }
   return page.Data.Bonanza_Page.Metadata, nil
}

var secret_key = []byte("2b84a073ede61c766e4c0b3f1e656f7f")

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
            airDate
            episodeNumber
            mpxAccountId
            mpxGuid
            seasonNumber
            secondaryTitle
            seriesShortTitle
         }
      }
   }
}
`

// this is better than strings.Replace and strings.ReplaceAll
func graphQL_compact(s string) string {
   f := strings.Fields(s)
   return strings.Join(f, " ")
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
