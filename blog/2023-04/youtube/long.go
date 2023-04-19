package main

import (
   "bufio"
   "fmt"
   "os"
)

func main() {
   file, err := os.Open("ignore.go")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   scan := bufio.NewScanner(file)
   var (
      line int
      top_length int
      top_line int
   )
   for scan.Scan() {
      line++
      length := len(scan.Bytes())
      if length > top_length {
         top_line = line
         top_length = length
      }
   }
   fmt.Println(top_line, top_length)
}
