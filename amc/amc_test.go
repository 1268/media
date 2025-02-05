package amc

import (
   "fmt"
   "os"
   "os/exec"
   "strings"
   "testing"
)

func TestLogin(t *testing.T) {
   data, err := exec.Command("password", "amcplus.com").Output()
   if err != nil {
      t.Fatal(err)
   }
   username, password, _ := strings.Cut(string(data), ":")
   var auth Authorization
   err = auth.Unauth()
   if err != nil {
      t.Fatal(err)
   }
   data, err = auth.Login(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("amc.txt", data, os.ModePerm)
}

var key_tests = []struct{
   key_id string
   url string
}{
   {
      key_id: "+7nUc5piRu2GY3lAiA4MvQ==",
      url: "amcplus.com/movies/nocebo--1061554",
   },
   {
      key_id: "vHkdO0RPSsqD3iPzeupPeA==",
      url: "amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152",
   },
}

var path_tests = []string{
   "/movies/nocebo--1061554",
   "amcplus.com/movies/nocebo--1061554",
   "https://www.amcplus.com/movies/nocebo--1061554",
   "www.amcplus.com/movies/nocebo--1061554",
}

func TestPath(t *testing.T) {
   for _, test := range path_tests {
      var web Address
      err := web.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(web)
   }
}

func TestRefresh(t *testing.T) {
   data, err := os.ReadFile("amc.txt")
   if err != nil {
      t.Fatal(err)
   }
   var auth Authorization
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   data, err = auth.Refresh()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("amc.txt", data, os.ModePerm)
}
