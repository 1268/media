package twitter

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
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

func access_token() (*header, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/oauth2/token",
   })
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   req.SetBasicAuth(
      "3nVuSoBZnx6U4vzUxf5w",
      "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys",
   )
   req.Body_String("grant_type=client_credentials")
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var head header
   {
      var s struct {
         Access_Token string
      }
      if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
         return nil, err
      }
      head.Header = make(http.Header)
      head.Set("Authorization", "Bearer " + s.Access_Token)
   }
   return &head, nil
}

func (h header) activate() error {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/guest/activate.json",
   })
   req.Header = h.Header
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   var s struct {
      Guest_Token string
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return err
   }
   h.Set("X-Guest-Token", s.Guest_Token)
   return nil
}

type header struct {
   http.Header
}

func welcome(h header) (*flow, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/onboarding/task.json",
      RawQuery: "flow_name=welcome",
   })
   req.Header = h.Header
   req.Header.Set("Content-Type", "application/json")
   req.Header.Set("User-Agent", "TwitterAndroid/99")
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

func (f *flow) next_link(h header) error {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/onboarding/task.json",
   })
   req.Header = h.Header
   req.Header.Set("Conten-Type", "application/json")
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
