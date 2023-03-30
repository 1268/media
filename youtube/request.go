package youtube

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
)

func (r Request) Search(query string) (*Search, error) {
   filter := New_Filters()
   filter.Type(Type["Video"])
   r.Params = filter.Marshal()
   r.Query = query
   body, err := json.MarshalIndent(r, "", " ")
   if err != nil {
      return nil, err
   }
   req := http.Post(body)
   req.Header.Set("X-Goog-API-Key", api_key)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/search"
   req.URL.Scheme = "https"
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
func (r Request) Player(id string, tok *Token) (*Player, error) {
   r.Video_ID = id
   body, err := json.MarshalIndent(r, "", " ")
   if err != nil {
      return nil, err
   }
   req := http.Post(body)
   req.URL.Scheme = "https"
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   if tok != nil {
      req.Header.Set("Authorization", "Bearer " + tok.Access_Token)
   } else {
      req.Header.Set("X-Goog-API-Key", api_key)
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

const (
   // com.google.android.youtube
   android_version = "18.07.34"
   api_key = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
   mweb_version = "2.20230106.01.00"
   origin = "https://www.youtube.com"
)

func Android() Request {
   var r Request
   r.Content_Check_OK = true
   r.Context.Client.Name = "ANDROID"
   r.Context.Client.Version = android_version
   return r
}

func Android_Check() Request {
   var r Request
   r.Content_Check_OK = true
   r.Context.Client.Name = "ANDROID"
   r.Context.Client.Version = android_version
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
   r.Context.Client.Version = android_version
   return r
}

type Request struct {
   Content_Check_OK bool `json:"contentCheckOk,omitempty"`
   Context struct {
      Client struct {
         Name string `json:"clientName"`
         Version string `json:"clientVersion"`
      } `json:"client"`
   } `json:"context"`
   Params []byte `json:"params,omitempty"`
   Query string `json:"query,omitempty"`
   Racy_Check_OK bool `json:"racyCheckOk,omitempty"`
   Video_ID string `json:"videoId,omitempty"`
}

