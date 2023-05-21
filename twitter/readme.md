# Twitter

https://github.com/virezox/mech/tree/6c5d82a8

OAuth:

https://github.com/virezox/mech/tree/ef55d038

## Android client

- https://github.com/httptoolkit/frida-android-unpinning
- https://play.google.com/store/apps/details?id=com.twitter.android

download:

~~~
googleplay -d com.twitter.android
~~~

install:

~~~
adb install-multiple (Get-ChildItem *.apk)
~~~

## How to get Bearer Token?

Make a request to this page, without Cookies:

https://twitter.com

in the response should be something like this:

~~~
"authorization":"Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZ...
~~~
