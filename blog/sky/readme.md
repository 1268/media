# show.sky.ch

~~~
url = https://show.sky.ch/de/filme/2035/a-knights-tale
monetization = FLATRATE
country = Switzerland
~~~

- https://github.com/sunsettrack4/plugin.video.skych/blob/master/addon.py
- https://justwatch.com/ch/Anbieter/sky
- https://proxy-seller.com

~~~
GET /b0c7c16365cdaef75e81/0/0/m_drm_widevine.mpd?z32=MF2WI2LPL5RW6ZDFMNZT2YLBMMTGG43JMQ6TCOBRIU2UCOKEIE2UIMJUIFBUGLJSGQ3TMMRRG4ZDSMBWII3TKNZREZSHE3J5MV4HA2LSMF2GS33OHIYTOMZYGUZTAOJTHETGS3TJORUWC3DSMF2GKPJSGAYDAJTNMF4HEYLUMU6TQMBQGATG22LOOJQXIZJ5GUYDAJTQOJSWMZLSOJSWIX3MMFXGO5LBM5ST2ZDFEZZWSZZ5GI3F6YZYGQ2DQY3DG4YWGNZWGAYDIZLEMMYDKOLCMVRGENJQMJSTANJVEZZXKYTUNF2GYZLTHVUGSZDEMVXC243ENATHK43FOJPWSZB5ONVXSX3DNA5DUMJZGM4DINBREZ3D2MA HTTP/1.1
Host: sunlau1-1-dashenc-vod.zahs.tv
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:128.0) Gecko/20100101 Firefox/128.0
Accept: */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate, br, zstd
Referer: https://show.sky.ch/
Origin: https://show.sky.ch
Connection: keep-alive
Sec-Fetch-Dest: empty
Sec-Fetch-Mode: cors
Sec-Fetch-Site: cross-site
content-length: 0
~~~

## tv

- https://play.google.com/store/apps/details?id=homedia.sky.sport
- https://apkmirror.com/apk/sky-switzerland/sky-android-tv

lets try it with only location proxy for now.

it seems 2.3.6.6 (514) (june 13 2024) is too old. can we find a newer version?
no

## phone

I found the Sky Switzerland TV app, but its from last year. when I intercepted
it all the request failed even with residential proxy - so likely its just an
old version of the the app, and I cant find a newer version online - its
possible they are just using the phone app for both. also I cant download it
directly from Google Play, even with residential proxy, since they tie your
location to your account location instead of IP - you can change it once a year
but its not worth it for this. I guess I could create a Google account from a
Switzerland IP, but making a TV call to Google Play returned the phone app so
its probably a waste of time. so looks like I am stuck with the web client. I
saw someone solving the Amazon CAPTCHA using https://huggingface.co/ - but I
don't wanna fuck with that 

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
