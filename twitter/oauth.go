// Support for this software's development was paid for by
// Fredrick R. Brennan's Modular Font Editor K Foundation, Inc.
package twitter

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha1"
   "encoding/base64"
   "net/url"
   "strconv"
   "time"
)

type oauth struct {
   consumer_key string
   consumer_secret string
   token string
   token_secret string
}

func (o oauth) sign(method string, ref *url.URL) string {
   m := value_map{
      "oauth_consumer_key": o.consumer_key,
      "oauth_nonce": "0",
      "oauth_signature_method": "HMAC-SHA-1",
      "oauth_timestamp": strconv.FormatInt(time.Now().Unix(), 10),
      "oauth_token": o.token,
   }
   key := value_slice{o.consumer_secret, o.token_secret}
   h := hmac.New(sha1.New, key.Bytes())
   {
      query := ref.Query()
      for key, value := range m {
         query.Set(key, value)
      }
      req := value_slice{
         method,
         ref.Scheme + "://" + ref.Host + ref.Path,
         query.Encode(),
      }
      h.Write(req.Bytes())
   }
   m["oauth_signature"] = base64.StdEncoding.EncodeToString(h.Sum(nil))
   return "OAuth " + m.String()
}

type value_map map[string]string

func (v value_map) String() string {
   var b bytes.Buffer
   for key, value := range v {
      if b.Len() >= 1 {
         b.WriteByte(',')
      }
      b.WriteString(key)
      b.WriteByte('=')
      b.WriteString(url.QueryEscape(value))
   }
   return b.String()
}

type value_slice []string

func (v value_slice) Bytes() []byte {
   var b bytes.Buffer
   for _, value := range v {
      if b.Len() >= 1 {
         b.WriteByte('&')
      }
      b.WriteString(url.QueryEscape(value))
   }
   return b.Bytes()
}
