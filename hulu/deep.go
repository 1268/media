package hulu

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
)

type DeepLink struct {
   EabId string `json:"eab_id"`
}

func (a *Authenticate) DeepLink(id *EntityId) (*DeepLink, error) {
   req, _ := http.NewRequest("", "https://discover.hulu.com", nil)
   req.URL.Path = "/content/v5/deeplink/playback"
   req.URL.RawQuery = url.Values{
      "id":        {id.s},
      "namespace": {"entity"},
   }.Encode()
   req.Header.Set("authorization", "Bearer "+a.Data.UserToken)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var deep DeepLink
   err = json.NewDecoder(resp.Body).Decode(&deep)
   if err != nil {
      return nil, err
   }
   if deep.EabId == "" {
      return nil, errors.New("eab_id")
   }
   return &deep, nil
}
