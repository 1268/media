package cbc

import "2a.pages.dev/rosso/http"

/*
content	
0	
lineups	
1	
items	
3	
idMedia	958273
*/
func gem(s string) (*http.Response, error) {
   req := http.Get()
   req.URL.Scheme = "https"
   req.URL.Host = "services.radio-canada.ca"
   req.URL.Path = "/ott/catalog/v2/gem/show/" + s
   // you can also use `web` here, but this is smaller
   req.URL.RawQuery = "device=phone_android"
   return http.Default_Client.Do(req)
}
