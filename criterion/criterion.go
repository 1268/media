package criterion

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (f *File) License(data []byte) (*http.Response, error) {
   req, err := http.NewRequest(
      "POST", "https://drm.vhx.com/v2/widevine", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "token=" + f.DrmAuthorizationToken
   return http.DefaultClient.Do(req)
}

func (f *File) Mpd() (*http.Response, error) {
   return http.Get(f.Links.Source.Href)
}

type File struct {
   DrmAuthorizationToken string `json:"drm_authorization_token"`
   Links                 struct {
      Source struct {
         Href string
      }
   } `json:"_links"`
   Method string
}

const client_id = "9a87f110f79cd25250f6c7f3a6ec8b9851063ca156dae493bf362a7faf146c78"

func (t *Token) Unmarshal(data []byte) error {
   return json.Unmarshal(data, t)
}

func (f Files) Dash() (*File, bool) {
   for _, file0 := range f {
      if file0.Method == "dash" {
         return &file0, true
      }
   }
   return nil, false
}

func (Token) Marshal(username, password string) ([]byte, error) {
   resp, err := http.PostForm("https://auth.vhx.com/v1/oauth/token", url.Values{
      "client_id":  {client_id},
      "grant_type": {"password"},
      "password":   {password},
      "username":   {username},
   })
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type Token struct {
   AccessToken string `json:"access_token"`
}

func (t *Token) Video(slug string) (*Video, error) {
   req, _ := http.NewRequest("", "https://api.vhx.com", nil)
   req.URL.Path = "/videos/" + slug
   req.URL.RawQuery = "url=" + url.QueryEscape(slug)
   req.Header.Set("authorization", "Bearer "+t.AccessToken)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var video0 Video
   err = json.NewDecoder(resp.Body).Decode(&video0)
   if err != nil {
      return nil, err
   }
   if video0.Message != "" {
      return nil, errors.New(video0.Message)
   }
   return &video0, nil
}

type Video struct {
   Links struct {
      Files struct {
         Href string
      }
   } `json:"_links"`
   Message string
   Name string
}

func (t *Token) Files(video0 *Video) (Files, error) {
   req, err := http.NewRequest("", video0.Links.Files.Href, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer "+t.AccessToken)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var data strings.Builder
      resp.Write(&data)
      return nil, errors.New(data.String())
   }
   var files0 Files
   err = json.NewDecoder(resp.Body).Decode(&files0)
   if err != nil {
      return nil, err
   }
   return files0, nil
}

type Files []File
