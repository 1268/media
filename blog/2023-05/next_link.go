package twitter

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
)

func (g Guest) next_link(t *task) (*http.Response, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/onboarding/task.json",
   })
   req.Header["Authorization"] = []string{"Bearer " + bearer}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["X-Guest-Token"] = []string{g.Guest_Token}
   /*
   req.Header["Accept"] = []string{"application/json"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Cache-Control"] = []string{"no-store"}
   req.Header["Cookie"] = []string{"guest_id_marketing=v1%3A168462381830737295; guest_id_ads=v1%3A168462381830737295; personalization_id=v1_5rPm5RI8w/lssWlTrCBkGw==; guest_id=v1%3A168462381830737295"}
   req.Header["Optimize-Body"] = []string{"true"}
   req.Header["Os-Security-Patch-Level"] = []string{"2016-09-06"}
   req.Header["Os-Version"] = []string{"23"}
   req.Header["System-User-Agent"] = []string{"Dalvik/2.1.0 (Linux; U; Android 6.0; Android SDK built for x86 Build/MASTER)"}
   req.Header["Timezone"] = []string{"America/Chicago"}
   req.Header["Twitter-Display-Size"] = []string{"1080x2040x400"}
   req.Header["User-Agent"] = []string{"TwitterAndroid/9.89.0-release.1 (29890001-r-1) Android+SDK+built+for+x86/6.0 (unknown;Android+SDK+built+for+x86;Android;sdk_google_phone_x86;0;;0;2013)"}
   req.Header["X-B3-Traceid"] = []string{"4ea0d2ad68a17019"}
   req.Header["X-Client-Uuid"] = []string{"b9d292b4-eb11-47c1-8df7-f90c4c89389d"}
   req.Header["X-Twitter-Active-User"] = []string{"yes"}
   req.Header["X-Twitter-Api-Version"] = []string{"5"}
   req.Header["X-Twitter-Client"] = []string{"TwitterAndroid"}
   req.Header["X-Twitter-Client-Adid"] = []string{"a00ec847-f729-4862-a0a0-ba1bffa385c3"}
   req.Header["X-Twitter-Client-Appsetid"] = []string{"1ace918f-2331-46ae-8a02-9a4e7f29fba7"}
   req.Header["X-Twitter-Client-Deviceid"] = []string{"c246251b2b00d602"}
   req.Header["X-Twitter-Client-Flavor"] = []string{""}
   req.Header["X-Twitter-Client-Language"] = []string{"en-US"}
   req.Header["X-Twitter-Client-Limit-Ad-Tracking"] = []string{"0"}
   req.Header["X-Twitter-Client-Version"] = []string{"9.89.0-release.1"}
   */
   {
      /*
      var i input
      i.Open_Link.Link = "next_link"
      i.Subtask_ID = "NextTaskOpenLink"
      t.Subtask_Inputs = []input{i}
      */
      t.Input_Flow_Data = nil
      b, err := json.MarshalIndent(t, "", " ")
      if err != nil {
         return nil, err
      }
      req.Body_Bytes(b)
   }
   /*
   req_body := strings.NewReader(`
   {
     "flow_token": "g;168462860984978255:-1684628609855:oO8bop6j0OMkYsvrvUZtd7HF:0",
     "subtask_inputs": [
       {
         "open_link": {
           "link": "next_link"
         },
         "subtask_id": "NextTaskOpenLink"
       }
     ]
   }
   `)
   */
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
