package http

import (
   "bufio"
   "bytes"
   "log"
   "net/http"
)

// go.dev/wiki/CompilerOptimizations#map-lookup-by-byte
type transport map[string][]byte

func (t transport) RoundTrip(req *http.Request) (*http.Response, error) {
   var buf bytes.Buffer
   req.Write(&buf)
   key := buf.String()
   data, ok := t[key]
   if ok {
      log.Print("get")
   } else {
      log.Print("set")
      req1, err := http.ReadRequest(bufio.NewReader(&buf))
      if err != nil {
         return nil, err
      }
      req.Body = req1.Body
      resp, err := http.DefaultTransport.RoundTrip(req)
      if err != nil {
         return nil, err
      }
      resp.Write(&buf)
      data = buf.Bytes()
      t[key] = data
   }
   return http.ReadResponse(bufio.NewReader(bytes.NewReader(data)), nil)
}
