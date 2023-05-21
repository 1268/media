// Support for this software's development was paid for by
// Fredrick R. Brennan's Modular Font Editor K Foundation, Inc.
package twitter

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "net/url"
)

type Guest struct {
   Guest_Token string
}

func New_Guest() (*Guest, error) {
   req := http.Post(&url.URL{
      Scheme: "https",
      Host: "api.twitter.com",
      Path: "/1.1/guest/activate.json",
   })
   req.Header.Set("Authorization", "Bearer " + bearer)
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   g := new(Guest)
   if err := json.NewDecoder(res.Body).Decode(g); err != nil {
      return nil, err
   }
   return g, nil
}

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

func (g Guest) Source(space *Audio_Space) (*Source, error) {
   req := http.Get(&url.URL{
      Scheme: "https",
      Host: "twitter.com",
      Path: "/i/api/1.1/live_video_stream/status/" + space.Metadata.Media_Key,
   })
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {g.Guest_Token},
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var video struct {
      Source Source
   }
   if err := json.NewDecoder(res.Body).Decode(&video); err != nil {
      return nil, err
   }
   return &video.Source, nil
}

func (g Guest) Space(id string) (*Audio_Space, error) {
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
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {g.Guest_Token},
   }
   res, err := http.Default_Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var space struct {
      Data struct {
         Audio_Space Audio_Space `json:"audioSpace"`
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&space); err != nil {
      return nil, err
   }
   return &space.Data.Audio_Space, nil
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
