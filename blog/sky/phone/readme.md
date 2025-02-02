## phone

login is protected:

~~~go
cookie: aws-waf-token=2e86b681-4c6d-40cd-9856-9ec0780664e5:HAoAkAsSO8kGAAAA:wW...
~~~

if you drop the Amazon request or the Cookie, the login fails

https://apkfab.com/it/sky/homedia.sky.sport

~~~
config.armeabi_v7a.apk
~~~

so need Android 9. this request:

~~~
GET https://clientapi.prd.sky.ch/stream/2035/MOVIE HTTP/2.0
devicecode: ANDROID_INAPP
authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxOTM4NDQ...
~~~

has no geo block, but the authorization only lasts five minutes, and the
refresh call is geo blocked, so the web client is better. also request above
does not accept these:

~~~
_ASP.NET_SessionId_ fail
SkyCheeseCake fail
sky-auth-token fail
~~~
