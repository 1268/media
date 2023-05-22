package twitter

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
)

const (
   consumer_key = "3nVuSoBZnx6U4vzUxf5w"
   consumer_secret = "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys"
)

type flow struct {
   Flow_Token *string `json:"flow_token"`
   Input_Flow_Data *flow_data `json:"input_flow_data"`
   Subtasks []subtask
}

func (f flow) open_account() *subtask {
   for _, task := range f.Subtasks {
      if task.Open_Account != nil {
         return &task
      }
   }
   return nil
}

type flow_data struct {
   Flow_Context struct {
      Start_Location struct {
         Location string `json:"location"`
      } `json:"start_location"`
   } `json:"flow_context"`
}

type subtask struct {
   Open_Account *struct {
      OAuth_Token string
      OAuth_Token_Secret string
   }
}

// this always returns the same output. how old are the username and password?
func access_token() (string, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/oauth2/token",
   })
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.SetBasicAuth(consumer_key, consumer_secret)
   req.Body_String("grant_type=client_credentials")
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   var s struct {
      Access_Token string
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return "", err
   }
   return s.Access_Token, nil
}

func guest_token(access_token string) (string, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/guest/activate.json",
   })
   req.Header.Set("Authorization", "Bearer " + access_token)
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   var s struct {
      Guest_Token string
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return "", err
   }
   return s.Guest_Token, nil
}

func welcome(access_token, guest_token string) (*flow, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/onboarding/task.json",
      RawQuery: "flow_name=welcome",
   })
   req.Header = http.Header{
      "Authorization": {"Bearer " + access_token},
      "Content-Type": {"application/json"},
      "User-Agent": {"TwitterAndroid/99"},
      "X-Guest-Token": {guest_token},
   }
   {
      var f flow
      f.Input_Flow_Data = new(flow_data)
      f.Input_Flow_Data.Flow_Context.Start_Location.Location = "splash_screen"
      b, err := json.MarshalIndent(f, "", " ")
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
   f := new(flow)
   if err := json.NewDecoder(res.Body).Decode(f); err != nil {
      return nil, err
   }
   return f, nil
}

func (f *flow) next_link(access_token, guest_token string) error {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/onboarding/task.json",
   })
   req.Header = http.Header{
      "Authorization": {"Bearer " + access_token},
      "Content-Type": {"application/json"},
      "X-Guest-Token": {guest_token},
   }
   {
      b, err := json.MarshalIndent(f, "", " ")
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
   return json.NewDecoder(res.Body).Decode(f)
}
