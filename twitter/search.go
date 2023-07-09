// Support for this software's development was paid for by
// Fredrick R. Brennan's Modular Font Editor K Foundation, Inc.
package twitter

import (
   "encoding/json"
   "net/url"
)

func (x subtask) search(q string) (*search, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/2/search/adaptive.json",
      RawQuery: url.Values{
         "q": {q},
         "tweet_search_mode": {"live"},
      }.Encode(),
   })
   auth := oauth{
      consumer_key: consumer_key,
      consumer_secret: consumer_secret,
      token: x.Open_Account.OAuth_Token,
      token_secret: x.Open_Account.OAuth_Token_Secret,
   }
   req.Header.Set("Authorization", auth.sign(req.Method, req.URL))
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   s := new(search)
   if err := json.NewDecoder(res.Body).Decode(s); err != nil {
      return nil, err
   }
   return s, nil
}

type search struct {
   GlobalObjects struct {
      Tweets map[int64]struct {
         Entities struct {
            URLs []struct {
               Expanded_URL string
            }
         }
      }
   }
}
