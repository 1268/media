package twitter

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
)

type tweet_result struct {
   Video struct {
      Variants []struct {
         Src string
      }
   }
}

func new_tweet_result(id int64) (*tweet_result, error) {
   req, err := http.NewRequest("GET", "https://cdn.syndication.twimg.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/tweet-result"
   req.URL.RawQuery = url.Values{
      "id": {strconv.FormatInt(id, 10)},
      "token": {"-"},
      //token=!
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tweet := new(tweet_result)
   if err := json.NewDecoder(res.Body).Decode(tweet); err != nil {
      return nil, err
   }
   return tweet, nil
}
