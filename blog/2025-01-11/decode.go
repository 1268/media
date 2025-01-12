package main

import (
   "bufio"
   "bytes"
   "fmt"
   "github.com/ugorji/go/codec"
   "io"
   "net/http"
   "os"
)

func main() {
   data, err := os.ReadFile("req.txt")
   if err != nil {
      panic(err)
   }
   req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(data)))
   if err != nil {
      panic(err)
   }
   data, err = io.ReadAll(req.Body)
   if err != nil {
      panic(err)
   }
   var value any
   err = codec.NewDecoderBytes(data, &codec.CborHandle{}).Decode(&value)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%#v\n", value)
}
