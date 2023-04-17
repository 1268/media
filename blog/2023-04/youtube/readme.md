# YouTube

<https://github.com/glubsy/livestream_saver/issues/63>

## Visitor

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
  }
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
`VISITOR_INFO1_LIVE` cookie.

~~~
POST /youtubei/v1/player HTTP/1.1
Host: www.youtube.com
X-Goog-Api-Key: AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8

{
 "contentCheckOk": true,
 "context": {
  "client": {
   "clientName": "ANDROID",
   "clientVersion": "18.14.40"
  }
 },
 "params": "CgIQAA==",
 "videoId": "-Lknlh0Qib0"
}
~~~
