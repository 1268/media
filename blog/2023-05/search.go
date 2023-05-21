package twitter

import (
   "crypto/hmac"
   "crypto/sha1"
   "encoding/base64"
   "math/rand"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

type OAuth1 struct {
   ConsumerKey string
   ConsumerSecret string
   AccessToken string
   AccessSecret string
}

// Params being any key-value url query parameter pairs
func (auth OAuth1) BuildOAuth1Header(method, path string, params map[string]string) string {
   vals := url.Values{}
   vals.Add("oauth_nonce", generateNonce())
   vals.Add("oauth_consumer_key", auth.ConsumerKey)
   vals.Add("oauth_signature_method", "HMAC-SHA1")
   vals.Add("oauth_timestamp", strconv.Itoa(int(time.Now().Unix())))
   vals.Add("oauth_token", auth.AccessToken)
   vals.Add("oauth_version", "1.0")

   for k, v := range params {
      vals.Add(k, v)
   }
   // net/url package QueryEscape escapes " " into "+", this replaces it with the percentage encoding of " "
   parameterString := strings.Replace(vals.Encode(), "+", "%20", -1)

   // Calculating Signature Base String and Signing Key
   signatureBase := strings.ToUpper(method) + "&" + url.QueryEscape(strings.Split(path, "?")[0]) + "&" + url.QueryEscape(parameterString)
   signingKey := url.QueryEscape(auth.ConsumerSecret) + "&" + url.QueryEscape(auth.AccessSecret)
   signature := calculateSignature(signatureBase, signingKey)

   return "OAuth oauth_consumer_key=\"" + url.QueryEscape(vals.Get("oauth_consumer_key")) + "\", oauth_nonce=\"" + url.QueryEscape(vals.Get("oauth_nonce")) +
      "\", oauth_signature=\"" + url.QueryEscape(signature) + "\", oauth_signature_method=\"" + url.QueryEscape(vals.Get("oauth_signature_method")) +
      "\", oauth_timestamp=\"" + url.QueryEscape(vals.Get("oauth_timestamp")) + "\", oauth_token=\"" + url.QueryEscape(vals.Get("oauth_token")) +
      "\", oauth_version=\"" + url.QueryEscape(vals.Get("oauth_version")) + "\""
}

func calculateSignature(base, key string) string {
   hash := hmac.New(sha1.New, []byte(key))
   hash.Write([]byte(base))
   signature := hash.Sum(nil)
   return base64.StdEncoding.EncodeToString(signature)
}

func generateNonce() string {
   const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
   b := make([]byte, 48)
   for i := range b {
      b[i] = allowed[rand.Intn(len(allowed))]
   }
   return string(b)
}

func (x subtask) search() (*http.Response, error) {
   req, err := http.NewRequest(
      "GET", "https://api.twitter.com/2/search/adaptive.json", nil,
   )
   if err != nil {
      return nil, err
   }
   val := make(url.Values)
   val["cards_platform"] = []string{"Android-12"}
   val["earned"] = []string{"true"}
   val["ext"] = []string{"mediaRestrictions,altText,mediaStats,mediaColor,info360,highlightedLabel,superFollowMetadata,hasNftAvatar,unmentionInfo"}
   val["include_blocked_by"] = []string{"true"}
   val["include_blocking"] = []string{"true"}
   val["include_cards"] = []string{"true"}
   val["include_carousels"] = []string{"true"}
   val["include_composer_source"] = []string{"true"}
   val["include_entities"] = []string{"true"}
   val["include_ext_enrichments"] = []string{"true"}
   val["include_ext_has_nft_avatar"] = []string{"true"}
   val["include_ext_media_availability"] = []string{"true"}
   val["include_ext_professional"] = []string{"true"}
   val["include_ext_replyvoting_downvote_perspective"] = []string{"true"}
   val["include_ext_sensitive_media_warning"] = []string{"true"}
   val["include_media_features"] = []string{"true"}
   val["include_profile_interstitial_type"] = []string{"true"}
   val["include_quote_count"] = []string{"true"}
   val["include_reply_count"] = []string{"true"}
   val["include_user_entities"] = []string{"true"}
   val["include_viewer_quick_promote_eligibility"] = []string{"true"}
   val["q"] = []string{"filter:spaces"}
   val["query_source"] = []string{"typed_query"}
   val["simple_quoted_tweet"] = []string{"true"}
   val["spelling_corrections"] = []string{"true"}
   val["tweet_mode"] = []string{"extended"}
   val["tweet_search_mode"] = []string{"top"}
   param := make(map[string]string)
   for key := range val {
      param[key] = val.Get(key)
   }
   req.URL.RawQuery = val.Encode()
   auth := OAuth1{
      AccessToken: x.Open_Account.OAuth_Token,
      AccessSecret: x.Open_Account.OAuth_Token_Secret,
      ConsumerKey: "3nVuSoBZnx6U4vzUxf5w",
      ConsumerSecret: "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys",
   }
   req.Header["Authorization"] = []string{
      auth.BuildOAuth1Header(req.Method, req.URL.String(), param),
   }
   return new(http.Transport).RoundTrip(req)
}
