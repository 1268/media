# YouTube

> I’m wide awake\
> I’m wide awake\
> Wide awake\
> I’m not sleeping
>
> [U2 - Bad (1984)](//youtube.com/watch?v=0s30qw4XV_E)

## Android client

Must use Android API 31 or higher.

https://android.stackexchange.com/a/245551

## Device OAuth

- https://datatracker.ietf.org/doc/html/rfc8628
- https://developers.google.com/identity/sign-in/devices

## How to get X-Goog-API-Key

Make a request like this:

~~~
GET / HTTP/2
Host: www.youtube.com
~~~

in the response you should see something like this:

~~~
"INNERTUBE_API_KEY":"AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
~~~

https://cloud.google.com/apis/docs/system-parameters

## How to get `client_id` and `client_secret`

Set User-Agent to [1]:

~~~
Mozilla/5.0 (ChromiumStylePlatform) Cobalt/Version
~~~

Then visit:

https://www.youtube.com/tv

Next open your browser menu, and click Web Developer, Network or similar. Then
go back to the page, and click "Sign in", then "Sign in with a web browser". On
the next page, dont bother with any of the instructions, just go back to
Developer Tools, and after about five seconds you should see a JSON request like
this:

~~~
POST /o/oauth2/token HTTP/1.1
Host: www.youtube.com

{"client_id":"861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com",
"client_secret":"SboVhoG9s0rNafixCSGGKXAT",
"code":"AH-1Ng14qVvEj76OeM_h14Mgklgyhchbyc67MhULhCKPY6K-0DTYJqaKng2ULVFTmTzU...",
"grant_type":"http://oauth.net/grant_type/device/1.0"}
~~~

1. <https://github.com/youtube/cobalt/blob/master/src/cobalt/browser/user_agent_string.cc>

## Image

Is `maxres1` always available? No:

- <http://i.ytimg.com/vi_webp/hq2KgzKETBw/maxres1.webp>
- http://i.ytimg.com/vi/hq2KgzKETBw/maxres1.jpg

Is `sd1` always available? No:

- <http://i.ytimg.com/vi_webp/hq2KgzKETBw/sd1.webp>
- http://i.ytimg.com/vi/hq2KgzKETBw/sd1.jpg

If `hq1` always available? Yes:

http://i.ytimg.com/vi/hq2KgzKETBw/hq1.jpg

## MWEB client

~~~
curl -A iPad https://m.youtube.com
~~~
