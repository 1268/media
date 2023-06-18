package widevine

import (
   "encoding/base64"
   "os"
   "testing"
)

// therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76
const (
   raw_PSSH = "AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQvfpNbNs5cC5baB+QYX+afhoKaW50ZXJ0cnVzdCIBKg=="
   raw_response = "CAISpAEKGgoAEghBPi3uL18dhiABKAA4AEAASOrh758GEh4IARAAGAAgACgAMMnFCjgAQgBIAFAAWABgAHAAeAEaXgoQvfpNbNs5cC5baB+QYX+afhIQUkaPkIyfx8bplsyEg21YKhogQYRNMicID9i4twGGFm/fdKEmJGUVQuNnBXXfUtmYXH4gAigBMggIABAqGAAgADoICAAQKhgAIAAg6uHvnwY4ABoga0R69yOAW/lmax2ng2W92p0JzaIz+GNvKhgucUFs0qsigAKU3+gmgfhY4c6YXXDpAi5lt192PUBfXUCJ6WX+zu3haUI1sY9VxA389iUQ470xm6SW5mF2vcFg7NUeCFEeq+Y15GgZnN9JuLWOq7GQIqsQioLpKvQMIwEamd/KP16KtL8UQD6cCU6/tFQipbiGYw4XSeSQjSeqScZjWhkwzu69V6gYWgCeL7BuLqE4BrfgyUVjuGt2CeChhAxOZC8n1McIxxpZ50ST7F1HdWlpjRj1WrlsMLzLH/FaBWo4zqtGQxyg92d28AlwpiLgVIt1pDNt8PgrvHsyjCKhL4Vh373wmSMuiwVbBcG8Rl2x29ek2ot9UP9iZjSnNuC6BSsO9d6FOggKBjE3LjAuMUABWAA="
)

func Test_Response(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "2a/mech/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/2a/mech/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   pssh, err := base64.StdEncoding.DecodeString(raw_PSSH)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := New_Module(private_key, client_ID, pssh)
   if err != nil {
      t.Fatal(err)
   }
   response, err := base64.StdEncoding.DecodeString(raw_response)
   if err != nil {
      t.Fatal(err)
   }
   if _, err := mod.signed_response(response); err != nil {
      t.Fatal(err)
   }
}
