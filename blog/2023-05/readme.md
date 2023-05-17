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

for active account, you always get:

~~~
.data.properties.entitlements[1] = "ob-sub-amcplus"
~~~

using this:

https://amcplus.com/movies/nocebo--1061554

active account, path/watch:

~~~
active-account-path-watch.txt
~~~

active account, path:

~~~
active-account-path.txt
~~~

expired account, path:

~~~
expired-account-path.txt
~~~

expired account, path/watch
