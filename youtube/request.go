package youtube

import (
   "bytes"
   "encoding/json"
   "net/http"
   "net/url"
   "path"
   "strconv"
)

func (r Request) Search(query string) (*Search, error) {
   body, err := func() ([]byte, error) {
      var p parameters
      p.Filter = new(filter)
      p.Filter.Type = values["TYPE"]["Video"]
      r.Params = p.Marshal()
      r.Query = query
      return json.MarshalIndent(r, "", " ")
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/search",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("X-Goog-API-Key", api_key)
   res, err := http.DefaultClient.Do(req)
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

func (v version) String() string {
   var b []byte
   b = strconv.AppendInt(b, v.major, 10)
   b = append(b, '.')
   b = strconv.AppendInt(b, v.minor, 10)
   b = append(b, '.')
   b = strconv.AppendInt(b, v.patch, 10)
   return string(b)
}

const user_agent = "com.google.android.youtube/"

func (r Request) Player(tok *Token) (*Player, error) {
   r.Context.Client.Android_SDK_Version = 99
   body, err := json.MarshalIndent(r, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("User-Agent", user_agent + r.Context.Client.Version)
   if tok != nil {
      req.Header.Set("Authorization", "Bearer " + tok.Access_Token)
   }
   res, err := http.DefaultClient.Do(req)
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

func (r *Request) Android() {
   r.Content_Check_OK = true
   r.Context.Client.Name = "ANDROID"
   r.Context.Client.Version = max_android.String()
}

func (r *Request) Android_Check() {
   r.Content_Check_OK = true
   r.Context.Client.Name = "ANDROID"
   r.Context.Client.Version = max_android.String()
   r.Racy_Check_OK = true
}

func (r *Request) Mobile_Web() {
   r.Context.Client.Name = "MWEB"
   r.Context.Client.Version = mweb_version
}

func (r *Request) Android_Embed() {
   r.Context.Client.Name = "ANDROID_EMBEDDED_PLAYER"
   r.Context.Client.Version = max_android.String()
}

func (r *Request) Set(s string) error {
   ref, err := url.Parse(s)
   if err != nil {
      return err
   }
   r.Video_ID = ref.Query().Get("v")
   if r.Video_ID == "" {
      r.Video_ID = path.Base(ref.Path)
   }
   return nil
}

func (r Request) String() string {
   return r.Video_ID
}

const (
   api_key = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
   mweb_version = "2.20230405.01.00"
)
