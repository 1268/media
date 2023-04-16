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
   req.Header["Authorization"] = []string{"SAPISIDHASH 1681675603_b1677ab3f62c826218fd106051f1d6e87ea00da3"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Cookie"] = []string{"CONSENT=PENDING+443", "DEVICE_INFO=ChxOekl3TVRVMk5ESTJNRGd6T0RrME16VTROZz09EIqxxJ8GGIqxxJ8G", "GPS=1", "PREF=tz=America.Chicago&f5=30000&autoplay=true&f4=4000000", "ST-1cwo4nj=itct=CKMBEKQwGAIiEwjx7-bama_-AhUltOUHHcFeCIYyB3JlbGF0ZWRIpsqRu__njoA0mgEFCAEQ-B0%3D&csn=MC4wMjIxNDM2NjAxNTYxODM0MQ..&endpoint=%7B%22clickTrackingParams%22%3A%22CKMBEKQwGAIiEwjx7-bama_-AhUltOUHHcFeCIYyB3JlbGF0ZWRIpsqRu__njoA0mgEFCAEQ-B0%3D%22%2C%22commandMetadata%22%3A%7B%22webCommandMetadata%22%3A%7B%22url%22%3A%22%2Fwatch%3Fv%3DcoLCY15P6Bw%22%2C%22webPageType%22%3A%22WEB_PAGE_TYPE_WATCH%22%2C%22rootVe%22%3A3832%7D%7D%2C%22watchEndpoint%22%3A%7B%22videoId%22%3A%22coLCY15P6Bw%22%2C%22nofollow%22%3Atrue%2C%22watchEndpointSupportedOnesieConfig%22%3A%7B%22html5PlaybackOnesieConfig%22%3A%7B%22commonConfig%22%3A%7B%22url%22%3A%22https%3A%2F%2Frr3---sn-q4fl6n66.googlevideo.com%2Finitplayback%3Fsource%3Dyoutube%26oeis%3D1%26c%3DWEB%26oad%3D3200%26ovd%3D3200%26oaad%3D11000%26oavd%3D11000%26ocs%3D700%26oewis%3D1%26oputc%3D1%26ofpcc%3D1%26msp%3D1%26odepv%3D1%26id%3D7282c2635e4fe81c%26ip%3D72.181.23.38%26initcwndbps%3D1750000%26mt%3D1681675269%26oweuc%3D%26pxtags%3DCg4KAnR4EggyNDQ4NjU3MA%26rxtags%3DCg4KAnR4EggyNDQ4NjU3MA%252CCg4KAnR4EggyNDQ4NjU3MQ%252CCg4KAnR4EggyNDQ4NjU3Mg%22%7D%7D%7D%7D%7D", "VISITOR_INFO1_LIVE=Ba9qACMRbm8", "YSC=pV8QPaSBXrA", "YT_DEVICE_MEASUREMENT_ID=9iccTww=", "__Secure-3PAPISID=pojN-Dn4H1sDNhDL/AiTuoE8bB7InOwAPh", "__Secure-3PSID=VAhNhuH78IBA-U3Su9CwZGvZ6-634Y3IKK73wqRmD9r0F_FcXGZlTkxRiUwu1EYd2Nn7vg.", "__Secure-YEC=CgtucUx3c2NxUWhacyjyptCfBg%3D%3D", "app-start-timestamp-cookie=NIL", "yt-dev.storage-integrity=true"}
   req.Header["Origin"] = []string{"https://www.youtube.com"}
   req.Header["Referer"] = []string{"https://www.youtube.com/watch?v=NAA7P_dkZSY"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"same-origin"}
   req.Header["Sec-Fetch-Site"] = []string{"same-origin"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0"}
   req.Header["X-Goog-Authuser"] = []string{"0"}
   req.Header["X-Goog-Visitor-Id"] = []string{"CgtCYTlxQUNNUmJtOCi4qvGhBg%3D%3D"}
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
  "context": {
    "client": {
      "hl": "en",
      "gl": "US",
      "remoteHost": "72.181.23.38",
      "deviceMake": "",
      "deviceModel": "",
      "visitorData": "CgtCYTlxQUNNUmJtOCi4qvGhBg%3D%3D",
      "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0,gzip(gfe)",
        "clientName": "ANDROID",
   "clientVersion": "18.14.40",
      "osName": "Windows",
      "osVersion": "10.0",
      "originalUrl": "https://www.youtube.com/watch?v=coLCY15P6Bw",
      "screenPixelDensity": 1,
      "platform": "DESKTOP",
      "clientFormFactor": "UNKNOWN_FORM_FACTOR",
      "configInfo": {
        "appInstallData": "CLiq8aEGEMyu_hIQvbauBRDX_64FEMz1rgUQzN-uBRDi1K4FEKC3_hIQ5aD-EhDn964FEInorgUQ6aGvBRD-7q4FEIWDrwUQ5LP-EhC24K4FEO2GrwUQt5GvBRCi7K4FELiLrgUQpJmvBRDW7q4FELjUrgUQuI-vBRCY2q4FEM6ZrwU%3D"
      },
      "screenDensityFloat": 1.25,
      "timeZone": "America/Chicago",
      "browserName": "Firefox",
      "browserVersion": "90.0",
      "acceptHeader": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
      "deviceExperimentId": "ChxOekl3TVRVMk5ESTJNRGd6T0RrME16VTROZz09ELiq8aEGGIqxxJ8G",
      "screenWidthPoints": 1212,
      "screenHeightPoints": 626,
      "utcOffsetMinutes": -300,
      "userInterfaceTheme": "USER_INTERFACE_THEME_LIGHT",
      "clientScreen": "WATCH",
      "mainAppWebInfo": {
        "graftUrl": "/watch?v=coLCY15P6Bw",
        "webDisplayMode": "WEB_DISPLAY_MODE_BROWSER",
        "isWebNativeShareAvailable": false
      }
    },
    "user": {
      "lockedSafetyMode": false
    },
    "request": {
      "useSsl": true,
      "internalExperimentFlags": [],
      "consistencyTokenJars": []
    },
    "clickTracking": {
      "clickTrackingParams": "CKMBEKQwGAIiEwjx7-bama_-AhUltOUHHcFeCIYyB3JlbGF0ZWRIpsqRu__njoA0mgEFCAEQ-B0="
    },
    "adSignalsInfo": {
      "params": [
        {
          "key": "dt",
          "value": "1681675591916"
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
          "value": "1195"
        },
        {
          "key": "brdim",
          "value": "196,33,196,33,1536,0,1226,776,1212,626"
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
    }
  },
  "videoId": "coLCY15P6Bw",
  "playbackContext": {
    "contentPlaybackContext": {
      "currentUrl": "/watch?v=coLCY15P6Bw",
      "vis": 0,
      "splay": false,
      "autoCaptionsDefaultOn": false,
      "autonavState": "STATE_OFF",
      "html5Preference": "HTML5_PREF_WANTS",
      "signatureTimestamp": 19459,
      "referer": "https://www.youtube.com/watch?v=NAA7P_dkZSY",
      "lactMilliseconds": "-1",
      "watchAmbientModeContext": {
        "watchAmbientModeEnabled": true
      }
    }
  },
  "racyCheckOk": false,
  "contentCheckOk": false
}
`
