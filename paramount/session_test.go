package paramount

import (
   "2a.pages.dev/mech/widevine"
   "encoding/base64"
   "encoding/json"
   "os"
   "testing"
   "time"
)

func Test_Post(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   test := tests[0]
   pssh, err := base64.StdEncoding.DecodeString(test.pssh)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, pssh)
   if err != nil {
      t.Fatal(err)
   }
   sess, err := New_Session(test.guid)
   if err != nil {
      t.Fatal(err)
   }
   keys, err := mod.Post(sess)
   if err != nil {
      t.Fatal(err)
   }
   if keys.Content().String() != test.key {
      t.Fatal(keys)
   }
}
func Test_Session(t *testing.T) {
   test := tests[0]
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.SetEscapeHTML(false)
   for version, secret := range app_secrets {
      if secret != "" {
         sess, err := session_secret(test.guid, secret)
         if err != nil {
            t.Fatal(version, err)
         }
         if err := enc.Encode(sess); err != nil {
            t.Fatal(err)
         }
         time.Sleep(99 * time.Millisecond)
      }
   }
}
