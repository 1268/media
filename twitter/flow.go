// Support for this software's development was paid for by
// Fredrick R. Brennan's Modular Font Editor K Foundation, Inc.
package twitter

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

func welcome(access_token, guest_token string) (*flow, error) {
   body, err := func() ([]byte, error) {
      var f flow
      f.Input_Flow_Data = new(flow_data)
      f.Input_Flow_Data.Flow_Context.Start_Location.Location = "splash_screen"
      return json.MarshalIndent(f, "", " ")
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST",
      "https://api.twitter.com/1.1/onboarding/task.json?flow_name=welcome",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + access_token},
      "Content-Type": {"application/json"},
      "User-Agent": {"TwitterAndroid/99"},
      "X-Guest-Token": {guest_token},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   f := new(flow)
   if err := json.NewDecoder(res.Body).Decode(f); err != nil {
      return nil, err
   }
   return f, nil
}

func (f *flow) next_link(access_token, guest_token string) error {
   body, err := json.MarshalIndent(f, "", " ")
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://api.twitter.com/1.1/onboarding/task.json",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + access_token},
      "Content-Type": {"application/json"},
      "X-Guest-Token": {guest_token},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return json.NewDecoder(res.Body).Decode(f)
}

func (f flow) open_account() *subtask {
   for _, task := range f.Subtasks {
      if task.Open_Account != nil {
         return &task
      }
   }
   return nil
}

const (
   consumer_key = "3nVuSoBZnx6U4vzUxf5w"
   consumer_secret = "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys"
)

type flow struct {
   Flow_Token *string `json:"flow_token"`
   Input_Flow_Data *flow_data `json:"input_flow_data"`
   Subtasks []subtask
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

func access_token() (string, error) {
   req, err := http.NewRequest(
      "POST", "https://api.twitter.com/oauth2/token",
      strings.NewReader("grant_type=client_credentials"),
   )
   if err != nil {
      return "", err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.SetBasicAuth(consumer_key, consumer_secret)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return "", errors.New(res.Status)
   }
   var s struct {
      Access_Token string
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return "", err
   }
   return s.Access_Token, nil
}

func guest_token(access_token string) (string, error) {
   req, err := http.NewRequest(
      "POST", "https://api.twitter.com/1.1/guest/activate.json", nil,
   )
   if err != nil {
      return "", err
   }
   req.Header.Set("Authorization", "Bearer " + access_token)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return "", errors.New(res.Status)
   }
   var s struct {
      Guest_Token string
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return "", err
   }
   return s.Guest_Token, nil
}
