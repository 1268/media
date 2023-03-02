# SoundCloud

## Android client

~~~
com.soundcloud.android
~~~

Install system certificate.

## How to get `client_id`

First, make a request like this:

~~~
GET / HTTP/2
Host: m.soundcloud.com
~~~

In the HTML response, you should see something like this:

~~~
"clientId":"iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX"
~~~

The `client_id` seems to last at least a year:

https://github.com/rrosajp/soundcloud-archive/commit/c02809dc

## Image

artworks:

~~~
https://soundcloud.com/oembed?format=json&url=https://soundcloud.com/western_vinyl/jessica-risker-cut-my-hair
https://i1.sndcdn.com/artworks-000308141235-7ep8lo-t500x.jpg
~~~

avatars:

~~~
https://soundcloud.com/oembed?format=json&url=https://soundcloud.com/pdis_inpartmaint/harold-budd-perhaps-moss
https://i1.sndcdn.com/avatars-000274827119-0dxutu-t500x.jpg
~~~

## Why does this exist?

January 28 2022.

I use the site myself.

https://soundcloud.com/afterhour-sounds/premiere-ele-bisu-caradamom-coffee
