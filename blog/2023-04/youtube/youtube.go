package main

import (
   "2a.pages.dev/rosso/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "time"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   req_body := req_mes.Marshal()
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["X-Goog-Visitor-Id"] = []string{"Cgtkb0NEOGQzRTNiayi5wvyhBg%3D%3D"}
   for range [16]struct{}{} {
      req.Body = io.NopCloser(bytes.NewReader(req_body))
      res, err := new(http.Transport).RoundTrip(&req)
      if err != nil {
         panic(err)
      }
      res_body, err := io.ReadAll(res.Body)
      if err != nil {
         panic(err)
      }
      if err := res.Body.Close(); err != nil {
         panic(err)
      }
      res_mes, err := protobuf.Unmarshal(res_body)
      if err != nil {
         panic(err)
      }
      view_count, err := res_mes.Get(11).Get_String(32)
      if err != nil {
         panic(err)
      }
      adaptive_formats := res_mes.Get(4).Get_Messages(3)
      fmt.Println(view_count, len(adaptive_formats))
      time.Sleep(time.Second)
   }
}

var req_mes = protobuf.Message{
   4: protobuf.Message{
      1: protobuf.Message{
         3: protobuf.String("android-google"),
         5:  protobuf.Varint(496),
         6:  protobuf.Varint(0),
         8:  protobuf.Varint(0),
         10: protobuf.Varint(0),
         38: protobuf.Varint(0),
         4:  protobuf.Varint(0),
         12: protobuf.String("sdkv=a.18.14.40&output=xml_vast2"),
         7:  protobuf.Varint(3),
         29: protobuf.Varint(0),
         37: protobuf.Varint(0),
         11: protobuf.Varint(0),
         41: protobuf.Varint(0),
      },
   },
   5:  protobuf.Varint(0),
   8:  protobuf.Varint(0),
   15: protobuf.Varint(0),
   23: protobuf.String("eMNvk_p-rPhpyW5b"),
   1:protobuf.Message{
      1:protobuf.Message{
         19: protobuf.String("8.0.0"),
         37: protobuf.Varint(384),
         50: protobuf.Varint(231312022),
         39: protobuf.Fixed32(1075419546),
         41: protobuf.Varint(2),
         46: protobuf.Varint(1),
         55: protobuf.Varint(384),
         61: protobuf.Varint(3),
         78: protobuf.Varint(1),
         13: protobuf.String("Phone"),
         38: protobuf.Varint(592),
         92: protobuf.String("vbox86;vbox86"),
         98: protobuf.String("Custom"),
         12: protobuf.String("Genymobile"),
         25: protobuf.String("4486089941071504633"),
         56: protobuf.Varint(592),
         16: protobuf.Varint(3),
         18: protobuf.String("Android"),
         40: protobuf.Fixed32(1080872141),
         64: protobuf.Varint(26),
         100: protobuf.Message{
            1: protobuf.Message{
               1: protobuf.Varint(1658),
               3: protobuf.Varint(1),
            },
         },
         21: protobuf.String("en-US"),
         22: protobuf.String("US"),
         65: protobuf.Fixed32(1073741824),
         80: protobuf.String("GMT"),
         102: protobuf.Message{
            1: protobuf.String("Unknown Renderer"),
            2: protobuf.Varint(2),
            3: protobuf.Varint(0),
         },
         17: protobuf.String("18.14.40"),
         52: protobuf.Varint(4),
         67: protobuf.Varint(0),
      },
      3: protobuf.Message{
         7:  protobuf.Varint(0),
         15: protobuf.Varint(0),
      },
      6: protobuf.Message{
         2: protobuf.Message{
            4: protobuf.Message{
               1: protobuf.Varint(1681858873523252),
               2: protobuf.Fixed32(183321733),
               3: protobuf.Fixed32(2768959010),
            },
         },
         6: protobuf.String("external"),
      },
   },
   9: protobuf.Message{
      1: protobuf.Message{
         1: protobuf.String("ms"),
         2: protobuf.String("CoACYi-fyjXIPhEg5_hlafLdmGV94aRDuMNQvupS-2SPkB-jS9LZXDLwFpht8BEePWnIky57ZMXZx4t1F7MF7ybQcAKvq70s8FLMbE3TICFK-WJ_s3mFyIFNsapW3M41BjishdbCA53kL_GY35c-9gMz3BBvzb72RUjeAtl8M__GfUId4EfTyel2OW-A0XR7fweioyEoA18EDfCdFBXO37Rzn9TOr96mo5X5fuatc9xD_7e3_YrD-NwpwQYtFP5iDnrWK9_U4scGDuTf3sKGqX013HeF_uAgjiWdHDD3LlOhocZ71hVm7Gn1UR5DTJZP5VBz16t7X-11q43NatXWlXWy8wqAAkQBofSFp639YmOCtjX8ZK3P4JTDVqGml4xVVgbhRWn-rZhjn2oSNUN1vWcvNsEdyVEK6XvEqtgFmh7MPJuR-baBs4AAvKj2xKJ05b03QpowR9oipEWg-XBUx1UjzGab9cJ35Nc1_s81emfhzBswwY9KheOUHIa_qM-Am3maHB2Hyx_L2hM40AubVM8lAousZ6Y_H-w31QfbZ8vklWpA3zN29Yy46xbFP7ugSD359f5GxvqdYAXmQ_SUWotgM_0cwuuF4Wfdy9qUg8JNmngL3s95WP5SjWv9EjgXIgdUyhu7QFJ5iAS84r7oijjqIIuAO_QU2KzD4jmaxnYB3RKRJw0SELgzlk9NzTrOAGrLcdHQ6RE"),
      },
   },
   2: protobuf.String("Xxk-ryO6J2I"),
   3: protobuf.Varint(0),
}
