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
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   val := make(url.Values)
   val["key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
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
