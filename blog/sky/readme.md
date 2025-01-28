# show.sky.ch

~~~
url = https://show.sky.ch/de/filme/2035/a-knights-tale
monetization = FLATRATE
country = Switzerland
~~~

- https://github.com/sunsettrack4/plugin.video.skych/blob/master/addon.py
- https://justwatch.com/ch/Anbieter/sky

## web

Amazon CAPTCHA is required

## tv google play

if you request TV app, phone app is returned:

~~~
> play -i homedia.sky.sport -leanback
details[8] = 0 USD
details[13][1][4] = 1.18.1.142
details[13][1][16] = Jan 21, 2025
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 8.0 and up
details[15][18] = https://support.sky.ch/hc/en-us/articles/9520105066140
downloads = 468.36 thousand
name = Sky
size = 35.09 megabyte
version code = 584

> play -i homedia.sky.sport
details[8] = 0 USD
details[13][1][4] = 1.18.1.142
details[13][1][16] = Jan 21, 2025
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 8.0 and up
details[15][18] = https://support.sky.ch/hc/en-us/articles/9520105066140
downloads = 468.36 thousand
name = Sky
size = 35.09 megabyte
version code = 584
~~~

## tv apk mirror

https://apkmirror.com/apk/sky-switzerland/sky-android-tv

if you try version 2.3.6.6 with just residential proxy, all the request fail,
which means version is too old.

## tv other locations

cannot find a newer TV APK at other locations, which means it was likely dropped
in favor of phone app

## phone apk mirror

too old

## phone apk fab

login is protected:

~~~go
req.Header["Cookie"] = []string{
   "aws-waf-token=2e86b681-4c6d-40cd-9856-9ec0780664e5:HAoAkAsSO8kGAAAA:wWotxIx/qIxwEPx20cZJqorgSm4bt5YuAhntIxvP7HAXyKYgrnJD39XjU8Vlcwcb88umfrKppm+luczkW5DnyMk7l+eU7KbxOIi76foo8gRgpdS9e18/BwJVciM=",
}
~~~

if you drop the Amazon request or the Cookie, the login fails

https://apkfab.com/it/sky/homedia.sky.sport
