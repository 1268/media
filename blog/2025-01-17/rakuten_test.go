package rakuten

import (
   "fmt"
   "testing"
)

func TestHello(t *testing.T) {
   fmt.Println("Hello")
}

func TestWorld(t *testing.T) {
   fmt.Println("World")
}

func TestMain(m *testing.M) {
   fmt.Println("Main")
   m.Run()
}
