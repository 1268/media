package paramount

type Data struct {
   Items struct {
      Item []struct {
         Asset_Type string `xml:"assetType"`
      } `xml:"item"`
   } `xml:"items"`
}
