package max

import (
   "net/http"
   "net/url"
)

func video(ref string) (*http.Response, error) {
   return http.Get(ref)
}

func media() (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "medium.ngtv.io"
   req.URL.Path = "/v2/media/meb52c30a61e34b63a0ca66946f1b515e5aaad5f9d/desktop"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["appId"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuZXR3b3JrIjoiaGJvbWF4IiwicHJvZHVjdCI6ImJlYW0iLCJwbGF0Zm9ybSI6IndlYi10b3AyIiwiYXBwSWQiOiJoYm9tYXgtYmVhbS13ZWItdG9wMi1wMnFiMnAifQ.8XAm_qxXSa7REssRiJsrAnO0eh_Ljs24muIfZyaLE-8"}
   req.URL.RawQuery = val.Encode()
   return new(http.Transport).RoundTrip(&req)
}
