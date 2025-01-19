package hulu

import (
   "os"
   "os/exec"
   "strings"
   "testing"
)

func TestAuthenticate(t *testing.T) {
   data, err := exec.Command("password", "hulu.com").Output()
   if err != nil {
      t.Fatal(err)
   }
   username, password, _ := strings.Cut(string(data), ":")
   data, err = new(Authenticate).Marshal(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.txt", data, os.ModePerm)
}

var tests = []struct {
   content string
   key_id  string
   url     string
}{
   {
      content: "episode",
      key_id:  "21b82dc2ebb24d5aa9f8631f04726650",
      url:     "hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d",
   },
   {
      content: "film",
      url:     "hulu.com/watch/f70dfd4d-dbfb-46b8-abb3-136c841bba11",
   },
}
