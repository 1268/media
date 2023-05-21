package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["Authorization"] = []string{"OAuth " + strings.Join([]string{
      `oauth_consumer_key="3nVuSoBZnx6U4vzUxf5w"`,
      `oauth_nonce="5775992602006906429184916339582"`,
      `oauth_signature_method="HMAC-SHA1"`,
      `oauth_timestamp="1684614059"`,
      `oauth_version="1.0"`,
      `oauth_token="1660017733327630336-FClvhvGL5hO3eahU1O2cuCBHbEwOyc"`,
      // needs to be encoded:
      `oauth_signature="I64WGB%2BOG0Pe87pCKzi4w091w2w%3D"`,
   }, ",")}
   req.URL = new(url.URL)
   req.URL.Host = "api.twitter.com"
   req.URL.Path = "/2/search/adaptive.json"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["q"] = []string{"filter:spaces"}
   val["ext"] = []string{"mediaRestrictions,altText,mediaStats,mediaColor,info360,highlightedLabel,hasNftAvatar,unmentionInfo,editControl,previousCounts,limitedActionResults,superFollowMetadata"}
   val["cards_platform"] = []string{"Android-12"}
   val["earned"] = []string{"true"}
   val["include_blocked_by"] = []string{"true"}
   val["include_blocking"] = []string{"true"}
   val["include_cards"] = []string{"true"}
   val["include_carousels"] = []string{"true"}
   val["include_composer_source"] = []string{"true"}
   val["include_entities"] = []string{"true"}
   val["include_ext_birdwatch_pivot"] = []string{"true"}
   val["include_ext_edit_control"] = []string{"true"}
   val["include_ext_has_nft_avatar"] = []string{"true"}
   val["include_ext_is_blue_verified"] = []string{"true", "true"}
   val["include_ext_is_tweet_translatable"] = []string{"true", "true"}
   val["include_ext_limited_action_results"] = []string{"true"}
   val["include_ext_media_availability"] = []string{"true"}
   val["include_ext_previous_counts"] = []string{"true"}
   val["include_ext_professional"] = []string{"true"}
   val["include_ext_profile_image_shape"] = []string{"true", "true"}
   val["include_ext_sensitive_media_warning"] = []string{"true"}
   val["include_ext_trusted_friends_metadata"] = []string{"true"}
   val["include_ext_verified_type"] = []string{"true", "true"}
   val["include_ext_views"] = []string{"true"}
   val["include_media_features"] = []string{"true"}
   val["include_profile_interstitial_type"] = []string{"true"}
   val["include_quote_count"] = []string{"true"}
   val["include_reply_count"] = []string{"true"}
   val["include_user_entities"] = []string{"true"}
   val["include_viewer_quick_promote_eligibility"] = []string{"true"}
   val["query_source"] = []string{"typed_query"}
   val["simple_quoted_tweet"] = []string{"true"}
   val["spelling_corrections"] = []string{"true"}
   val["tweet_mode"] = []string{"extended"}
   val["tweet_search_mode"] = []string{"top"}
   req.URL.RawQuery = val.Encode()
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   b, err := httputil.DumpRequest(req, false)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(b)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}
