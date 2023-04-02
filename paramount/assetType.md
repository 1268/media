# assetType

We know that `Downloadable` is a valid value for `assetTypes`:

~~~
> curl -i link.theplatform.com/s/dJ5BDC/media/guid/2198311517/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_?assetTypes=Downloadable
HTTP/1.1 302 Found
Location: https://candownloads.cbsaavideo.com/vr/cbsnews/2023/01/05/2158351939508/0115_60minutes_full_1_1627388_5192.mp4?hdnea=st=1680441369~exp=1680700599~acl=/vr/cbsnews/2023/01/05/2158351939508/0115_60minutes_full_1_1627388_5192.mp4*~hmac=e2b1f173121395635d0e236efee7d78cfbf1d2f41c1264080250f4b42a30dd8b
~~~

because if you provide an invalid value, the request fails:

~~~
> curl -i link.theplatform.com/s/dJ5BDC/media/guid/2198311517/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_?assetTypes=invalid
HTTP/1.1 412 Precondition Failed
~~~

the question is, how do we enumerate the available `assetTypes` for a given
video? Previously we could use a URL like this:

<http://can.cbs.com/thunder/player/videoPlayerService.php?partner=cbs&contentId=YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_>

but it now returns an empty response. If we visit this page logged in:

<https://paramountplus.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_>

in the response we can see this:

~~~
"playerLocUrl":"http:\/\/can.cbs.com\/thunder\/player\/chrome\/canplayer.swf?allowScriptAccess=always&allowFullScreen=true&pid=4D_TMl1QG_vO&partner=pp_us_lcp_desktop&autoPlayVid=false&owner=CBS Production News"
~~~

we can download this URL:

~~~
curl -o canplayer.swf 'can.cbs.com/thunder/player/chrome/canplayer.swf?allowScriptAccess=always&allowFullScreen=true&pid=4D_TMl1QG_vO&partner=pp_us_lcp_desktop&autoPlayVid=false&owner=CBS%20Production%20News'
~~~

Then extract using 7-Zip or similar:

https://github.com/mcmilk/7-Zip-zstd

we get another file `canplayer~.swf`. We can then extract that file. If we check
the largest file, in this case named `15.82`, we find some interesting data:

~~~
http://can.cbs.com/thunder/player/videoPlayerService.php?pid=$PID$&domain=$DOMAIN$&partner=$PARTNER$
http://can.cbs.com/thunder/player/videoPlayerService.php?pid=$PID$&domain=$DOMAIN$&partner=$PARTNER$&auth=$AUTH$
~~~

From the URL above, we have the `pid` and `partner`, so we can try that:

<http://can.cbs.com/thunder/player/videoPlayerService.php?pid=4D_TMl1QG_vO&partner=pp_us_lcp_desktop>

but again empty response. So we need to also add the `domain`, or `auth` or both.
How do we get the `domain`? Searching the rest of the SWF returns no match. If
we visit this page logged in:

<https://paramountplus.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_>

we should see a request like this:

https://player-services.paramountplus.com/1.14.0/smart-tag/smart.tag.js

in the response is this:

~~~
cid:"https://can-services.cbs.com/canServices/playerService/video/search.xml?contentId=#CONTENT_ID#&domain=#DOMAIN#&partner=#PARTNER#&showEncodes=true&st=1",
pid:"https://can-services.cbs.com/canServices/playerService/video/search.xml?pid=#PID#&domain=#DOMAIN#&partner=#PARTNER#"
~~~

we can try both URL without `domain`:

- <https://can-services.cbs.com/canServices/playerService/video/search.xml?contentId=YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_&partner=pp_us_lcp_desktop&showEncodes=true&st=1>
- <https://can-services.cbs.com/canServices/playerService/video/search.xml?pid=4D_TMl1QG_vO&partner=pp_us_lcp_desktop>

but still empty response. If we check the JavaScript Debugger:

~~~
player-services.paramountplus.com
   1.14.0
      smart-tag
         smart.tag.js
~~~

right click the tab, then click Pretty print source. Click any line with `domain`
to set Breakpoints, then refresh the page. The Debugger did not pause, so those
lines must have not been called. However I did find this in the same JavaScript:

~~~
domain: {
   localhost: {
      token: 'localhost',
      value: 'localhost'
   },
   automation: {
      token: 'automation-player-services.cbs.com',
      value: 'automation-player-services.cbs.com'
   },
   dev: {
      token: 'branch-player-services.cbs.com',
      value: 'branch-player-services.cbs.com'
   },
   branch: {
      token: 'branch-player-services.cbs.com',
      value: 'branch-player-services.cbs.com'
   },
   stage: {
      token: 'stage-player-services.cbs.com',
      value: 'stage-player-services.cbs.com'
   },
   perf: {
      token: 'perf-player-services.cbs.com',
      value: 'perf-player-services.cbs.com'
   },
   'stage-hub': {
      token: 'stage-hub-player-services.cbs.com',
      value: 'stage-hub-player-services.cbs.com'
   },
   preview: {
      token: 'preview-player-services.cbs.com',
      value: 'preview-player-services.cbs.com'
   },
   prod: {
      token: 'player-services.cbs.com',
      value: 'player-services.cbs.com'
   }
},
~~~

so we can try again with the `domain=prod`:

- <https://can-services.cbs.com/canServices/playerService/video/search.xml?contentId=YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_&partner=pp_us_lcp_desktop&showEncodes=true&st=1&domain=prod>
- <https://can-services.cbs.com/canServices/playerService/video/search.xml?pid=4D_TMl1QG_vO&partner=pp_us_lcp_desktop&domain=prod>

and `domain=player-services.cbs.com`:

- <https://can-services.cbs.com/canServices/playerService/video/search.xml?contentId=YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_&partner=pp_us_lcp_desktop&showEncodes=true&st=1&domain=player-services.cbs.com>
- <https://can-services.cbs.com/canServices/playerService/video/search.xml?pid=4D_TMl1QG_vO&partner=pp_us_lcp_desktop&domain=player-services.cbs.com>

but still empty response. If we search the logged in HTML response of the
public URL:

<https://paramountplus.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_>

we find some `assetType` values:

~~~
"assetType":"DASH_CENC"
"assetType":"ThumbSheet"
"assetType":"Thumbnail"
~~~

but not all, since `Downloadable` is missing. We can download the APK and
extract:

~~~
googleplay -d com.cbs.app -v 211204029
jadx com.cbs.app-211204029.apk
~~~

searching `playerservice` returns no results. searching `assettype` returns
nothing interesting. We can then start the APK:

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://www.paramountplus.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_/
~~~

searching `[pP]layerService` returns no results. Searching `assetType` returns
this request:

https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_.json?locale=en-us&at=ABDXaI84LZDFGjCtn56%2F6e72mswXRbGsgOoBYKTRIIBnpzVXj4y%2BrbihkUg%2BTw7RO2k%3D

but again, not all `assetType` values are returned:

~~~
"assetType": "DASH_CENC_PRECON",
"assetType": "Thumbnail",
"assetType": "ThumbSheet",
~~~

So it seems no method is currently available to enumerate all available
`assetType` values, with some being hidden.
