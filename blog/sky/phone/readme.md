# show.sky.ch

~~~
GET https://clientapi.prd.sky.ch/stream/2035/MOVIE HTTP/2.0
appversion: 1.18.1.125
devicecode: ANDROID_INAPP
devicename: Google AOSP on IA Emulator
osversion: 28
deviceplatform: ANDROID_INAPP
accept-encoding: compress
api-version: 1.13
accept-language: de
macaddress: 31f08176-dd2f-43a3-ae94-9e24190600fe
environment: SkyShow
serialnumber: 31f08176-dd2f-43a3-ae94-9e24190600fe
bundles: SHOW_PREMIUM
user-agent: okhttp/5.0.0-alpha.10
content-length: 0
authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxOTM4NDQxIiwiUHJmSWQiOiIyMTE5ODA5IiwiUm9sZXMiOiIiLCJCdW5kbGVzIjoie1wiU2t5U2hvd1wiOlwiU0hPV19QUkVNSVVNXCIsXCJTa3lTcG9ydFwiOlwiXCJ9IiwiZXhwIjoxNzM3OTQ1NjU5LCJpc3MiOiJodHRwczovL3d3dy5za3kuY2giLCJhdWQiOiJTa3kgVXNlcnMifQ.7MMfzGZw4HIqpNOloFsVi4LSNSWYNjTJRT-mhyIG3y4
~~~

## phone

login is protected:

~~~go
req.Header["Cookie"] = []string{
   "aws-waf-token=2e86b681-4c6d-40cd-9856-9ec0780664e5:HAoAkAsSO8kGAAAA:wWotxIx/qIxwEPx20cZJqorgSm4bt5YuAhntIxvP7HAXyKYgrnJD39XjU8Vlcwcb88umfrKppm+luczkW5DnyMk7l+eU7KbxOIi76foo8gRgpdS9e18/BwJVciM=",
}
~~~

if you drop the Amazon request or the Cookie, the login fails

https://apkfab.com/it/sky/homedia.sky.sport

~~~
config.armeabi_v7a.apk
~~~

so need Android 9
