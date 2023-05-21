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

links:

~~~
developer.twitter.com/docs/authentication/api-reference/token
developer.twitter.com/docs/authentication/oauth-2-0
developer.twitter.com/docs/authentication/oauth-2-0/application-only

developer.twitter.com/docs/authentication/oauth-1-0a/authorizing-a-request
developer.twitter.com/docs/authentication/oauth-1-0a/creating-a-signature
~~~
