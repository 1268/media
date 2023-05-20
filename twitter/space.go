// Support for this software's development was paid for by
// Fredrick R. Brennan's Modular Font Editor K Foundation, Inc.
package twitter

import (
   "2a.pages.dev/rosso/http"
   "encoding/json"
   "errors"
   "net/url"
   "path"
   "strings"
   "time"
)

func NewGuest() (*Guest, error) {
   req, err := http.NewRequest(
      "POST", "https://api.twitter.com/1.1/guest/activate.json", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Authorization", "Bearer " + bearer)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   guest := new(Guest)
   if err := json.NewDecoder(res.Body).Decode(guest); err != nil {
      return nil, err
   }
   return guest, nil
}

func (g Guest) Source(space *AudioSpace) (*Source, error) {
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

const bearer =
   "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=" +
   "1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

type Guest struct {
   Guest_Token string
}

type AudioSpace struct {
   Metadata struct {
      Media_Key string
      Title string
      State string
      Started_At int64
      Ended_At int64 `json:"ended_at,string"`
   }
   Participants struct {
      Admins []struct {
         Display_Name string
      }
   }
}

// https://twitter.com/i/spaces/1jMJgednpreKL?s=20
func SpaceID(addr string) (string, error) {
   parse, err := url.Parse(addr)
   if err != nil {
      return "", err
   }
   return path.Base(parse.Path), nil
}

const spacePersistedQuery = "lFpix9BgFDhAMjn9CrW6jQ"

func (a AudioSpace) Time() time.Time {
   return time.UnixMilli(a.Metadata.Started_At)
}

func (g Guest) AudioSpace(id string) (*AudioSpace, error) {
   var str strings.Builder
   str.WriteString("https://twitter.com/i/api/graphql/")
   str.WriteString(spacePersistedQuery)
   str.WriteString("/AudioSpaceById")
   req, err := http.NewRequest("GET", str.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + bearer},
      "X-Guest-Token": {g.Guest_Token},
   }
   buf, err := json.Marshal(spaceRequest{ID: id})
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "variables=" + url.QueryEscape(string(buf))
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var space struct {
      Data struct {
         AudioSpace AudioSpace
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&space); err != nil {
      return nil, err
   }
   return &space.Data.AudioSpace, nil
}

type Source struct {
   Location string // Segment
}

type spaceRequest struct {
   ID string `json:"id"`
   IsMetatagsQuery bool `json:"isMetatagsQuery"`
   WithBirdwatchPivots bool `json:"withBirdwatchPivots"`
   WithDownvotePerspective bool `json:"withDownvotePerspective"`
   WithReactionsMetadata bool `json:"withReactionsMetadata"`
   WithReactionsPerspective bool `json:"withReactionsPerspective"`
   WithReplays bool `json:"withReplays"`
   WithScheduledSpaces bool `json:"withScheduledSpaces"`
   WithSuperFollowsTweetFields bool `json:"withSuperFollowsTweetFields"`
   WithSuperFollowsUserFields bool `json:"withSuperFollowsUserFields"`
}
