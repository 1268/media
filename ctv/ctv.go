package ctv

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strconv"
   "strings"
)

// YOU CANNOT USE ANONYMOUS QUERY!
const query_axis = `
query axisContent($id: ID!) {
   axisContent(id: $id) {
      axisId
      axisPlaybackLanguages {
         ... on AxisPlayback {
            destinationCode
         }
      }
   }
}
`

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(data string) string {
   field := strings.Fields(data)
   return strings.Join(field, " ")
}

const query_resolve = `
query resolvePath($path: String!) {
   resolvedPath(path: $path) {
      lastSegment {
         content {
            ... on AxisObject {
               id
               ... on AxisMedia {
                  firstPlayableContent {
                     id
                  }
               }
            }
         }
      }
   }
}
`

func (a Address) Resolve() (*Content, error) {
   var value struct {
      Query         string `json:"query"`
      Variables     struct {
         Path string `json:"path"`
      } `json:"variables"`
   }
   value.Query = graphql_compact(query_resolve)
   value.Variables.Path = a.s
   data, err := json.MarshalIndent(value, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.ctv.ca/space-graphql/apq/graphql",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   // you need this for the first request, then can omit
   req.Header.Set("graphql-client-platform", "entpay_web")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   var value1 struct {
      Data struct {
         ResolvedPath *struct {
            LastSegment struct {
               Content Content
            }
         }
      }
   }
   err = json.Unmarshal(data, &value1)
   if err != nil {
      return nil, err
   }
   if path := value1.Data.ResolvedPath; path != nil {
      return &path.LastSegment.Content, nil
   }
   return nil, errors.New(string(data))
}

func (c *Content) Axis() (*AxisContent, error) {
   var value struct {
      Query         string `json:"query"`
      Variables     struct {
         Id string `json:"id"`
      } `json:"variables"`
   }
   value.Query = graphql_compact(query_axis)
   value.Variables.Id = c.get_id()
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.ctv.ca/space-graphql/apq/graphql",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   // you need this for the first request, then can omit
   req.Header.Set("graphql-client-platform", "entpay_web")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var value1 struct {
      Data struct {
         AxisContent AxisContent
      }
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value1)
   if err != nil {
      return nil, err
   }
   if err := value1.Errors; len(err) >= 1 {
      return nil, errors.New(err[0].Message)
   }
   return &value1.Data.AxisContent, nil
}

type Address struct {
   s string
}

// https://www.ctv.ca/shows/friends/the-one-with-the-bullies-s2e21
func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   a.s = strings.TrimPrefix(s, "ctv.ca")
   return nil
}

func (a *Address) String() string {
   return a.s
}

func (Widevine) Wrap(data []byte) ([]byte, error) {
   resp, err := http.Post(
      "https://license.9c9media.ca/widevine", "application/x-protobuf",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

// hard geo block
func (a *AxisContent) Manifest(media *MediaContent) (string, error) {
   req, err := http.NewRequest("", "https://capi.9c9media.com", nil)
   if err != nil {
      return "", err
   }
   req.URL.Path = func() string {
      b := []byte("/destinations/")
      b = append(b, a.AxisPlaybackLanguages[0].DestinationCode...)
      b = append(b, "/platforms/desktop/playback/contents/"...)
      b = strconv.AppendInt(b, a.AxisId, 10)
      b = append(b, "/contentPackages/"...)
      b = strconv.AppendInt(b, media.ContentPackages[0].Id, 10)
      b = append(b, "/manifest.mpd"...)
      return string(b)
   }()
   req.URL.RawQuery = "action=reference"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var data strings.Builder
      resp.Write(&data)
      return "", errors.New(data.String())
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return "", err
   }
   return strings.Replace(string(data), "/best/", "/ultimate/", 1), nil
}

func (m *MediaContent) Unmarshal(data []byte) error {
   return json.Unmarshal(data, m)
}

func (MediaContent) Marshal(axis *AxisContent) ([]byte, error) {
   req, _ := http.NewRequest("", "https://capi.9c9media.com", nil)
   req.URL.Path = func() string {
      b := []byte("/destinations/")
      b = append(b, axis.AxisPlaybackLanguages[0].DestinationCode...)
      b = append(b, "/platforms/desktop/contents/"...)
      b = strconv.AppendInt(b, axis.AxisId, 10)
      return string(b)
   }()
   req.URL.RawQuery = "$include=[ContentPackages,Media,Season]"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

func (c *Content) get_id() string {
   if c.FirstPlayableContent != nil {
      return c.FirstPlayableContent.Id
   }
   return c.Id
}

type Widevine struct{}

type Content struct {
   Id                   string
   FirstPlayableContent *struct {
      Id string
   }
}

///

type AxisContent struct {
   AxisId                int64
   AxisPlaybackLanguages []struct {
      DestinationCode string
   }
}

type MediaContent struct {
   ContentPackages []struct {
      Id int64
   }
   Episode int
   Media   struct {
      Name string
      Type string
   }
   Name   string
   Season struct {
      Number int
   }
}
