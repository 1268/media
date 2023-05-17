# May 2023

https://amcplus.com/create

if you try to sign up using the web client without these:

~~~
google.com
gstatic.com
~~~

you get this:

> Please, confirm you are human

can we automate the Android app sign up instead? Also, can we see if an error
is due to lapsed account, or missing media?

this fails if logged out:

https://amcplus.com/movies/nocebo--1061554

this fails if logged out:

https://amcplus.com/shows/orphan-black/episodes--1010634

this always fails:

- https://amcplus.com/movies/stop-making-sense--1059031
- https://reddit.com/r/MovieOfTheDay/comments/yqv7bu

For AMC+ expired media, you always get:

~~~
HTTP/2.0 404 
~~~

For active media, you always get:

~~~
HTTP/2.0 200
~~~

using this:

https://amcplus.com/movies/nocebo--1061554

with active account and `path/watch`, you get:

~~~
.data.children[1].type = "video-player-ap"
~~~

with active account and `path`, you get:

~~~
.data.children[0].type = "image"
.data.children[1].type = "image"
.data.children[2].type = "container"
.data.children[3].type = "navigation-ap"
~~~

with expired account and `path`, you get:

~~~
.data.children[0].type = "image"
.data.children[1].type = "image"
.data.children[2].type = "container"
.data.children[3].type = "navigation-ap"
.data.children[4].type = "modal"
~~~

with expired account and `path/watch`, you get:

~~~
.data.type = "redirect"
~~~
