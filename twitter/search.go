// Support for this software's development was paid for by
// Fredrick R. Brennan's Modular Font Editor K Foundation, Inc.
package twitter

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
)

func (x subtask) search(q string) (*search, error) {
   req, err := http.NewRequest(
      "GET", "https://api.twitter.com/2/search/adaptive.json", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "q": {q},
      "tweet_search_mode": {"live"},
   }.Encode()
   {
      o := oauth{
         consumer_key: consumer_key,
         consumer_secret: consumer_secret,
         token: x.Open_Account.OAuth_Token,
         token_secret: x.Open_Account.OAuth_Token_Secret,
      }
      req.Header.Set("Authorization", o.sign(req.Method, req.URL))
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
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
