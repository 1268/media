# Paramount

`DASH_CENC`:

~~~
pass "formats=MPEG-DASH"
pass "formats=MPEG-DASH&mbr=true"
pass "format=smil&formats=MPEG-DASH"
pass "assetTypes=DASH_CENC&formats=MPEG-DASH"
pass "format=smil&formats=MPEG-DASH&mbr=true"
pass "assetTypes=DASH_CENC&formats=MPEG-DASH&mbr=true"
pass "assetTypes=DASH_CENC&format=smil&formats=MPEG-DASH"
pass "assetTypes=DASH_CENC&format=smil&formats=MPEG-DASH&mbr=true"
~~~

Downloadable:

~~~
pass ""
pass "mbr=true"
pass "format=smil"
pass "formats=MPEG-DASH"
pass "format=smil&mbr=true"
pass "assetTypes=Downloadable"
pass "formats=MPEG-DASH&mbr=true"
pass "format=smil&formats=MPEG-DASH"
pass "assetTypes=Downloadable&mbr=true"
pass "assetTypes=Downloadable&format=smil"
pass "format=smil&formats=MPEG-DASH&mbr=true"
pass "assetTypes=Downloadable&format=smil&mbr=true"
~~~

StreamPack:

~~~
pass ""
pass "mbr=true"
pass "format=smil"
pass "formats=MPEG-DASH"
pass "format=smil&mbr=true"
pass "assetTypes=StreamPack"
pass "formats=MPEG-DASH&mbr=true"
pass "format=smil&formats=MPEG-DASH"
pass "assetTypes=StreamPack&mbr=true"
pass "assetTypes=StreamPack&format=smil"
pass "format=smil&formats=MPEG-DASH&mbr=true"
pass "assetTypes=StreamPack&format=smil&mbr=true"
~~~

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

URL is here:

~~~
com.cbs.app_12.0.28> rg 2198311517
String uri = new Uri.Builder().scheme(ProxyConfig.MATCH_HTTP).
authority("link.theplatform.com").
appendPath(Constants.APPBOY_PUSH_SUMMARY_TEXT_KEY).appendPath("dJ5BDC").
appendPath(CommonUtil.Directory.MEDIA_ROOT).
appendPath(DistributedTracing.NR_GUID_ATTRIBUTE).appendPath("2198311517").
appendPath(contentId).appendQueryParameter("assetTypes", "DASH_CENC").
appendQueryParameter("formats", "MPEG-DASH").
appendQueryParameter("format", "smil").build().toString();
~~~

this is interesting:

http://can.cbs.com/thunder/player/videoPlayerService.php?partner=cbs&contentId=YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_

result:

~~~xml
<assetType>Downloadable</assetType>
<assetType>PDL_MP4</assetType>
~~~

search

~~~
theplatform assetType "Downloadable"
~~~

