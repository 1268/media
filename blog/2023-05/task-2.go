package twitter

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func task_two() {
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
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"application/json"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Authorization"] = []string{"Bearer AAAAAAAAAAAAAAAAAAAAAFXzAwAAAAAAMHCxpeSDG1gLNLghVe8d74hl6k4%3DRUMF4xAQLsbeBhTSRrCiQpJtxoGWeyHrDb5te2jpGskWDFW82F"}
   req.Header["Cache-Control"] = []string{"no-store"}
   req.Header["Content-Type"] = []string{"application/json"}
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
   req.Header["X-Guest-Token"] = []string{"1660058730438836226"}
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
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.twitter.com"
   req.URL.Path = "/1.1/onboarding/task.json"
   req.URL.RawPath = ""
   val := make(url.Values)
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
   res, err := new(http.Transport).RoundTrip(&req)
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

