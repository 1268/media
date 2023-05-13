package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "services.radio-canada.ca"
   req.URL.Path = "/ott/subscription/v2/gem/Subscriber/profile"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["device"] = []string{"phone_android"}
   req.URL.RawQuery = val.Encode()
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IjkzQURGMUNFNDhGNENCQUVCOTBDM0YyNEU5RDc0QkU5RjU0REJDMTIiLCJ4NXQiOiJrNjN4emtqMHk2NjVERDhrNmRkTDZmVk52QkkiLCJ0eXAiOiJKV1QifQ.eyJhenBDb250ZXh0IjoiY2JjZ2VtIiwiZ2l2ZW5fbmFtZSI6InN0ZXZlbiIsImZhbWlseV9uYW1lIjoicGVubnkiLCJuYW1lIjoic3RldmVuIHBlbm55Iiwic3ViIjoiZmE4N2M1NDAtMDRlNy00NDYxLTk3N2EtYzhiZDQ3MGQxNDBhIiwicmNpZCI6ImZhODdjNTQwLTA0ZTctNDQ2MS05NzdhLWM4YmQ0NzBkMTQwYSIsInBpY3R1cmUiOiIiLCJlbWFpbCI6InNycGVuNkBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaWRwVXNlcklkIjoiMGY2NTBmOGQyMjBkNDEyYmFhYTU3OGQ1YzUwYjk2ZWQiLCJscmF0IjoiNjQ1ZmVlMzMtMmE0Mi00YjhhLWE3ZWQtYzg1ODBiNmUwMjExIiwib2lkIjoiOWQxZDBlNTMtYWVjZi00ODUxLTgxZTYtM2RmMzJhN2E0NTRkIiwiaWRwIjoicmFkaW9jYW5hZGEiLCJzY3AiOiJlbWFpbCBtZWRpYS1kcm10IG1lZGlhLW1ldGEgbWVkaWEtdmFsaWRhdGlvbiBtZXRyaWsgb2lkYzRyb3BjIHByb2ZpbGUgdG91dHYgdG91dHYtcHJlc2VudGF0aW9uIHRvdXR2LXByb2ZpbGluZyBpZC5hY2NvdW50LmNyZWF0ZSBpZC5hY2NvdW50LmRlbGV0ZSBpZC5hY2NvdW50LmluZm8gaWQuYWNjb3VudC5tb2RpZnkgaWQuYWNjb3VudC5yZXNldC1wYXNzd29yZCBpZC5hY2NvdW50LnNlbmQtY29uZmlybWF0aW9uLWVtYWlsIGlkLndyaXRlIG1lZGlhLXZhbGlkYXRpb24ucmVhZCBzdWJzY3JpcHRpb25zLnZhbGlkYXRlIHN1YnNjcmlwdGlvbnMud3JpdGUgb3R0LXByb2ZpbGluZyBvdHQtc3Vic2NyaXB0aW9uIiwiYXpwIjoiN2Y0NGM5MzUtNjU0Mi00Y2U3LWFlMDUtZWI4ODc4MDk3NDFjIiwidmVyIjoiMS4wIiwiaWF0IjoxNjg0MDA4NTAwLCJhdWQiOiI4NDU5M2I2NS0wZWY2LTRhNzItODkxYy1kMzUxZGRkNTBhYWIiLCJleHAiOjE2ODQwMzAxMDAsImlzcyI6Imh0dHBzOi8vcmNtbmIyY3Byb2QuYjJjbG9naW4uY29tL2JlZjFiNTM4LTE5NTAtNDI4My05YjI3LWIwOTZjYmMxODA3MC92Mi4wLyIsIm5iZiI6MTY4NDAwODUwMH0.RshE1crQpQ8NfTmgBqMJcCOdlgKpZ4wTqdzAYsvtRB-gNWseJSNAuBRh21QTqtT5tJZfRIgGVgtfvymzWhLwuUG1GP1ItaYmfVt3uhjPt6itUCDL38nKwDhBddTVP1ehoKTSt4Uip7By0ifB62MFFcbZiPxAWHQzb9gMknPHtacKMnzDihLZoerf59ZrBQ9-yhBPOYurovTWeh6chHWPQ5wBiQGt_qHoqv7oWMJYA30hj_VVsIDTzaGV_Sy_o397v3LnWkyyzJasCo2JjPEzh78EeIgFMPBGfwxGFFkfzV4NIBPcj5Hrl6ddquyN873xjNJL1homppvOgLPp2XZOiw"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}
