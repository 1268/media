# Paramount+

https://justwatch.com/us/provider/paramount-plus

## Android client

~~~
com.cbs.app
~~~

Create device with Android API 23 or higher. Install user certificate. After the
commercial you might get an error, try again.

## How to get aid?

In the response to the same request, you should see something like this:

~~~xml
<param name="trackingData" value="aid=2198311517|b=1000|bc=CBSI-NEW|ci=1|cid=1...
~~~

## How to get app\_secret?

~~~
com\cbs\app\dagger\DataLayerModule.java
dataSourceConfiguration.setCbsAppSecret("6c70b33080758409");
~~~

https://github.com/matthuisman/slyguy.addons/issues/136

## How to get secret\_key?

~~~
com\cbs\app\androiddata\retrofit\util\RetrofitUtil.java
SecretKeySpec secretKeySpec = new SecretKeySpec(b("302a6a0d70a7e9b967f91d39fef3e387816e3095925ae4537bce96063311f9c5"), "AES");
~~~

https://github.com/matthuisman/slyguy.addons/issues/136

## How to get sid?

https://play.google.com/store/apps/details?id=com.cbs.app

Install user certificate. Start video, and you should see a request like this:

~~~
GET /s/dJ5BDC/fNsRH_fjko5T?format=SMIL&Tracking=true&sig=006229620e7f3db019fc0...
Host: link.theplatform.com
X-NewRelic-ID: VQ4FVlJUARABVVRXAwEOVFc=
User-Agent: Dalvik/2.1.0 (Linux; U; Android 7.0; Android SDK built for x86 Bui...
Connection: Keep-Alive
Accept-Encoding: gzip
content-length: 0
~~~

## Why DASH and HLS?

We have FairPlay HLS:

~~~
https://vod-gcs-cedexis.cbsaavideo.com/intl_vms/2020/05/07/1735196227871/2367_fp_hls/master.m3u8
~~~

and StreamPack HLS:

~~~
https://cbsios-vh.akamaihd.net/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2020/05/07/1735196227871/0_0_3436402_ful01_2588_503000.mp4.csmil/master.m3u8
~~~

and CENC DASH:

~~~
https://vod-gcs-cedexis.cbsaavideo.com/intl_vms/2020/05/07/1735196227871/2367_cenc_dash/stream.mpd
~~~

but no StreamPack DASH.
