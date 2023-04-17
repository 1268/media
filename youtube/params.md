# params

If you visit this page:

https://www.youtube.com

and hover over one of the videos, you should capture a request like this:

~~~
POST https://www.youtube.com/youtubei/v1/player?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8&prettyPrint=false HTTP/2.0
cookie: VISITOR_INFO1_LIVE=Ba9qACMRbm8
x-goog-visitor-id: CgtCYTlxQUNNUmJtOCiE8fahBg%3D%3D

{
  "context": {
    "client": {
      "visitorData": "CgtCYTlxQUNNUmJtOCiE8fahBg%3D%3D"
    }
  },
  "params": "YAHIAQE%3D"
}
~~~

lets start with the `visitorData`/`x-goog-visitor-id` value. we can unmarshal
the data:

~~~go
package main

import (
   "2a.pages.dev/rosso/protobuf"
   "encoding/base64"
   "fmt"
   "net/url"
)

func main() {
   s, err := url.QueryUnescape("CgtCYTlxQUNNUmJtOCiE8fahBg%3D%3D")
   if err != nil {
      panic(err)
   }
   b, err := base64.StdEncoding.DecodeString(s)
   if err != nil {
      panic(err)
   }
   m, err := protobuf.Unmarshal(b)
   if err != nil {
      panic(err)
   }
   fmt.Println(m)
}
~~~

result:

~~~go
protobuf.Message{
   1: protobuf.String("Ba9qACMRbm8"),
   5: protobuf.Varint(1681766532),
}
~~~

value `5` is just the current timestamp:

~~~go
package main

import "time"

func main() {
   s := time.Unix(1681766532, 0).String()
   println(s) // 2023-04-17 16:22:12 -0500 CDT
}
~~~

I have found that it can be omitted, which leaves:

~~~go
protobuf.Message{
   1: protobuf.String("Ba9qACMRbm8"),
}
~~~

whats interesting if you look above, this string also matches the
`VISITOR_INFO1_LIVE` cookie. If we marshal the message, we get:

~~~go
package main

import (
   "2a.pages.dev/rosso/protobuf"
   "encoding/base64"
)

func main() {
   b := protobuf.Message{
      1: protobuf.String("Ba9qACMRbm8"),
   }.Marshal()
   println(base64.StdEncoding.EncodeToString(b)) // CgtCYTlxQUNNUmJtOA==
}
~~~

if we plug that into the request body, everything works as expected:

~~~go
package main

import (
   "io"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["X-Goog-Api-Key"] = []string{"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"}
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   req_body := `
   {
    "context": {
     "client": {
      "clientName": "ANDROID",
      "clientVersion": "18.14.40",
      "visitorData": "CgtCYTlxQUNNUmJtOCiE8fahBg%3D%3D"
     }
    },
    "videoId": "sG8DXQhNEGU"
   }
   `
   for range [16]struct{}{} {
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
      text := string(body)
      println(
         "adaptiveFormats", strings.Contains(text, `"adaptiveFormats"`),
         "viewCount", strings.Contains(text, `"viewCount"`),
      )
      time.Sleep(time.Second)
   }
}
~~~

If we remove `visitorData`, we get random failures:

~~~
adaptiveFormats true viewCount true
adaptiveFormats true viewCount true
adaptiveFormats true viewCount true
adaptiveFormats true viewCount true
adaptiveFormats false viewCount false
adaptiveFormats true viewCount true
adaptiveFormats true viewCount true
adaptiveFormats true viewCount true
adaptiveFormats false viewCount false
adaptiveFormats false viewCount false
adaptiveFormats true viewCount true
adaptiveFormats true viewCount true
adaptiveFormats true viewCount true
adaptiveFormats true viewCount true
adaptiveFormats false viewCount false
adaptiveFormats false viewCount false
~~~

note setting the `x-goog-visitor-id` header instead is also an option:

~~~go
req.Header = http.Header{
   "x-goog-api-key": {"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"},
   "x-goog-visitor-id": {"CgtCYTlxQUNNUmJtOCiE8fahBg%3D%3D"},
}
~~~

but the `cookie` header also fails randomly for some reason:

~~~go
req.Header = http.Header{
   "cookie": {"VISITOR_INFO1_LIVE=Ba9qACMRbm8"},
   "x-goog-api-key": {"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"},
}
~~~

note trying `params` instead of `visitorData`/`x-goog-visitor-id` as in the
original request:

~~~
"params": "YAHIAQE%3D"
~~~

fails randomly as well. Note I also found this interesting comment:

> the required parameter is field number 12 in the protobuf (very long base64
> string), the JSON name is "params" -- it seems to be some kind of ad token?
>
> strangely even though using the JSON API without the parameter gets you a 403,
> when using the protobuf interface you get the famous 'content is not available
> on this app' message instead.
>
> edit: since params contains protobuf-encoded data, removing all the fields
> except the ones the server needs gets you the following working (for now)
> value: CgIQBg%3D%3D

https://github.com/TeamNewPipe/NewPipe/issues/9038#issuecomment-1289756816

Sadly this person didnt explain what "field number 12 in the protobuf" is. they
could mean part of a web client HTTP request or response body, which was then
base64 decoded and then ProtoBuf unmarshaled. or they could mean part of an
Android client HTTP request or response body. At any rate, adding that value to
the request body:

~~~
"params": "CgIQBg%3D%3D"
~~~

does fix the issue. in addition, if you unmarshal the value, you get:

~~~go
protobuf.Message{
   1: protobuf.Message{
      2: protobuf.Varint(6)
   }
}
~~~

I found you can replace the Varint with any value and the request still works,
so here are the encoded versions:

~~~
CgIQAA==
CgIQAQ==
CgIQAg==
CgIQAw==
CgIQBA==
CgIQBQ==
CgIQBg==
CgIQBw==
CgIQCA==
CgIQCQ==
~~~
