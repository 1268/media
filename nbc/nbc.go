package nbc

import "strings"

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
