package main

import (
   "bufio"
   "fmt"
   "github.com/ugorji/go/codec"
   "io"
   "net/http"
   "os"
)

func main() {
   file, err := os.Open("req.txt")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   req, err := http.ReadRequest(bufio.NewReader(file))
   if err != nil {
      panic(err)
   }
   data, err := io.ReadAll(req.Body)
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
