package main

import (
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
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
   req.Body = io.NopCloser(req_body)
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if bytes.Contains(body, []byte(`"viewCount"`)) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

var req_body = strings.NewReader(`
{
   "context": {
      "client": {
         "hl": "en",
         "gl": "US",
         "remoteHost": "72.181.23.38",
         "deviceMake": "",
         "deviceModel": "",
         "visitorData": "CgtCYTlxQUNNUmJtOCjx1_ChBg%3D%3D",
         "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0,gzip(gfe)",
         "clientName": "WEB",
         "clientVersion": "2.20230414.01.00",
         "osName": "Windows",
         "osVersion": "10.0",
         "originalUrl": "https://www.youtube.com/watch?v=k5dX9sjXYVk",
         "screenPixelDensity": 1,
         "platform": "DESKTOP",
         "clientFormFactor": "UNKNOWN_FORM_FACTOR",
         "configInfo": {
            "appInstallData": "CPHX8KEGEO2GrwUQ5LP-EhDn964FEOLUrgUQuIuuBRDM9a4FEOmhrwUQvbauBRCgt_4SEInorgUQzK7-EhCi7K4FENburgUQ1_-uBRCFg68FEOWg_hIQpJmvBRC24K4FELeRrwUQuNSuBRD-7q4FEMzfrgUQuI-vBRCY2q4FEM6ZrwU%3D"
         },
         "screenDensityFloat": 1.25,
         "timeZone": "America/Chicago",
         "browserName": "Firefox",
         "browserVersion": "90.0",
         "acceptHeader": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
         "deviceExperimentId": "ChxOekl3TVRVMk5ESTJNRGd6T0RrME16VTROZz09EPHX8KEGGIqxxJ8G",
         "screenWidthPoints": 1212,
         "screenHeightPoints": 626,
         "utcOffsetMinutes": -300,
         "userInterfaceTheme": "USER_INTERFACE_THEME_LIGHT",
         "mainAppWebInfo": {
            "graftUrl": "https://www.youtube.com/watch?v=k5dX9sjXYVk",
            "webDisplayMode": "WEB_DISPLAY_MODE_BROWSER",
            "isWebNativeShareAvailable": false
         },
         "playerType": "UNIPLAYER",
         "tvAppInfo": {
            "livingRoomAppMode": "LIVING_ROOM_APP_MODE_UNSPECIFIED"
         },
         "clientScreen": "WATCH_FULL_SCREEN"
      },
      "request": {
         "useSsl": true,
         "internalExperimentFlags": [],
         "consistencyTokenJars": []
      },
      "user": {
         "lockedSafetyMode": false
      }
   },
   "videoId": "k5dX9sjXYVk"
}
`)
