# May 2023

https://amcplus.com/create

can web client sign up without these:

~~~
google.com
gstatic.com
~~~

If yes, can we automate the web sign up? If no, can we automate the Android app
sign up? Also, can we see if an error is due to lapsed account, or missing
media? Need to capture responses:

- active account, active media
- active account, expired media
- expired account, active media
- expired account, expired media

or, we could just use an email like this:

~~~
2023-05-15@mailsac.com
~~~

this fails if logged out:

https://amcplus.com/movies/nocebo--1061554

~~~
GET https://gw.cds.amcn.com/content-compiler-cr/api/v1/content/amcn/amcplus/path/movies/nocebo--1061554? HTTP/2.0

HTTP/2.0 404 
date: Mon, 15 May 2023 22:54:30 GMT
content-type: application/json; charset=utf-8
x-amcn-cache: MISS
x-krakend: Version 2.1.0-ee
x-krakend-completed: false
x-powered-by: Express
x-llid: ab3cd1cdbdb7b80adde1a535a486570a
content-length: 131
access-control-allow-origin: https://www.amcplus.com
access-control-expose-headers: Content-Length,Content-Type,X-Amcn-Bc-Jwt

{
  "error": "Manifest not found for [path]:(/movies/nocebo--1061554) with error: This ID returned a 204",
  "status": 404,
  "success": false
}
~~~

this fails if logged out:

https://amcplus.com/shows/orphan-black/episodes--1010634

~~~
GET https://gw.cds.amcn.com/content-compiler-cr/api/v1/content/amcn/amcplus/path/shows/orphan-black/episodes--1010634? HTTP/2.0

HTTP/2.0 404 
date: Mon, 15 May 2023 22:56:50 GMT
content-type: application/json; charset=utf-8
x-amcn-cache: MISS
x-krakend: Version 2.1.0-ee
x-krakend-completed: false
x-powered-by: Express
x-llid: be5fc7f945fb2742d1a2dc08394d8bc1
content-length: 145
access-control-allow-origin: https://www.amcplus.com
access-control-expose-headers: Content-Length,Content-Type,X-Amcn-Bc-Jwt

{
  "error": "Manifest not found for [path]:(/shows/orphan-black/episodes--1010634) with error: This ID returned a 204",
  "status": 404,
  "success": false
}
~~~

this always fails:

- https://amcplus.com/movies/stop-making-sense--1059031
- https://reddit.com/r/MovieOfTheDay/comments/yqv7bu

~~~
GET https://gw.cds.amcn.com/content-compiler-cr/api/v1/content/amcn/amcplus/path/movies/stop-making-sense--1059031? HTTP/2.0

HTTP/2.0 404 
date: Mon, 15 May 2023 22:58:07 GMT
content-type: application/json; charset=utf-8
x-amcn-cache: MISS
x-krakend: Version 2.1.0-ee
x-krakend-completed: false
x-powered-by: Express
x-llid: 84345429e18a4c77098d77b1e8a931dd
content-length: 142
access-control-allow-origin: https://www.amcplus.com
access-control-expose-headers: Content-Length,Content-Type,X-Amcn-Bc-Jwt

{
  "error": "Manifest not found for [path]:(/movies/stop-making-sense--1059031) with error: This ID returned a 204",
  "status": 404,
  "success": false
}
~~~
