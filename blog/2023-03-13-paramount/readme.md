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

works with lowercase too:

~~~
link.theplatform.com/s/dJ5BDC/media/guid/2198311517/
tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD?formats=mpeg-dash
~~~

## CBS Android client

https://play.google.com/store/apps/details?id=com.cbs.tve

## CBS web client

- https://cbs.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_
- https://paramountplus.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_

## Paramount Android client

https://play.google.com/store/apps/details?id=com.cbs.app

Android client request this URL:

~~~
https://vod-gcs-cedexis.cbsaavideo.com/intl_vms/2023/01/25/2165477955514/
1977711_cenc_precon_dash/
MTV_THECHALLENGEWORLDCHAMPIONSHIP_101_HD_V1_en-US_1961718_3_aac_128/seg_13.m4s
~~~

which comes from:

~~~
https://pubads.g.doubleclick.net/ondemand/dash/content/2497752/vid/
SQGw5gWeMWBeZfvz7D78tJ9jmkMUbL4X/TUL/streams/
3ac204df-51c7-46d8-99e7-4960bb84c4a4/manifest.mpd
~~~

avoid this:

~~~
https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/
SQGw5gWeMWBeZfvz7D78tJ9jmkMUbL4X.json
~~~

because the resulting MPD is:

~~~
assetType	"DASH_CENC_PRECON"
~~~

which means the video is broken up into pieces. Maybe cbs.com videos are better,
or CBS app?

## Paramount web client

searching the web client for MPEG-DASH returns nothing however. Web client is
requesting this URL:

~~~
https://vod-gcs-cedexis.cbsaavideo.com/intl_vms/2017/09/23/1053238851585/dubs/
1664335_cenc_precon_dash/CBS_SEAL_TEAM_101_HD_R_1563332_4500/seg_3.m4s
~~~

which comes from here:

~~~
https://pubads.g.doubleclick.net/ondemand/dash/content/2497752/vid/
rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW/TUL/streams/
279e4e07-11a6-4ebd-800c-ee5720115139/manifest.mpd
~~~
