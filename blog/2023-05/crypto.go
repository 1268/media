package twitter

import (
   "crypto/hmac"
   "crypto/sha1"
   "encoding/base64"
   "net/url"
   "strconv"
   "strings"
   "time"
)

func calculateSignature(base, key string) string {
   hash := hmac.New(sha1.New, []byte(key))
   hash.Write([]byte(base))
   signature := hash.Sum(nil)
   return base64.StdEncoding.EncodeToString(signature)
}

func generateNonce() string {
   return strings.Repeat("0", 48)
}

// Params being any key-value url query parameter pairs
func (auth OAuth1) BuildOAuth1Header(method, path string, val url.Values) string {
   val.Add("oauth_nonce", generateNonce())
   val.Add("oauth_consumer_key", auth.ConsumerKey)
   val.Add("oauth_signature_method", "HMAC-SHA1")
   val.Add("oauth_timestamp", strconv.Itoa(int(time.Now().Unix())))
   val.Add("oauth_token", auth.AccessToken)
   val.Add("oauth_version", "1.0")
   // Calculating Signature Base String and Signing Key
   signatureBase := strings.ToUpper(method) + "&" + url.QueryEscape(strings.Split(path, "?")[0]) + "&" + url.QueryEscape(val.Encode())
   signingKey := url.QueryEscape(auth.ConsumerSecret) + "&" + url.QueryEscape(auth.AccessSecret)
   signature := calculateSignature(signatureBase, signingKey)
   return "OAuth " + strings.Join([]string{
      "oauth_consumer_key=" + val.Get("oauth_consumer_key"),
      "oauth_nonce=" + val.Get("oauth_nonce"),
      "oauth_signature=" + url.QueryEscape(signature),
      "oauth_signature_method=" + val.Get("oauth_signature_method"),
      "oauth_timestamp=" + val.Get("oauth_timestamp"),
      "oauth_token=" + val.Get("oauth_token"),
      "oauth_version=" + val.Get("oauth_version"),
   }, ",")
}

type OAuth1 struct {
   ConsumerKey string
   ConsumerSecret string
   AccessToken string
   AccessSecret string
}

