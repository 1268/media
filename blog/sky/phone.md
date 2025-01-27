# phone





## apkfab.com

https://apkfab.com/it/sky/homedia.sky.sport

no x86, but does have armeabi-v7a

1.18.1.125

so we will need Android 9. install system certificate. first try with no proxy

~~~
adb install-multiple (Get-ChildItem *.apk)
adb shell input text EMAIL
~~~

1. email
2. password
3. continue

no connection to sky. if we try Mullvad it just times out. smart proxy works

~~~
mitmproxy --upstream-auth USERNAME:PASSWORD `
--mode upstream:http://ch.smartproxy.com:29001

mitmproxy --upstream-auth USERNAME:PASSWORD `
--mode upstream:http://res.proxy-seller.com:10000
~~~

## apkmirror.com

I cant download it with my tool. this:

https://apkmirror.com/apk/sky-switzerland/sky-3

is old:

Sky 1.18.1.82 June 13, 2024 CDT

## apkpure.com

https://apkpure.com/sky/homedia.sky.sport

~~~
config.arm64_v8a.apk
~~~

## google.com

~~~
> play -i homedia.sky.sport
details[8] = 0 USD
details[13][1][4] = 1.18.1.142
details[13][1][16] = Jan 21, 2025
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 8.0 and up
details[15][18] = https://support.sky.ch/hc/en-us/articles/9520105066140
downloads = 468.02 thousand
name = Sky
size = 35.09 megabyte
version code = 584
~~~

https://play.google.com/store/apps/details?id=homedia.sky.sport

## uptodown.com

https://homedia-sky-sport.en.uptodown.com

is newer:

1.18.1.109

~~~
config.arm64_v8a.apk
~~~

so we would need to switch to Android 11 device. can we find newer x86?
