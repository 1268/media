package main

import (
   "net/http"
   "net/url"
   "os"
   "bytes"
   "encoding/json"
)

/*
x-forwarded-for fail
mullvad.net fail
proxy-seller.com pass
*/
func main() {
   var req http.Request
   req.Header = http.Header{}
   req.URL = &url.URL{}
   req.URL.Host = "show.sky.ch"
   req.URL.Path = "/de/SkyPlayerAjax/SkyPlayer"
   req.URL.Scheme = "https"
   value := url.Values{}
   value["id"] = []string{"2035"}
   value["contentType"] = []string{"2"}
   req.Header["X-Requested-With"] = []string{"XMLHttpRequest"}
   var dst bytes.Buffer
   err := json.Compact(&dst, []byte(data))
   if err != nil {
      panic(err)
   }
   req.Header["Cookie"] = []string{
      "sky-auth-token=" + dst.String(),
   }
   req.URL.RawQuery = value.Encode()
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

const data = `
{
   "cstId": 1938441,
   "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxOTM4NDQxIiwibWFjQWRkcmVzcyI6IjBlY2UwZTI1LTFkYzUtNDI4ZC1iNmNiLTlmYmU4MTdjNTgxNSIsInRpbWUiOiIyMDI1LDEsMjYsMjEsMCwwIiwiaXNzIjoiaHR0cHM6Ly93d3cuc2t5LmNoIiwiYXVkIjoiU2t5IFVzZXJzIn0.f4UHiZUopfFqmGE48qOmaYmJRH5y8MqwbADDhnfTBxY"
}
`
