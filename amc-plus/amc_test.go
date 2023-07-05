package amc

var tests = []struct {
   address string
   key string
   pssh string
} {
   // amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
   episode: {
      address: "/shows/orphan-black/episodes/season-1-instinct--1011152",
      key: "95f11e40064f47007e7d950bd52d7b95",
      pssh: "AAAAVnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADYIARIQJqlCz6NjSI2kDWew20wbGRoNd2lkZXZpbmVfdGVzdCIIMTIzNDU2NzgyB2RlZmF1bHQ=",
   },
   // amcplus.com/movies/nocebo--1061554
   movie: {address: "/movies/nocebo--1061554"},
}

const (
   episode = iota
   movie
)
