package twitter

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
)

const bearer = "AAAAAAAAAAAAAAAAAAAAAFXzAwAAAAAAMHCxpeSDG1gLNLghVe8d74hl6k4%3DRUMF4xAQLsbeBhTSRrCiQpJtxoGWeyHrDb5te2jpGskWDFW82F"

type Guest struct {
   Guest_Token string
}

func New_Guest() (*Guest, error) {
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
   g := new(Guest)
   if err := json.NewDecoder(res.Body).Decode(g); err != nil {
      return nil, err
   }
   return g, nil
}

func (g Guest) flow_welcome() (*task, error) {
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
   var t task
   {
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
   if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
      return nil, err
   }
   return &t, nil
}

func (g Guest) next_link(t *task) (*http.Response, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/onboarding/task.json",
   })
   req.Header["Authorization"] = []string{"Bearer " + bearer}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["X-Guest-Token"] = []string{g.Guest_Token}
   {
      t.Input_Flow_Data = nil
      b, err := json.MarshalIndent(t, "", " ")
      if err != nil {
         return nil, err
      }
      req.Body_Bytes(b)
   }
   return http.Default_Client.Do(req)
}

type input struct {
   Open_Link struct {
      Link string
   }
   Subtask_ID string
}

type flow_data struct {
   Flow_Context struct {
      Start_Location struct {
         Location string `json:"location"`
      } `json:"start_location"`
   } `json:"flow_context"`
}

type task struct {
   Flow_Token *string `json:"flow_token"`
   Input_Flow_Data *flow_data `json:"input_flow_data,omitempty"`
   Subtask_Inputs []input `json:",omitempty"`
}
