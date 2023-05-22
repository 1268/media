// Support for this software's development was paid for by
// Fredrick R. Brennan's Modular Font Editor K Foundation, Inc.
package twitter

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
)

const persisted_query = "lFpix9BgFDhAMjn9CrW6jQ"

type Audio_Space struct {
   Metadata struct {
      Ended_At int64 `json:"ended_at,string"`
      Media_Key string
      Started_At int64
      State string
      Title string
   }
   Participants struct {
      Admins []struct {
         Display_Name string
      }
   }
}

type Source struct {
   Location string // Segment
}

type space_request struct {
   ID string `json:"id"`
   Is_Metatags_Query bool `json:"isMetatagsQuery"`
   With_Birdwatch_Pivots bool `json:"withBirdwatchPivots"`
   With_Downvote_Perspective bool `json:"withDownvotePerspective"`
   With_Reactions_Metadata bool `json:"withReactionsMetadata"`
   With_Reactions_Perspective bool `json:"withReactionsPerspective"`
   With_Replays bool `json:"withReplays"`
   With_Scheduled_Spaces bool `json:"withScheduledSpaces"`
   With_Super_Follows_Tweet_Fields bool `json:"withSuperFollowsTweetFields"`
   With_Super_Follows_User_Fields bool `json:"withSuperFollowsUserFields"`
}

func (h header) Space(id string) (*Audio_Space, error) {
   query := func(r *http.Request) error {
      b, err := json.Marshal(space_request{ID: id})
      if err != nil {
         return err
      }
      r.URL.RawQuery = "variables=" + url.QueryEscape(string(b))
      return nil
   }
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "twitter.com",
      Path: "/i/api/graphql/" + persisted_query + "/AudioSpaceById",
   })
   err := query(req)
   if err != nil {
      return nil, err
   }
   req.Header = h.Header
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      Data struct {
         Audio_Space Audio_Space `json:"audioSpace"`
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   return &s.Data.Audio_Space, nil
}

func (h header) Source(space *Audio_Space) (*Source, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "twitter.com",
      Path: "/i/api/1.1/live_video_stream/status/" + space.Metadata.Media_Key,
   })
   req.Header = h.Header
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      Source Source
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   return &s.Source, nil
}
