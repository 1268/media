package youtube

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/json"
   "io"
   "net/url"
   "strconv"
)

func new_config() (*config, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "m.youtube.com",
   })
   req.Header.Set("User-Agent", "iPad")
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   sep := []byte("\nytcfg.set(")
   con := new(config)
   if err := json.Cut(data, sep, con); err != nil {
      return nil, err
   }
   return con, nil
}

const (
   api_key = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
   mweb_version = "2.20230405.01.00"
)

func (r Request) Search(query string) (*Search, error) {
   body := func(req *http.Request) error {
      p := New_Params()
      p.Type(Type["Video"])
      r.Params = p.Marshal()
      r.Query = query
      b, err := json.MarshalIndent(r, "", " ")
      if err != nil {
         return err
      }
      req.Body_Bytes(b)
      return nil
   }
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "www.youtube.com",
      Path: "/youtubei/v1/search",
   })
   req.Header.Set("X-Goog-API-Key", api_key)
   err := body(req)
   if err != nil {
      return nil, err
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   search := new(Search)
   if err := json.NewDecoder(res.Body).Decode(search); err != nil {
      return nil, err
   }
   return search, nil
}

type version struct {
   major int64
   minor int64
   patch int64
}

var max_android = version{18, 22, 99}

func Android() Request {
   var r Request
   r.Content_Check_OK = true
   r.Context.Client.Name = "ANDROID"
   r.Context.Client.Version = max_android.String()
   return r
}

type Request struct {
   Content_Check_OK bool `json:"contentCheckOk,omitempty"`
   Context struct {
      Client struct {
         Android_SDK_Version int32 `json:"androidSdkVersion,omitempty"`
         Name string `json:"clientName"`
         Version string `json:"clientVersion"`
      } `json:"client"`
   } `json:"context"`
   Params []byte `json:"params,omitempty"`
   Query string `json:"query,omitempty"`
   Racy_Check_OK bool `json:"racyCheckOk,omitempty"`
   Video_ID string `json:"videoId,omitempty"`
}

func (v version) String() string {
   var b []byte
   b = strconv.AppendInt(b, v.major, 10)
   b = append(b, '.')
   b = strconv.AppendInt(b, v.minor, 10)
   b = append(b, '.')
   b = strconv.AppendInt(b, v.patch, 10)
   return string(b)
}


func (r Request) Player(id string, tok *Token) (*Player, error) {
   body := func(req *http.Request) error {
      r.Context.Client.Android_SDK_Version = 99
      r.Video_ID = id
      b, err := json.MarshalIndent(r, "", " ")
      if err != nil {
         return err
      }
      req.Body_Bytes(b)
      return nil
   }
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "www.youtube.com",
      Path: "/youtubei/v1/player",
   })
   req.Header.Set("User-Agent", user_agent + r.Context.Client.Version)
   if tok != nil {
      req.Header.Set("Authorization", "Bearer " + tok.Access_Token)
   }
   err := body(req)
   if err != nil {
      return nil, err
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   play := new(Player)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

const user_agent = "com.google.android.youtube/"

func Android_Check() Request {
   var r Request
   r.Content_Check_OK = true
   r.Context.Client.Name = "ANDROID"
   r.Context.Client.Version = max_android.String()
   r.Racy_Check_OK = true
   return r
}

func Mobile_Web() Request {
   var r Request
   r.Context.Client.Name = "MWEB"
   r.Context.Client.Version = mweb_version
   return r
}

func Android_Embed() Request {
   var r Request
   r.Context.Client.Name = "ANDROID_EMBEDDED_PLAYER"
   r.Context.Client.Version = max_android.String()
   return r
}

type config struct {
   Innertube_API_Key string
   Innertube_Client_Name string
   Innertube_Client_Version string
}
