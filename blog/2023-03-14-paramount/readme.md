# Paramount

This is what I have been using:

~~~
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/
tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD?assetTypes=DASH_CENC&formats=MPEG-DASH
~~~

which returns up to 1080. Whats interesting is, this returns up to 2160:

~~~
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/
tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD?formats=MPEG-DASH
~~~

responses:

~~~
302 Found [assetTypes=DASH_CENC formats=MPEG-DASH]
302 Found [formats=MPEG-DASH]
412 Precondition Failed []
412 Precondition Failed [assetTypes=DASH_CENC]
~~~

## Paramount Android client

https://play.google.com/store/apps/details?id=com.cbs.app

Note searching the APK for `theplatform` fails. So the client is requesting
this URL:

~~~
https://pubads.g.doubleclick.net/ondemand/dash/content/2497752/vid/
7H97Abj_yordO5mKdGAyWOrnM87UCpE9/TUL/streams/
cb0d5705-ff9b-4640-bd76-e05dbd2225c3/manifest.mpd
~~~

which splits the video into multiple Periods to allow for ads. Also we have a
request that returns single video in some cases:

~~~
https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/
7H97Abj_yordO5mKdGAyWOrnM87UCpE9.json

https://vod-gcs-cedexis.cbsaavideo.com/intl_vms/2023/02/14/2172339267761/
1965235_cenc_fmp4_hdr_dash/stream.mpd
~~~

and split video in others:

~~~
https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/
rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW.json

https://vod-gcs-cedexis.cbsaavideo.com/intl_vms/2017/09/23/1053238851585/dubs/
1664335_cenc_precon_dash/stream.mpd
~~~

## Paramount Android client 2021

With version 8.0.46 (com.cbs.app), I get this message:

> Update Needed to Continue

I could block the request that checks the version, but I could not figure out
which one that is. I searched for the `versionCode` and `versionString`. Same
result with version 8.0.00. Similar error with 7.2.5 (CBS - Full Episodes &
Live TV):

> We are currently experiencing some technical difficulties

CBS app (com.cbs.tve) from 2021, uses the same 8.0.46 `versionString`, so I
assume the result would be the same as Paramount.
## CBS web client

<https://cbs.com/shows/video/_3vZxZoTZV2s2ig346INss95UDzokC20>

so the client is requesting this URL:

~~~
https://pubads.g.doubleclick.net/ondemand/dash/content/2497752/vid/
_3vZxZoTZV2s2ig346INss95UDzokC20/TUL/streams/
f48bcd8c-2c2a-4094-85af-a2cdc3e64bae/manifest.mpd
~~~

which splits the video into multiple Periods to allow for ads. However, in the
HTML response of the original URL, we also have this:

~~~
https://vod-gcs-cedexis.cbsaavideo.com/intl_vms/2023/03/07/2179633731600/
2006101_cenc_dash/stream.mpd
~~~

which is the unsplit video.

## Paramount web client

https://paramountplus.com/shows/video/rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW

so the client is requesting this URL:

~~~
https://pubads.g.doubleclick.net/ondemand/dash/content/2497752/vid/
rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW/TUL/streams/
dedd6f4c-a4f3-4aa6-8c65-192d525ab127/manifest.mpd
~~~

which splits the video into multiple Periods to allow for ads. However, in the
HTML response of the original URL, we also have this:

~~~
https://vod-gcs-cedexis.cbsaavideo.com/intl_vms/2017/09/23/1053238851585/dubs/
1664335_cenc_dash/stream.mpd
~~~

which is the unsplit video.

## CBS Android client

https://play.google.com/store/apps/details?id=com.cbs.tve

so the client is requesting this URL:

~~~
https://pubads.g.doubleclick.net/ondemand/dash/content/2497752/vid/
_3vZxZoTZV2s2ig346INss95UDzokC20/TUL/streams/
dc860dc5-1cf1-4a52-91cb-13654e097007/manifest.mpd
~~~

which splits the video into multiple Periods to allow for ads. Also in the
response to this request:

~~~
https://cbsdigital.cbs.com/apps-api/v2.0/androidphone/video/cid/
_3vZxZoTZV2s2ig346INss95UDzokC20.json
~~~

we have another split:

~~~
https://vod-gcs-cedexis.cbsaavideo.com/intl_vms/2023/03/07/2179633731600/
2006101_cenc_precon_dash/stream.mpd
~~~

