package media

import (
   "fmt"
   "testing"
   "time"
)

func Test_Name(t *testing.T) {
   for _, test := range tests {
      s, err := Name(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(s)
   }
}

type tester struct {
   series string
   season int64
   episode int64
}

func (t tester) Date() (time.Time, error) { return time.Now(), nil }
func (t tester) Episode() (int64, error) { return t.episode, nil }
func (t tester) Season() (int64, error) { return t.season, nil }
func (t tester) Series() string { return t.series }
func (t tester) Title() string { return "title" }

var tests = []tester{
   // amc movie
   {},
   // amc show
   {"Orphan Black", 1, 2},
   // cbc\gem.go
   // nbc\nbc.go
   // paramount\item.go
   // roku\roku.go
}

func Test_Clean(t *testing.T) {
   hello := Clean("one * two ? three")
   fmt.Println(hello)
}
