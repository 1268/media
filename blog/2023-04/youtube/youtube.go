package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Encoding"] = []string{"identity"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Authorization"] = []string{"SAPISIDHASH 1681671982_19363636b505fc01fdcb4a6ad68bb8ef7fded946"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Cookie"] = []string{"CONSENT=PENDING+443", "DEVICE_INFO=ChxOekl3TVRVMk5ESTJNRGd6T0RrME16VTROZz09EIqxxJ8GGIqxxJ8G", "GPS=1", "PREF=tz=America.Chicago&f5=30000&autoplay=true&f4=4000000", "VISITOR_INFO1_LIVE=Ba9qACMRbm8", "YSC=pV8QPaSBXrA", "YT_DEVICE_MEASUREMENT_ID=9iccTww=", "__Secure-3PAPISID=pojN-Dn4H1sDNhDL/AiTuoE8bB7InOwAPh", "__Secure-3PSID=VAhNhuH78IBA-U3Su9CwZGvZ6-634Y3IKK73wqRmD9r0F_FcXGZlTkxRiUwu1EYd2Nn7vg.", "__Secure-YEC=CgtucUx3c2NxUWhacyjyptCfBg%3D%3D", "app-start-timestamp-cookie=NIL", "yt-dev.storage-integrity=true"}
   req.Header["Origin"] = []string{"https://www.youtube.com"}
   req.Header["Referer"] = []string{"https://www.youtube.com/"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"same-origin"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0"}
   req.Header["X-Goog-Authuser"] = []string{"0"}
   req.Header["X-Goog-Visitor-Id"] = []string{"CgtCYTlxQUNNUmJtOCitjvGhBg%3D%3D"}
   req.Header["X-Origin"] = []string{"https://www.youtube.com"}
   req.Header["X-Youtube-Bootstrap-Logged-In"] = []string{"false"}
   req.Header["X-Youtube-Client-Name"] = []string{"1"}
   req.Header["X-Youtube-Client-Version"] = []string{"2.20230414.01.00"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.RawPath = ""
   val := make(url.Values)
   val["key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   val["prettyPrint"] = []string{"false"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   for i := 1; i <= 9; i++ {
      req.Body = io.NopCloser(strings.NewReader(req_body))
      res, err := new(http.Transport).RoundTrip(&req)
      if err != nil {
         panic(err)
      }
      body, err := io.ReadAll(res.Body)
      if err != nil {
         panic(err)
      }
      if err := res.Body.Close(); err != nil {
         panic(err)
      }
      fmt.Println(
         "adaptive_formats",
         bytes.Contains(body, []byte(`"adaptiveFormats"`)),
         "view_count",
         bytes.Contains(body, []byte(`"viewCount"`)),
      )
      time.Sleep(time.Second)
   }
}

const req_body = `
{
  "videoId": "3UNsV5H7xxc",
  "context": {
    "client": {
      "hl": "en",
      "gl": "US",
      "remoteHost": "72.181.23.38",
      "deviceMake": "",
      "deviceModel": "",
      "visitorData": "CgtCYTlxQUNNUmJtOCitjvGhBg%3D%3D",
      "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0,gzip(gfe)",
      "clientName": "WEB",
      "clientVersion": "2.20230414.01.00",
      "osName": "Windows",
      "osVersion": "10.0",
      "originalUrl": "https://www.youtube.com/watch?v=3UNsV5H7xxc",
      "screenPixelDensity": 1,
      "platform": "DESKTOP",
      "clientFormFactor": "UNKNOWN_FORM_FACTOR",
      "configInfo": {
        "appInstallData": "CK2O8aEGENburgUQ5_euBRD-7q4FEMyu_hIQt5GvBRC4i64FEOWg_hIQ5LP-EhDi1K4FEKC3_hIQouyuBRCFg68FEKSZrwUQzPWuBRC24K4FENf_rgUQ7YavBRC9tq4FEMzfrgUQuNSuBRCJ6K4FEOmhrwUQuI-vBRCY2q4FEM6ZrwU%3D"
      },
      "screenDensityFloat": 1.25,
      "timeZone": "America/Chicago",
      "browserName": "Firefox",
      "browserVersion": "90.0",
      "acceptHeader": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
      "deviceExperimentId": "ChxOekl3TVRVMk5ESTJNRGd6T0RrME16VTROZz09EK2O8aEGGIqxxJ8G",
      "screenWidthPoints": 817,
      "screenHeightPoints": 626,
      "utcOffsetMinutes": -300,
      "userInterfaceTheme": "USER_INTERFACE_THEME_LIGHT",
      "mainAppWebInfo": {
        "graftUrl": "https://www.youtube.com/watch?v=3UNsV5H7xxc",
        "webDisplayMode": "WEB_DISPLAY_MODE_BROWSER",
        "isWebNativeShareAvailable": false
      },
      "playerType": "UNIPLAYER",
      "tvAppInfo": {
        "livingRoomAppMode": "LIVING_ROOM_APP_MODE_UNSPECIFIED"
      },
      "clientScreen": "WATCH_FULL_SCREEN"
    },
    "user": {
      "lockedSafetyMode": false
    },
    "request": {
      "useSsl": true,
      "internalExperimentFlags": [],
      "consistencyTokenJars": []
    },
    "adSignalsInfo": {
      "params": [
        {
          "key": "dt",
          "value": "1681671982116"
        },
        {
          "key": "flash",
          "value": "0"
        },
        {
          "key": "frm",
          "value": "0"
        },
        {
          "key": "u_tz",
          "value": "-300"
        },
        {
          "key": "u_his",
          "value": "1"
        },
        {
          "key": "u_h",
          "value": "864"
        },
        {
          "key": "u_w",
          "value": "1536"
        },
        {
          "key": "u_ah",
          "value": "824"
        },
        {
          "key": "u_aw",
          "value": "1536"
        },
        {
          "key": "u_cd",
          "value": "24"
        },
        {
          "key": "bc",
          "value": "31"
        },
        {
          "key": "bih",
          "value": "626"
        },
        {
          "key": "biw",
          "value": "800"
        },
        {
          "key": "brdim",
          "value": "132,31,132,31,1536,0,1226,776,817,626"
        },
        {
          "key": "vis",
          "value": "1"
        },
        {
          "key": "wgl",
          "value": "true"
        },
        {
          "key": "ca_type",
          "value": "image"
        }
      ]
    },
    "clickTracking": {
      "clickTrackingParams": "CAAQu2kiEwiYg8OojK_-AhW2huUHHW_yANw="
    },
    "activePlayers": [
      {
        "playerContextParams": "Q0FFU0FnZ0I="
      }
    ]
  },
  "playbackContext": {
    "contentPlaybackContext": {
      "html5Preference": "HTML5_PREF_WANTS",
      "lactMilliseconds": "111",
      "referer": "https://www.youtube.com/watch?v=3UNsV5H7xxc",
      "signatureTimestamp": 19459,
      "autonavState": "STATE_OFF",
      "autoCaptionsDefaultOn": false,
      "mdxContext": {},
      "playerWidthPixels": 343,
      "playerHeightPixels": 193
    }
  },
  "cpn": "KnhBJLIExSai1kLU",
  "captionParams": {
    "deviceCaptionsOn": true
  },
  "attestationRequest": {
    "omitBotguardData": true
  }
}
`
