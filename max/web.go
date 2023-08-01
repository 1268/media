package max

import (
   "encoding/json"
   "errors"
   "net/http"
)

type media struct {
   Media struct {
      Desktop struct {
         Unprotected struct {
            Unencrypted struct {
               URL string
            }
         }
      }
   }
}

func (p page_data) media() (*media, error) {
   req, err := http.NewRequest("GET", "https://medium.ngtv.io/v2/media/", nil)
   if err != nil {
      return nil, err
   }
   video, err := p.video_URL()
   if err != nil {
      return nil, err
   }
   req.URL.Path += video.Value.Large + "/desktop"
   req.URL.RawQuery = "appId=" + p.Page.Media_App_ID
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   med := new(media)
   if err := json.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}

type page_data struct {
   Mappings []struct {
      Property string
      Value string
   }
   Page struct {
      Media_App_ID string
   }
}

type value struct {
   Value struct {
      Large string
   }
}

func (p page_data) video_URL() (*value, error) {
   for _, m := range p.Mappings {
      if m.Property == "videoUrl" {
         v := new(value)
         err := json.Unmarshal([]byte(m.Value), v)
         if err != nil {
            return nil, err
         }
         return v, nil
      }
   }
   return nil, errors.New("videoUrl not found")
}

type next_data struct {
   Props struct {
      Page_Props struct {
         Page_Data string `json:"pageData"`
      } `json:"pageProps"`
   }
}

func (n next_data) page_data() (*page_data, error) {
   page := new(page_data)
   err := json.Unmarshal([]byte(n.Props.Page_Props.Page_Data), page)
   if err != nil {
      return nil, err
   }
   return page, nil
}
