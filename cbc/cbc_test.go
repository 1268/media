package cbc

import (
   "fmt"
   "os"
   "testing"
)

const downton = "downton-abbey/s01e05"

func Test_Media(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   profile, err := Open_Profile(home + "/mech/cbc.json")
   if err != nil {
      t.Fatal(err)
   }
   asset, err := New_Asset(downton)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(asset)
   media, err := profile.Media(asset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}

func Test_Profile(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   Client.Log_Level = 2
   login, err := New_Login(email, password)
   if err != nil {
      t.Fatal(err)
   }
   web, err := login.Web_Token()
   if err != nil {
      t.Fatal(err)
   }
   top, err := web.Over_The_Top()
   if err != nil {
      t.Fatal(err)
   }
   profile, err := top.Profile()
   if err != nil {
      t.Fatal(err)
   }
   if err := profile.Create(home + "/mech/cbc.json"); err != nil {
      t.Fatal(err)
   }
}

