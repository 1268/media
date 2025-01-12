package rakuten

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

func (a *address) gizmo_movie() (*gizmo_movie, error) {
   req, err := http.NewRequest("", "https://gizmo.rakuten.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v3/")
      if a.movie != "" {
         b.WriteString("movies/")
         b.WriteString(a.movie)
      } else {
         b.WriteString("seasons/")
         b.WriteString(a.season)
      }
      return b.String()
   }()
   req.URL.RawQuery = url.Values{
      "classification_id": {
         strconv.Itoa(classification_id[a.market_code]),
      },
      "device_identifier": {"atvui40"},
      "market_code":       {a.market_code},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   movie := &gizmo_movie{}
   err = json.NewDecoder(resp.Body).Decode(movie)
   if err != nil {
      return nil, err
   }
   return movie, nil
}

type gizmo_movie struct {
   Data struct {
      Title string
      Year  int
   }
}

func streamings() (*http.Response, error) {
   var body = strings.NewReader(`
   {
      "audio_quality": "2.0",
      "classification_id": "272",
      "content_id": "transvulcania-the-people-s-run",
      "content_type": "movies",
      "video_type": "stream",
      "subtitle_language": "MIS",
      "device_serial": "not implemented",
      "device_identifier": "atvui40",
      "player": "atvui40:DASH-CENC:WVM",
      "device_stream_video_quality": "FHD",
      
      "audio_language": "SPA"
   }
   `)
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.URL = &url.URL{}
   req.URL.Host = "gizmo.rakuten.tv"
   req.URL.Path = "/v3/avod/streamings"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/json"}
   return http.DefaultClient.Do(&req)
}
