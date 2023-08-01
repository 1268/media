# web

- https://github.com/ytdl-org/youtube-dl/issues/32470
- https://max.com/a/video/the-white-lotus-s1-e1

this is it:

~~~
GET https://clips-media-aka.warnermediacdn.com/beam/clips/2023-05/1187541-6160a70bb990464da92734e0597b23e6/cmaf/master_cl_de.m3u8 HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: identity
origin: https://www.max.com
referer: https://www.max.com/
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: cross-site
te: trailers
content-length: 0
~~~

from:

~~~
GET https://medium.ngtv.io/v2/media/meb52c30a61e34b63a0ca66946f1b515e5aaad5f9d/desktop?appId=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuZXR3b3JrIjoiaGJvbWF4IiwicHJvZHVjdCI6ImJlYW0iLCJwbGF0Zm9ybSI6IndlYi10b3AyIiwiYXBwSWQiOiJoYm9tYXgtYmVhbS13ZWItdG9wMi1wMnFiMnAifQ.8XAm_qxXSa7REssRiJsrAnO0eh_Ljs24muIfZyaLE-8 HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0
accept: */*
accept-language: en-US,en;q=0.5
accept-encoding: identity
referer: https://www.max.com/
origin: https://www.max.com
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: cross-site
te: trailers
content-length: 0
~~~

from:

~~~
GET https://www.max.com/a/video/the-white-lotus-s1-e1 HTTP/2.0
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0
accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
accept-language: en-US,en;q=0.5
accept-encoding: identity
upgrade-insecure-requests: 1
sec-fetch-dest: document
sec-fetch-mode: navigate
sec-fetch-site: none
sec-fetch-user: ?1
te: trailers
content-length: 0
~~~
