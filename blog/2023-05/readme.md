# May 2023

this request:

~~~
GET /2/search/adaptive.json?q=filter%3Aspaces&tweet_mode=extended HTTP/1.1
Host: api.twitter.com
Authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs=1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA
~~~

is returning this response:

~~~
HTTP/2.0 429 Too Many Requests
~~~
