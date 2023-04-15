# SoundCloud

## Android client

https://play.google.com/store/apps/details?id=com.soundcloud.android

Must use Android API 24 or higher:

~~~
sdkVersion:'24'
~~~

which means API 23 will fail:

~~~
APK Installer
---------------------------
The APK failed to install.<br/> Error: INSTALL_FAILED_OLDER_SDK
~~~

So you must install system certificate.

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
