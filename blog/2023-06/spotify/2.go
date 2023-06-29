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
   req.URL.Host = "stitcher2.acast.com"
   req.URL.Path = "/livestitches/558fa97de2e566cc7877a838a1ee27d7.mp3"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["Signature"] = []string{"BAp-STHRCqIzKfmO2v80CLLFd20PMhnjJpLqppVS7Bpg6AMYio7iksyOYuDcSkv5Tf6iomIQqYPJrIko35Zeb~czbdV3tEEkXSgammSwl6pO0LQIH7gO0goALWk8-hcTGzzsoviPdT9TzXQEpOCsq72nCWAZBZrhiOP1zShtk0eFU6CaaI3yhTdNC24TO78iMQ4r5uPByMWM8zmN3c0Nwl9WPVkG7A93t4SCFsQM0~xCSiw-3-Xeld0hER85Wmm~NS2l3eD8LQVkwm6oYRbv7KqcZEG7lplf-2bVMyvlspXGGmvAPTdW82JItQ0VaxTiS22WZDVVpH82nQlT-iJ3qA__"}
   val["Expires"] = []string{"1688013298570"}
   val["Key-Pair-Id"] = []string{"K38CTQXUSD0VVB"}
   val["aid"] = []string{"64620c6df695d60011f6fa45"}
   val["chid"] = []string{"8c0337e1-761f-512d-b221-8e10f115da57"}
   val["ci"] = []string{"AKYxauximxGCVJcQ08-YfStSrseuCNCre4Zglht9f1TBndqqQAHAZg=="}
   val["pf"] = []string{"rss"}
   val["range"] = []string{"bytes=0-"}
   val["sv"] = []string{"sphinx@1.162.1"}
   val["uid"] = []string{"246a78f692da14b9878a825ff506bb1e"}
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, false)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}
