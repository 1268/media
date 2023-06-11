package youtube

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Android_Embed(t *testing.T) {
   var req Request
   req.Android_Embed()
   for _, embed := range android_embeds {
      req.Video_ID = embed
      play, err := req.Player(nil)
      if err != nil {
         t.Fatal(err)
      }
      if play.Playability_Status.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

func Test_Search(t *testing.T) {
   var req Request
   req.Mobile_Web()
   search, err := req.Search("oneohtrix point never along")
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   if err := enc.Encode(search); err != nil {
      t.Fatal(err)
   }
   for _, item := range search.Items() {
      fmt.Println(item.Video_With_Context_Renderer)
   }
}

func Test_Config(t *testing.T) {
   con, err := new_config()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(con)
}

var androids = []string{
   "H1BuwMTrtLQ", // content check
   "zv9NimPx3Es",
}

func Test_Android(t *testing.T) {
   for _, android := range androids {
      play, err := Android().Player(android, nil)
      if err != nil {
         t.Fatal(err)
      }
      if play.Playability_Status.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

var android_checks = []string{
   "Cr381pDsSsA", // racy check
   "HsUATh_Nc2U", // racy check
   "SZJvDhaSDnc", // racy check
   "Tq92D6wQ1mg", // racy check
   "dqRZDebPIGs", // racy check
   "nGC3D_FkCmg", // content check
}

func Test_Android_Check(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   req := Android_Check()
   tok, err := Read_Token(home + "/mech/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, check := range android_checks {
      play, err := req.Player(check, tok)
      if err != nil {
         t.Fatal(err)
      }
      if play.Playability_Status.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

var android_embeds = []string{
   "HtVdAasjOgU",
   "WaOKSUlf4TM",
}
