package twitter

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
)

func (x subtask) search() (*http.Response, error) {
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
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/2/search/adaptive.json",
      RawQuery: val.Encode(),
   })
   auth := OAuth1{
      AccessToken: x.Open_Account.OAuth_Token,
      AccessSecret: x.Open_Account.OAuth_Token_Secret,
      ConsumerKey: "3nVuSoBZnx6U4vzUxf5w",
      ConsumerSecret: "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys",
   }
   req.Header["Authorization"] = []string{
      auth.BuildOAuth1Header(req.Method, req.URL.String(), val),
   }
   return http.Default_Client.Do(req)
}

func flow_welcome(g *guest) (*task, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/onboarding/task.json",
      RawQuery: "flow_name=welcome",
   })
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "Content-Type": {"application/json"},
      "User-Agent": {"TwitterAndroid/99"},
      "X-Guest-Token": {g.Guest_Token},
   }
   {
      var t task
      t.Input_Flow_Data = new(flow_data)
      t.Input_Flow_Data.Flow_Context.Start_Location.Location = "splash_screen"
      b, err := json.MarshalIndent(t, "", " ")
      if err != nil {
         return nil, err
      }
      req.Body_Bytes(b)
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   t := new(task)
   if err := json.NewDecoder(res.Body).Decode(t); err != nil {
      return nil, err
   }
   return t, nil
}

const bearer = "AAAAAAAAAAAAAAAAAAAAAFXzAwAAAAAAMHCxpeSDG1gLNLghVe8d74hl6k4=RUMF4xAQLsbeBhTSRrCiQpJtxoGWeyHrDb5te2jpGskWDFW82F"

type flow_data struct {
   Flow_Context struct {
      Start_Location struct {
         Location string `json:"location"`
      } `json:"start_location"`
   } `json:"flow_context"`
}

type guest struct {
   Guest_Token string
}

func new_guest() (*guest, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/guest/activate.json",
   })
   req.Header.Set("Authorization", "Bearer " + bearer)
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   g := new(guest)
   if err := json.NewDecoder(res.Body).Decode(g); err != nil {
      return nil, err
   }
   return g, nil
}

type subtask struct {
   Open_Account *struct {
      OAuth_Token string
      OAuth_Token_Secret string
   }
}

type task struct {
   Flow_Token *string `json:"flow_token"`
   Input_Flow_Data *flow_data `json:"input_flow_data"`
   Subtasks []subtask
}

func (t *task) next_link(g *guest) error {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/onboarding/task.json",
   })
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "Content-Type": {"application/json"},
      "X-Guest-Token": {g.Guest_Token},
   }
   {
      b, err := json.MarshalIndent(t, "", " ")
      if err != nil {
         return err
      }
      req.Body_Bytes(b)
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(t)
}

func (t task) open_account() *subtask {
   for _, sub := range t.Subtasks {
      if sub.Open_Account != nil {
         return &sub
      }
   }
   return nil
}
