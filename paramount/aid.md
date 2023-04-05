# aid

If we take a video like this:

<https://paramountplus.com/shows/video/vLJLNTktnWmP_KzDtTm9X7ki0SRvpZ2w>

we can get the info like this:

<https://paramountplus.com/apps-api/v2.0/androidphone/video/cid/vLJLNTktnWmP_KzDtTm9X7ki0SRvpZ2w.json?at=ABDSJZgxTScBBBO4aTPohedBu%2BEYBA1PYEHSGCtyFepqrLX%2BARQN5S9eZ8VWQ2zwNk8%3D>

which includes everything we need:

~~~
"cmsAccountId": "dJ5BDC",
"pid": "i_iqDzG1odu1",
~~~

but other videos:

<https://paramountplus.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_>

are missing the `pid`:

<https://paramountplus.com/apps-api/v2.0/androidphone/video/cid/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_.json?at=ABDSJZgxTScBBBO4aTPohedBu%2BEYBA1PYEHSGCtyFepqrLX%2BARQN5S9eZ8VWQ2zwNk8%3D>

unless we request with cookie:

~~~
GET https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_.json?at=ABAFzPkMhOzTnRSPIU8SHvxC3JgwGZ8eTCYf3NSJx3wm6fhNd5vz%2FfLw3TlCcaKYC%2Bc%3D HTTP/1.1
cookie: CBS_COM=N0EwMjY0MDVENTU3MzJCNzJBMEQzMkIyMDQ0MjQyQUU6MTcxMjA5OTQwMTg2OT...
~~~

which we can get like this:

~~~
POST https://www.paramountplus.com/apps-api/v2.0/androidphone/auth/login.json?at=ABDFhCKlU... HTTP/1.1
content-type: application/x-www-form-urlencoded

j_username=EMAIL&j_password=PASSWORD
~~~

we can instead use the `aid`:

<http://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_>

without authentication. Its a little messy though, because the current web client
and Android client no longer use the `aid`, so who knows how long it will be
around. Its in the Android source like this:

~~~
sources\com\cbs\downloader\impl\a.java
String uri = new Uri.Builder().scheme(ProxyConfig.MATCH_HTTP).
authority("link.theplatform.com").
appendPath(Constants.APPBOY_PUSH_SUMMARY_TEXT_KEY).appendPath("dJ5BDC").
appendPath(CommonUtil.Directory.MEDIA_ROOT).
appendPath(DistributedTracing.NR_GUID_ATTRIBUTE).appendPath("2198311517").
appendPath(contentId).appendQueryParameter("assetTypes", "DASH_CENC").
appendQueryParameter("formats", "MPEG-DASH").
appendQueryParameter("format", "smil").build().toString();
~~~

the last version using it is 12.0.28 (211202876). Or, if you already have a
`pid`, you can get it like this:

<http://link.theplatform.com/s/dJ5BDC/i_iqDzG1odu1?format=SMIL&Tracking=true>
