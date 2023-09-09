package media

import (
   "fmt"
   "testing"
)

// amc\video.go
// cbc\gem.go
// nbc\nbc.go
// paramount\item.go
// roku\roku.go
func Test_Name(t *testing.T) {
   fmt.Println("hello world")
}

func Test_Clean(t *testing.T) {
   hello := Clean("one * two ? three")
   fmt.Println(hello)
}
