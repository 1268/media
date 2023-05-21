package twitter

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
)

func Token() (*http.Response, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/oauth2/token",
   })
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.SetBasicAuth(
      "3nVuSoBZnx6U4vzUxf5w",
      "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys",
   )
   req.Body_String("grant_type=client_credentials")
   return http.Default_Client.Do(req)
}

func (x subtask) search(q string) (*search, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/2/search/adaptive.json",
      RawQuery: url.Values{
         "q": {q},
         // This ensures Spaces Tweets will include Spaces URL
         "tweet_mode": {"extended"},
      }.Encode(),
   })
   auth := oauth{
      consumer_key: "3nVuSoBZnx6U4vzUxf5w",
      consumer_secret: "Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys",
      token: x.Open_Account.OAuth_Token,
      token_secret: x.Open_Account.OAuth_Token_Secret,
   }
   req.Header["Authorization"] = []string{auth.sign(req.Method, req.URL)}
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

func New_Search(q string) (*Search, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/2/search/adaptive.json",
      RawQuery: url.Values{
         "q": {q},
         // This ensures Spaces Tweets will include Spaces URL
         "tweet_mode": {"extended"},
      }.Encode(),
   })
   req.Header.Set("Authorization", "Bearer " + bearer)
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   s := new(Search)
   if err := json.NewDecoder(res.Body).Decode(s); err != nil {
      return nil, err
   }
   return s, nil
}

type Search struct {
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
