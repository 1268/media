package youtube

import (
   "bytes"
   "encoding/json"
   "net/http"
   "net/url"
   "path"
)

type Request struct {
   Content_Check_OK bool `json:"contentCheckOk,omitempty"`
   Context struct {
      Client struct {
         Android_SDK_Version int `json:"androidSdkVersion"`
         Client_Name string `json:"clientName"`
         Client_Version string `json:"clientVersion"`
         // need this to get the correct:
         // This video requires payment to watch
         // instead of the invalid:
         // This video can only be played on newer versions of Android or other
         // supported devices.
         OS_Version string `json:"osVersion"`
      } `json:"client"`
   } `json:"context"`
   Racy_Check_OK bool `json:"racyCheckOk,omitempty"`
   Video_ID string `json:"videoId,omitempty"`
}

func (r Request) Player(tok *Token) (*Player, error) {
   r.Context.Client.Android_SDK_Version = 32
   r.Context.Client.OS_Version = "12"
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
   req.Header.Set("User-Agent", user_agent + r.Context.Client.Client_Version)
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

func (r *Request) Android() {
   r.Content_Check_OK = true
   r.Context.Client.Client_Name = "ANDROID"
   r.Context.Client.Client_Version = android_youtube
}

func (r *Request) Android_Check() {
   r.Content_Check_OK = true
   r.Context.Client.Client_Name = "ANDROID"
   r.Context.Client.Client_Version = android_youtube
   r.Racy_Check_OK = true
}

func (r *Request) Android_Embed() {
   r.Context.Client.Client_Name = "ANDROID_EMBEDDED_PLAYER"
   r.Context.Client.Client_Version = android_youtube
}

const android_youtube = "18.39.41"

const user_agent = "com.google.android.youtube/"

func (r *Request) Mobile_Web() {
   r.Context.Client.Client_Name = "MWEB"
   r.Context.Client.Client_Version = mweb_version
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
