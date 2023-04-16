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
      "clickTracking": {
         "clickTrackingParams": "CAAQu2kiEwiYg8OojK_-AhW2huUHHW_yANw="
      },
      "activePlayers": [
         {
            "playerContextParams": "Q0FFU0FnZ0I="
         }
      ]
   }
}
`
