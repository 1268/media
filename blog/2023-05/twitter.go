package twitter

import (
   "2a.pages.dev/rosso/http"
   "bytes"
   "crypto/hmac"
   "crypto/sha1"
   "encoding/base64"
   "encoding/json"
   "net/url"
   "strconv"
   "time"
)

func (x subtask) search(q string) (*search, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/2/search/adaptive.json",
      RawQuery: url.Values{
         "q": {q},
         // This ensures Spaces Tweets will include Spaces URL
         "tweet_mode": {"extended"},
      }.Encode(),
   })
   auth := oauth{
      consumer_key: "3nVuSoBZnx6U4vzUxf5w",
      consumer_secret: "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys",
      token: x.Open_Account.OAuth_Token,
      token_secret: x.Open_Account.OAuth_Token_Secret,
   }
   req.Header["Authorization"] = []string{auth.sign(req.Method, req.URL)}
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   s := new(search)
   if err := json.NewDecoder(res.Body).Decode(s); err != nil {
      return nil, err
   }
   return s, nil
}

type search struct {
   GlobalObjects struct {
      Tweets map[int64]struct {
         Entities struct {
            URLs []struct {
               Expanded_URL string
            }
         }
      }
   }
}

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
