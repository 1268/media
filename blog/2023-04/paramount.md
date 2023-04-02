# April 2023

OK from my comment here:

https://github.com/ytdl-org/youtube-dl/issues/31864#issuecomment-1474345789

I have this command:

~~~
curl -v link.theplatform.com/s/dJ5BDC/media/guid/2198311517/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_
~~~

which previously returned an HLS URL:

<https://cbsios-vh.akamaihd.net/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2020/05/07/1735196227871/0_0_3436402_ful01_2588_503000.mp4.csmil/master.m3u8>

but now does not:

~~~json
{
   "title": "No AssetType/ProtectionScheme/Format Matches",
   "description": "None of the available releases match the specified AssetType, ProtectionScheme, and/or Format preferences",
   "isException": true,
   "exception": "NoAssetTypeFormatMatches",
   "responseCode": "412"
}
~~~

and this is still a valid ID:

<https://paramountplus.com/movies/video/wQH9yE_y_Dt4ekDYm3yelhhY2KXvOra_>

change user agent:

~~~
general.useragent.override
~~~

to:

~~~
Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Safari/605.1.15
~~~

https://whatismybrowser.com/guides/the-latest-user-agent/safari

OK I think I made some progress. If you monitor the Android app, you see a
request like this:

~~~
GET https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_.json?at=ABB90oFocIWpreOn1w0QqC%2FEi1lZTogArlwKOuD7ZmbjVsg7oCYtHqHGM1UFjJU8dW0%3D HTTP/2.0
cookie: CBS_COM=N0EwMjY0MDVENTU3MzJCNzJBMEQzMkIyMDQ0MjQyQUU6MTcxMTkxODg2NTMyND...
~~~

In the response is this:

~~~
"playerLocUrl": "http://can.cbs.com/thunder/player/chrome/canplayer.swf?allowScriptAccess=always&allowFullScreen=true&pid=u4uKImT6OqQR&partner=pp_us_lcp_android&autoPlayVid=false&owner=CBS Production News",
~~~

that URL is actually invalid because of the spaces, here is correct:

<http://can.cbs.com/thunder/player/chrome/canplayer.swf?allowScriptAccess=always&allowFullScreen=true&pid=u4uKImT6OqQR&partner=pp_us_lcp_android&autoPlayVid=false&owner=CBS%20Production%20News>

If you extract that file, you get another file `canplayer~.swf`. If you extract
that file, the largest file should be named something like `15.82`. Inside that
file is some interesting stuff:

~~~
http://can.cbs.com/thunder/player/videoPlayerService.php?pid=$PID$&domain=$DOMAIN$&partner=$PARTNER$
http://can.cbs.com/thunder/player/videoPlayerService.php?pid=$PID$&domain=$DOMAIN$&partner=$PARTNER$&auth=$AUTH$
~~~

thats as far as I have gotten so far.
