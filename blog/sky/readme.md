# show.sky.ch

~~~
url = https://show.sky.ch/de/filme/2035/a-knights-tale
monetization = FLATRATE
country = Switzerland
~~~

- https://github.com/sunsettrack4/plugin.video.skych/blob/master/addon.py
- https://justwatch.com/ch/Anbieter/sky

## webDriver

no support for `https_proxy`:

<https://bugzilla.mozilla.org/show_bug.cgi?id=1944213>

no proxy authentication:

https://github.com/mozilla/geckodriver/issues/1872

no proxy authentication:

<https://bugzilla.mozilla.org/show_bug.cgi?id=1395886>

we could use a proxy ladder, but does not seem worth the trouble

## phone

login is protected:

~~~go
req.Header["Cookie"] = []string{
   "aws-waf-token=2e86b681-4c6d-40cd-9856-9ec0780664e5:HAoAkAsSO8kGAAAA:wWotxIx/qIxwEPx20cZJqorgSm4bt5YuAhntIxvP7HAXyKYgrnJD39XjU8Vlcwcb88umfrKppm+luczkW5DnyMk7l+eU7KbxOIi76foo8gRgpdS9e18/BwJVciM=",
}
~~~

if you drop the Amazon request or the Cookie, the login fails

https://apkfab.com/it/sky/homedia.sky.sport

## tv

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

## web

Amazon CAPTCHA is required
