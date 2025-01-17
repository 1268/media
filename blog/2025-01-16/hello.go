package main

import (
   "bufio"
   "bytes"
   "net/http"
   "os"
   "strings"
)

func main() {
   req, err := http.NewRequest(
      "POST", "http://httpbingo.org/post", strings.NewReader("hello world"),
   )
   if err != nil {
      panic(err)
   }
   buf := &bytes.Buffer{}
   req.Write(buf)
   req, err = http.ReadRequest(bufio.NewReader(buf))
   if err != nil {
      panic(err)
   }
   req.Write(os.Stdout)
}
